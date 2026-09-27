-- name: CreateEntry :one
Insert into entries (account_id,amount) VALUES ($1,$2) RETURNING *;

-- name: UpdateEntry :exec
Update entries set amount=$2 WHERE id=$1;

-- name: DeleteEntry :exec
Delete from entries where id=$1;

-- name: ListEntries :many
Select * from entries where account_id = $1 ORDER BY id LIMIT $2 OFFSET $3;

-- name: GetEntry :one
Select * from entries WHERE id=$1 LIMIT 1;