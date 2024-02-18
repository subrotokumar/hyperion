-- name: CreateProject :one
INSERT INTO
    "projects" (
        "created_by",
        "name",
        "github_url",
        "subdomain",
        "custom_domain"
    )
VALUES
    ($1, $2, $3, $4, $5) RETURNING *;


-- name: GetProjects :one
SELECT * FROM "projects"
WHERE "created_by" = $1;

-- name: GetProjectById :one
SELECT * FROM "projects"
WHERE "id" = $1;

-- name: UpdateProjects :one
UPDATE "projects"
SET
"name" = COALESCE(sqlc.narg('name'), "name"),
"custom_domain" = COALESCE(sqlc.narg('custom_domain'), "custom_domain")
WHERE "id" = sqlc.arg('id')
RETURNING "id", "created_by", "name", "github_url", "subdomain", "custom_domain", "created_at";

-- name: DeleteProjects :exec
DELETE FROM "projects"
WHERE id = $1;
