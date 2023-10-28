drop table if exists participant;
drop type if exists participation_status;

/**
 * 参加状況
 * - not_completed: 未報告（イベント終了後2時間まで） or 欠席
 * - completed: イベント参加済み
 */
create type participation_status as enum ('not_completed', 'completed');

create table participants
(
    user_id  varchar(255) references users (id),
    event_id varchar(255) references events (id),
    status   participation_status not null default 'not_completed',
    primary key (user_id, event_id)
);
