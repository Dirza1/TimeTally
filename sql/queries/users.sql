-- name: AddAdmin :one
INSERT INTO users(id,name,hashed_password,access_finance,access_timeregistration,administrator)
VALUES(
    gen_random_UUID(),
    $1,
    $2,
    TRUE,
    TRUE,
    TRUE
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
UPDATE users
SET name = $1,
    hashed_password = $2,
    access_finance = $3,
    access_timeregistration = $4,
    administrator = $5
WHERE id = $6
RETURNING name, access_finance,access_timeregistration,administrator;

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
    TRUE,
    TRUE,
    TRUE
) 
RETURNING *;

-- name: CheckOnAdministartor :many
SELECT id, name FROM users
WHERE administrator = TRUE;