-- Drop foreign key constraints
ALTER TABLE "deployments"
DROP CONSTRAINT IF EXISTS "deployments_project_id_fkey";

ALTER TABLE "projects"
DROP CONSTRAINT IF EXISTS "projects_created_by_fkey";

ALTER TABLE "refresh_token"
DROP CONSTRAINT IF EXISTS "refresh_token_user_id_fkey";

-- Drop tables
DROP TABLE IF EXISTS "deployments";

DROP TABLE IF EXISTS "projects";

DROP TABLE IF EXISTS "refresh_token";

DROP TABLE IF EXISTS "user";

-- Drop custom enum type
DROP TYPE IF EXISTS "deployment_status";