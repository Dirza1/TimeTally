-- name: AddAdmin :one
INSERT INTO users(id,name,hashed_password,access_finance,access_timeregistration,administrator)
VALUES(
    gen_random_UUID(),
    $1,
    $2,
    1,
    1,
    1
) 
RETURNING *;

-- name: AddUser :one
INSERT INTO users(id,name,hashed_password,access_finance,access_timeregistration,administrator)
VALUES(
    gen_random_UUID(),
    $1,
    $2,
    $3,
    $4,
    0
) 
RETURNING *;

-- name: UpdateUser :one
INSERT INTO users(id,name,hashed_password,access_finance,access_timeregistration,administrator)
VALUES(
    gen_random_UUID(),
    $1,
    $2,
    $3,
    $4,
    $5
) 
RETURNING *;

-- name: Login :one
SELECT name,hashed_password,id FROM users
WHERE name = $1;

-- name: GetUserPermissions :one
SELECT access_finance, access_timeregistration, administrator FROM users
WHERE name = $1;

-- name: CreateFirstAdministartor :one
INSERT INTO users(id,name,hashed_password,access_finance,access_timeregistration,administrator)
VALUES(
    gen_random_UUID(),
    $1,
    $2,
    $3,
    $4,
    0
) 
RETURNING *;

-- name: CheckOnAdministartor :many
SELECT * FROM users
WHERE administrator = 1;