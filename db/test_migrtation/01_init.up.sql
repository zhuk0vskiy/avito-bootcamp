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
    is_moderator boolean not null ,
    totp_secret bytea not null
);

create table if not exists houses
(
    id   uuid primary key default gen_random_uuid(),
    creation_time timestamp not null,
    creator_id uuid not null references users(id),
    address text not null,
    max_flats smallint not null,
    last_update_time timestamp not null
);

create table if not exists apartments
(
    id   uuid primary key default gen_random_uuid(),
    creation_time timestamp not null,
    creator_id uuid not null references users(id),
    house_id uuid not null references houses(id),
    price integer not null,
    rooms smallint not null,
    status_update_time timestamp,
    moderator_id uuid not null references users(id)
);