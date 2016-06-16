drop table if exists `sensor_data`;

create table sensor_data(
    `id` integer primary key AUTOINCREMENT,
    `uid` char(64) not null,
    `x` double not null,
    `y` double not null,
    `z` double not null,
    `time` bigint not null,
    `state` smallint null
);

drop table if exists `history_data`;

create table history_data(
    `id` integer primary key AUTOINCREMENT,
    `uid` char(64) not null,
    `latitude` double not null,
    `longitude` double not null,
    `time` bigint not null,
    `state` smallint null
);
