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

-- name: ResetTransaction :exec
DELETE from finances;

-- name: OverviewAllTransactions :many
SELECT 
    id,
    timestamp AS "Registration Time",
    date_transaction AS "Date Transaction",
    ammount_cent / 100.0 AS "Amount",
    type,
    description,
    catagory
 FROM finances;

-- name: OverviewTransactionsDate :many
SELECT 
    id,
    timestamp AS "Registration Time",
    date_transaction AS "Date Transaction",
    ammount_cent / 100.0 AS "Amount",
    type,
    description,
    catagory
FROM finances
WHERE date_transaction >= $1 AND date_transaction <= $2;

-- name: TotalTransactionsDates :one
SELECT sum(ammount_cent/100.0)
FROM finances
WHERE date_transaction >= $1 AND date_transaction <= $2;

-- name: DeleteTransaction :exec
DELETE FROM finances
WHERE id = $1;

-- name: UpdateTransaction :one
UPDATE finances
SET date_transaction = $1, ammount_cent = $2, type = $3, description = $4, catagory = $5
WHERE id = $6
RETURNING *;