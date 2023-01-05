-- name: CreateAccount :one
insert into accounts
    (owner, balance, currency)
    values (@owner, @balance, @currency)
        returning *;

-- name: GetAccountDetails :one
select * from accounts
    where id= @id
        limit 1;

-- name: GetAccountDetails_WithLock :one
select * from accounts
    where id= @id
        limit 1
            for no key update; --TODO: understand why "no key" was added

-- name: ListAccounts :many
select * from accounts
    order by id
        limit $1
        offset $2;

-- name: UpdateAccount :exec
update accounts
    set balance= @balance
        where id= @id;

-- name: UpdateAccountBalance :one
update accounts
    set balance= balance + @amount
        where id= @id
            returning *;

-- name: DeleteAccount :exec
delete from accounts
    where id= @id;