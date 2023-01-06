create table users (
    username varchar primary key,
    hashed_password varchar not null,
    name varchar not null,
    email varchar unique not null,
    created_at timestamptz not null default(now( ))
);

alter table accounts
    add foreign key ("owner")
        references users ("username");

alter table accounts
    add constraint "owner_currency_key"
        unique ("owner", "currency");