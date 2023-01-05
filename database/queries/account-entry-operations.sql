-- name: CreateEntry :one
insert into entries
    (account_id, amount)
    values (@account_id, @amount)
        returning *;

-- name: GetEntryDetails :one
select * from entries
    where id = @id
        limit $1;

-- name: ListEntries :many
select * from entries
    where account_id = @account_id
        order by id
            limit $1 offset $2;
