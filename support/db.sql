drop table if exists `users`;

create table `users`(
    `uid`  char(64) primary key
);

create index `users_index_uid` on `users`(`uid`);

drop table if exists `sensor_data`;

create table `sensor_data`(
    `id`    integer primary key AUTOINCREMENT,
    `uid`   char(64) not null,
    `x`     double not null,
    `y`     double not null,
    `z`     double not null,
    `time`  bigint not null,
    `state` smallint null
);

drop index if exists `sensor_data_index_id`;
drop index if exists `sensor_data_index_uid`;
drop index if exists `sensor_data_index_state`;

create index `sensor_data_index_id` on `sensor_data`(`id`);
create index `sensor_data_index_uid` on `sensor_data`(`uid`);
create index `sensor_data_index_state` on `sensor_data`(`state`);



drop table if exists `history_data`;

create table `history_data`(
    `id`        integer primary key AUTOINCREMENT,
    `uid`       char(64) not null,
    `latitude`  double not null,
    `longitude` double not null,
    `time`      bigint not null,
    `state`     smallint null
);

drop index if exists `history_data_id`;
drop index if exists `history_data_uid`;
drop index if exists `history_data_state`;

create index `history_data_id` on `history_data`(`id`);
create index `history_data_uid` on `history_data`(`uid`);
create index `history_data_state` on `history_data`(`state`);
