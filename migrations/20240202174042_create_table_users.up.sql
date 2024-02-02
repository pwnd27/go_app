create table "users"
(
    "id"                  UUID primary key default (uuid_generate_v4()),
    "username"            varchar not null,
    "hashed_password"     varchar        not null,
    "full_name"           varchar        not null,
    "email"               varchar unique not null,
    "password_changed_at" timestamptz    not null default '0001-01-01 00:00:00z',
    "created_at"          timestamptz    not null default (now()),
    "image"               varchar                 default null
);
