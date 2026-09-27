-- name: CreateTransfer :one
Insert into transfers (from_account_id,to_account_id,amount) VALUES ($1,$2,$3) RETURNING *;

-- name: UpdateTransfer :exec
Update transfers set amount=$2 WHERE id=$1;

-- name: DeleteTransfer :exec
Delete from transfers where id=$1;

-- name: ListTransfers :many
SELECT * FROM transfers WHERE  from_account_id = $1 OR to_account_id = $2 ORDER BY id LIMIT $3 OFFSET $4;

-- name: GetTransfer :one
Select * from transfers WHERE id=$1 LIMIT 1;