package db

import (
	"fmt"
	"log"
	"os"

	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"golang.org/x/xerrors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/jphacks/TK_2310_1/config"
	"github.com/jphacks/TK_2310_1/entity"
)

var db *gorm.DB

type postgresImpl struct {
	client *gorm.DB
}

func init() {
	var err error
	if os.Getenv("APP_ENV") == "prd" || os.Getenv("APP_ENV") == "stg" || os.Getenv("APP_ENV") == "dev" {
		db, err = gorm.Open(postgres.New(postgres.Config{
			DriverName: "cloudsqlpostgres",
			DSN: fmt.Sprintf(
				"host=%s user=%s password=%s dbname=%s sslmode=disable",
				config.Get().PostgresHost,
				config.Get().PostgresUser,
				config.Get().PostgresPass,
				config.Get().PostgresDB,
			),
		}))
	} else {
		dsn := fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=verify-ca sslrootcert=%s sslcert=%s sslkey=%s",
			config.Get().PostgresHost,
			config.Get().PostgresPort,
			config.Get().PostgresUser,
			config.Get().PostgresPass,
			config.Get().PostgresDB,
			config.Get().PostgresCA,
			config.Get().PostgresCert,
			config.Get().PostgresKey,
		)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	}
	if err != nil {
		log.Fatalf("データベースへの接続に失敗しました: %+v", err)
	}
}

func New() DB {
	return &postgresImpl{
		client: db,
	}
}

// Migrate は　DB のマイグレーションをする関数です
func (p *postgresImpl) Migrate() error {
	// テーブル作成に必要な enum を作成する
	createEnum(db)

	err := db.AutoMigrate(
		&entity.User{},
		&entity.Company{},
		&entity.Event{},
		&entity.Application{},
		&entity.Participant{},
	)
	if err != nil {
		return xerrors.Errorf("db のマイグレーションに失敗しました: %w", err)
	}

	// テーブルで使われる function を作成する
	createFunction(db)
	// テーブルで使われる trigger を作成する
	createTrigger(db)

	return nil
}

func (p *postgresImpl) Insert(model interface{}) error {
	err := db.Create(model).Error
	if err != nil {
		return xerrors.Errorf("データの挿入に失敗しました。model -> %+v: %w", model, err)
	}

	return nil
}

func (p *postgresImpl) GetDB() *gorm.DB {
	return db
}

// createEnum はテーブル作成の際に使われる enum を作成する関数です
// 参考: https://github.com/go-gorm/gorm/issues/1978
func createEnum(db *gorm.DB) {
	db.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'sex_type') THEN
				CREATE TYPE sex_type AS ENUM ('male', 'female', 'other');
			END IF;
		END
		$$;
	`)
	db.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'company_category') THEN
				CREATE TYPE company_category AS ENUM ('it', 'manufacturing', 'service', 'others');
			END IF;
		END
		$$;
	`)
	db.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'application_status') THEN
				CREATE TYPE application_status AS ENUM ('participant', 'absent');
			END IF;
		END
		$$;
	`)
	db.Exec(`
		DO $$
		BEGIN
			IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'participation_status') THEN
				CREATE TYPE participation_status AS ENUM ('not_completed', 'completed');
			END IF;
		END
		$$;
	`)
}

// createFunction はテーブルで使われる function を作成する関数です
func createFunction(db *gorm.DB) {
	db.Exec(`
		DO $$
		BEGIN
		   IF NOT EXISTS (SELECT 1 FROM pg_proc WHERE proname = 'refresh_updated_at_step1') THEN
			  EXECUTE '
			  CREATE FUNCTION refresh_updated_at_step1() RETURNS trigger AS
			  $BODY$
			  BEGIN
				 IF NEW.updated_at = OLD.updated_at THEN
					NEW.updated_at := NULL;
				 END IF;
				 RETURN NEW;
			  END;
			  $BODY$ LANGUAGE plpgsql;';
		   END IF;
		END $$;
	`)
	db.Exec(`
		DO $$
		BEGIN
		   IF NOT EXISTS (SELECT 1 FROM pg_proc WHERE proname = 'refresh_updated_at_step2') THEN
			  EXECUTE '
			  CREATE FUNCTION refresh_updated_at_step2() RETURNS trigger AS
			  $BODY$
			  BEGIN
				 IF NEW.updated_at IS NULL THEN
					NEW.updated_at := OLD.updated_at;
				 END IF;
				 RETURN NEW;
			  END;
			  $BODY$ LANGUAGE plpgsql;';
		   END IF;
		END $$;
	`)
	db.Exec(`
		DO $$
		BEGIN
		   IF NOT EXISTS (SELECT 1 FROM pg_proc WHERE proname = 'refresh_updated_at_step3') THEN
			  EXECUTE '
			  CREATE FUNCTION refresh_updated_at_step3() RETURNS trigger AS
			  $BODY$
			  BEGIN
				 IF NEW.updated_at IS NULL THEN
					NEW.updated_at := CURRENT_TIMESTAMP;
				 END IF;
				 RETURN NEW;
			  END;
			  $BODY$ LANGUAGE plpgsql;';
		   END IF;
		END $$;
	`)
}

// createTrigger はテーブルで使われる trigger を作成する関数です
func createTrigger(db *gorm.DB) {
	db.Exec(`
		DO
		$$
		BEGIN
			-- トリガーが存在しない場合のみ実行
			IF NOT EXISTS (
				SELECT 1
				FROM pg_trigger
				WHERE tgname = 'refresh_users_updated_at_step1'
			) THEN
				-- トリガーの作成
				EXECUTE '
					CREATE TRIGGER refresh_users_updated_at_step1
					BEFORE UPDATE
					ON applications
					FOR EACH ROW
					EXECUTE PROCEDURE refresh_updated_at_step1();
				';
			END IF;
		END
		$$;
	`)
	db.Exec(`
		DO
		$$
		BEGIN
			-- トリガーが存在しない場合のみ実行
			IF NOT EXISTS (
				SELECT 1
				FROM pg_trigger
				WHERE tgname = 'refresh_users_updated_at_step2'
			) THEN
				-- トリガーの作成
				EXECUTE '
					CREATE TRIGGER refresh_users_updated_at_step2
					BEFORE UPDATE
					ON applications
					FOR EACH ROW
					EXECUTE PROCEDURE refresh_updated_at_step2();
				';
			END IF;
		END
		$$;
	`)
	db.Exec(`
		DO
		$$
		BEGIN
			-- トリガーが存在しない場合のみ実行
			IF NOT EXISTS (
				SELECT 1
				FROM pg_trigger
				WHERE tgname = 'refresh_users_updated_at_step3'
			) THEN
				-- トリガーの作成
				EXECUTE '
					CREATE TRIGGER refresh_users_updated_at_step3
					BEFORE UPDATE
					ON applications
					FOR EACH ROW
					EXECUTE PROCEDURE refresh_updated_at_step3();
				';
			END IF;
		END
		$$;
	`)
}
