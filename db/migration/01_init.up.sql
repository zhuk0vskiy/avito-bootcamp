create extension if not exists pgcrypto;

-- CREATE ROLE rwaccess;
-- GRANT CONNECT ON DATABASE avito TO rwaccess;
-- GRANT INSERT, UPDATE, DELETE ON ALL TABLES IN SCHEMA public TO rwaccess;
-- CREATE USER avito WITH PASSWORD 'avito';
-- GRANT rwaccess TO avito;
-- ALTER DEFAULT PRIVILEGES IN SCHEMA public GRANT SELECT, INSERT, UPDATE, DELETE ON TABLES TO rwaccess;

create table if not exists users
(
    id   uuid primary key default gen_random_uuid() ,
    creation_time timestamp not null,
    email text not null unique,
    password bytea not null,
    is_moderator boolean not null default false,
    totp_secret bytea not null
);

create table if not exists houses
(
    id   uuid primary key default gen_random_uuid(),
    creation_time timestamp not null,
    creator_id uuid not null references users(id),
    address text not null,
    max_apartments integer not null,
    apartments_update_time timestamp not null
);

create table if not exists apartments
(
    id   uuid primary key default gen_random_uuid(),
    creation_time timestamp not null,
    creator_id uuid not null references users(id),
    house_id uuid not null references houses(id),
    price bigint not null,
    rooms integer not null,
    status_update_time timestamp,
    moderator_id uuid not null references users(id)
);

create table if not exists subscribers
(
    id uuid primary key default gen_random_uuid(),
    creation_time timestamp not null,
    user_id uuid not null references users(id),
    house_id uuid not null references houses(id)
);

create index idx_subsciber_house_id on subscribers using hash (house_id);

create table if not exists notices
(
    id   uuid primary key default gen_random_uuid(),
    creation_time timestamp not null,
    subscriber_id uuid not null references users(id),
    apartment_id uuid not null references apartments(id),
    house_id uuid not null references houses(id)
);

create table if not exists notices_outbox
(
    id uuid primary key default gen_random_uuid(),
    notice_id uuid not null references notices(id),
    is_send bool default false
);