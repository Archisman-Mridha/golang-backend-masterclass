-- name: CreateAccount :one
insert into accounts
    (owner, balance, currency)
    values (@owner, @balance, @currency)
        returning *;

-- name: GetAccount :one
select * from accounts
    where id= $1
        limit 1;

-- name: ListAccounts :many
select * from accounts
    order by id
        limit $1
        offset $2;

-- name: UpdateAccount :exec
update accounts
    set balance= $2
        where id= $1;

-- name: DeleteAccount :exec
delete from accounts
    where id= $1;