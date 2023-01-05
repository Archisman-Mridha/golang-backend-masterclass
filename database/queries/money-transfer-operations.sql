-- name: TransferMoney :one
insert into transfers
    (from_account_id, to_account_id, amount)
    values (@sender_account_id, @receiver_account_id, @amount)
        returning *;

-- name: GetMoneyTransferDetails :one
select * from transfers
    where id= @id
        limit 1;

-- name: ListMoneyTransfers :many
select * from transfers
    where from_account_id= @sender_amount_id or to_account_id= @receiver_amount_id
        order by id
            limit $1 offset $2;