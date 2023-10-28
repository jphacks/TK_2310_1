package main

import (
	"github.com/giraffe-org/backend/handler"
	"github.com/giraffe-org/backend/handler/auth/signup"
	"github.com/giraffe-org/backend/service"
	"log"

	"go.uber.org/dig"
	"golang.org/x/xerrors"

	APIServerApplication "github.com/giraffe-org/backend/application/apiserver"
	DBRepository "github.com/giraffe-org/backend/repository/db"
)

func initRepository(c *dig.Container) error {
	err := c.Provide(DBRepository.New)
	if err != nil {
		return xerrors.Errorf("DBRepository の DI に失敗しました: %w", err)
	}

	return nil
}

func initService(c *dig.Container) error {

	err := c.Provide(service.NewEventService)
	if err != nil {
		return xerrors.Errorf("signup の DI に失敗しました: %w", err)
	}
	return nil
}

func initHandler(c *dig.Container) error {
	err := c.Provide(handler.NewEventHandler)
	if err != nil {
		return xerrors.Errorf("signup の DI に失敗しました: %w", err)
	}
	err = c.Provide(signup.New)
	if err != nil {
		return xerrors.Errorf("signup の DI に失敗しました: %w", err)
	}

	return nil
}

func initApplication(c *dig.Container) error {
	err := c.Provide(APIServerApplication.New)
	if err != nil {
		return xerrors.Errorf("APIServerApplication の DI に失敗しました: %w", err)
	}

	return nil
}

func start(c *dig.Container) error {
	log.Println("サーバを起動しています...")

	err := c.Invoke(func(dbRepository DBRepository.DB) error {
		err := dbRepository.Migrate()
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return xerrors.Errorf("DB のマイグレーションに失敗しました: %w", err)
	}

	log.Println("DB のマイグレーションに成功しました。")

	err = c.Invoke(func(apiServerApplication APIServerApplication.ApiServer) {
		apiServerApplication.Start()
	})
	if err != nil {
		return xerrors.Errorf("APIServerApplication の起動に失敗しました: %w", err)
	}

	log.Println("APIServerApplication の起動に成功しました。")

	return nil
}

func main() {
	c := dig.New()

	err := initRepository(c)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	err = initService(c)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	err = initHandler(c)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	err = initApplication(c)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}

	err = start(c)
	if err != nil {
		log.Fatalf("%+v\n", err)
	}
}
