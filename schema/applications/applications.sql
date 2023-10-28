drop table if exists applications;
drop type if exists application_status;
drop function if exists refresh_updated_at_step1;
drop function if exists refresh_updated_at_step2;
drop function if exists refresh_updated_at_step3;

/**
 * 申請の状態
 * - participant: 参加者
 * - absent: 不参加 (キャンセル)
 */
create type application_status as enum (
    'participant',
    'absent'
    );

create table applications
(
    id         varchar(255) primary key,
    user_id    varchar(255)     not null references users (id),
    event_id   varchar(255)     not null references events (id),
    status     application_status not null default 'applied',
    created_at timestamptz        not null default current_timestamp,
    updated_at timestamptz        not null default current_timestamp
);

-- 以下、updated_at の更新を制御するためのトリガーです
-- 参照: https://zenn.dev/mpyw/articles/rdb-ids-and-timestamps-best-practices#updated_at-%EF%BC%88update-%E6%99%82%E3%81%AE%E3%83%87%E3%83%95%E3%82%A9%E3%83%AB%E3%83%88%E5%9F%8B%E3%82%81%EF%BC%89
CREATE FUNCTION refresh_updated_at_step1() RETURNS trigger AS
$$
BEGIN
    IF NEW.updated_at = OLD.updated_at THEN
        NEW.updated_at := NULL;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE FUNCTION refresh_updated_at_step2() RETURNS trigger AS
$$
BEGIN
    IF NEW.updated_at IS NULL THEN
        NEW.updated_at := OLD.updated_at;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE FUNCTION refresh_updated_at_step3() RETURNS trigger AS
$$
BEGIN
    IF NEW.updated_at IS NULL THEN
        NEW.updated_at := CURRENT_TIMESTAMP;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER refresh_users_updated_at_step1
    BEFORE UPDATE
    ON applications
    FOR EACH ROW
EXECUTE PROCEDURE refresh_updated_at_step1();
CREATE TRIGGER refresh_users_updated_at_step2
    BEFORE UPDATE OF updated_at
    ON applications
    FOR EACH ROW
EXECUTE PROCEDURE refresh_updated_at_step2();
CREATE TRIGGER refresh_users_updated_at_step3
    BEFORE UPDATE
    ON applications
    FOR EACH ROW
EXECUTE PROCEDURE refresh_updated_at_step3();
