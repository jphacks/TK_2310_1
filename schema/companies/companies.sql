drop table if exists companies;
drop type if exists company_category;

create type company_category as enum ('it', 'manufacturing', 'service', 'others');

create table companies
(
    id           varchar(255) primary key,
    display_name varchar(255) not null,
    description  text,
    company_url  text,
    icon_url     text
);
