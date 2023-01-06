alter table if exists accounts
    drop constraint if exists owner_currency_key;

alter table if exists accounts
    drop constraint if exists accounts_owner_fkey;

drop table if exists users;