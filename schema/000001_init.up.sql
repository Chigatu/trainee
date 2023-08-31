CREATE TABLE segments
(
    id serial primary key,
    slug varchar(255) not null unique
);

CREATE TABLE users_segments
(
    id serial primary key,
    user_id serial not null,
    segment_id serial not null,
    foreign key (segment_id) references segments (id) on delete cascade,
    unique (user_id, segment_id)
);

CREATE TYPE segment_operation AS ENUM('add', 'delete');