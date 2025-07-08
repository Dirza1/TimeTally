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
SELECT 
    id,
    timestamp AS "Registration Time",
    date_activity AS "Date Activity",
    length_minutes AS "Time minutes",
    description,
    catagory
FROM timeregistration;

-- name: OverviewTimeDates :many
SELECT 
    id,
    timestamp AS "Registration Time",
    date_activity AS "Date Activity",
    length_minutes AS "Time minutes",
    description,
    catagory
FROM timeregistration
WHERE date_activity >= $1 AND date_activity <= $2;

-- name: TotalTimeDates :one
SELECT sum(length_minutes/60.0)
FROM timeregistration
WHERE date_activity >= $1 AND date_activity <= $2;

-- name: DeleteTime :exec
DELETE FROM timeregistration
WHERE id = $1;

-- name: UpdateTime :one
UPDATE timeregistration
SET date_activity = $1, length_minutes = $2, description = $3, catagory = $4
WHERE id = $5
RETURNING *;