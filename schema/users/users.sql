drop table if exists users;
drop type if exists sex_type;

create type sex_type as enum ('male', 'female', 'other');

create table users
(
    id           varchar(255) primary key,
    company_id   varchar(255) references companies (id),
    display_name varchar(255) not null,
    age          integer,
    icon_url     text
);
