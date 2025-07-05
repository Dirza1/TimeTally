-- name: AddTimeRegistration :one
INSERT INTO timeregistration (id, timestamp, date_activity, length_minutes, description, catagory)
VALUES(
    gen_random_UUID(),
    NOW,
    $1,
    $2,
    $3,
    $4
)
RETURNING *;