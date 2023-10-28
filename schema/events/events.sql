drop table if exists events;

create table events
(
    id                           varchar(255) primary key,
    title                        varchar(255)   not null,
    host_company_id              varchar(255)   not null references companies (id),
    description                  text             not null,
    address                      varchar(255)   not null,
    latitude                     double precision not null,
    longitude                    double precision not null,
    participant_count            integer          not null,
    unit_price                   integer          not null,
    will_start_at                timestamptz      not null,
    will_complete_at             timestamptz      not null,
    application_deadline         timestamptz      not null,
    leader                       varchar(255) references users (id),
    started_at                   timestamptz,
    completed_at                 timestamptz,
    proof_participants_image_url text,
    proof_garbage_image_url      text,
    report                       text
);

-- https://developers.google.com/maps/documentation/javascript/geocoding?hl=ja
-- このページを見る限り、緯度、軽度は小数点以下7桁までしか帰ってきていないみたい。
-- そのため、必要となる桁数は整数部分3桁、小数部分7桁でたかだか10桁となる。

-- https://www.postgresql.jp/document/9.3/html/datatype-numeric.html
-- このページにも書いてある通り、double precision は小数点以下15桁まで扱える。そのため、latitude、longitude は double precision で扱います。

-- https://www.postgresql.jp/document/7.2/user/datatype-datetime.html
-- timestamptz はタイムゾーン付きの日時を扱う型です。これを使うことで、日本時間でのイベント開催日時を扱うことができます。

-- description カラムはその event の説明文を保存しておく用のカラムです。
