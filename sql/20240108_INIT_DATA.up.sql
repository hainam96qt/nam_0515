DROP TABLE IF EXISTS users;

create table users (
    id SERIAL primary key,
    name text not null,
    email text not null,
    address text not null,
    password text not null,
    created_at timestamp default current_timestamp
);

CREATE UNIQUE INDEX user_email_idx ON users (email);