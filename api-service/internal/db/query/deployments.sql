-- name: CreateDeployment :one
INSERT INTO "deployments" 
("project_id", "status")
VALUES ( $1, $2 ) RETURNING *;


-- name: GetDeploymentByProjectId :one
SELECT * FROM "deployments"
WHERE "project_id" = $1;

-- -- name: GetDeployments :many
-- SELECT * FROM "deployments"
-- WHERE 
-- "project_id" = $1;

-- -- name: UpdateDeployment :one
-- UPDATE "deployments"
-- SET "status" = COALESCE(sqlc.narg('status'), "status")
-- WHERE 
-- "id" = sqlc.arg("id")
-- RETURNING *;

-- -- -- name: DeleteUpdate :exec
-- -- DELETE FROM "deployments"
-- -- WHERE "id" = sqlc.arg("id");