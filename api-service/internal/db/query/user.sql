-- name: CreateUser :one
INSERT INTO "user" ("github_id", "name", "username", "email", "avatar")
VALUES ($1, $2, $3, $4, $5)
RETURNING *;


-- name: GetUser :one
SELECT * from "user" 
WHERE "id" = $1;

-- name: Update :one
UPDATE "user"
SET
    "github_id" = COALESCE(sqlc.narg('github_id'), "github_id"),
    "name" = COALESCE(sqlc.narg('name'), "name"),
    "avatar" = COALESCE(sqlc.narg('avatar'), "avatar"),
    "email" = COALESCE(sqlc.narg('email'), "email"),
    "username" = COALESCE(sqlc.narg('username'), "username")
WHERE "id" = sqlc.arg('id')
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM "user"
WHERE id = $1;

