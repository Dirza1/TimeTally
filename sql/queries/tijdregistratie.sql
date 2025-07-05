-- name: AddTimeRegistration :one
INSERT INTO timeregistration (id, timestamp, date_activity, length_minutes, description, catagory)
VALUES(
    gen_random_UUID(),
    NOW(),
    $1,
    $2,
    $3,
    $4
)
RETURNING *;

-- name: ResetTimeRegistration :exec
DELETE FROM timeregistration;

-- name: OverviewAllTime :many
SELECT * FROM timeregistration;

-- name: OverviewTimeMonth :many
SELECT * FROM timeregistration
WHERE EXTRACT(MONTH FROM date_activity) = $1
AND EXTRACT(YEAR FROM date_activity) = $2;

-- name: OverviewTimeYear :many
SELECT * FROM timeregistration
WHERE EXTRACT(YEAR FROM date_activity) = $1;

-- name: TotalTimeMonth :one
SELECT sum(length_minutes/60.0)
FROM timeregistration
WHERE EXTRACT(MONTH FROM date_activity) = $1
AND EXTRACT(YEAR FROM date_activity) = $2;

-- name: TotalTimeYear :one
SELECT sum(length_minutes/60.0)
FROM timeregistration
WHERE EXTRACT(YEAR FROM date_activity) = $1;

-- name: DeleteTime :exec
DELETE FROM timeregistration
WHERE id = $1;

-- name: UpdateTime :one
UPDATE timeregistration
SET date_activity = $1, length_minutes = $2, description = $3, catagory = $4
WHERE id = $5
RETURNING *;