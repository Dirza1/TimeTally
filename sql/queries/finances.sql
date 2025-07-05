-- name: AddTransaction :one
INSERT INTO finances (id, timestamp,date_transaction, ammount_cent, type, description, catagory)
VALUES(
    gen_random_UUID(),
    NOW,
    $1,
    $2,
    $3,
    $4,
    $5
)
RETURNING *;