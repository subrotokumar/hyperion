CREATE TYPE "deployment_status" AS ENUM ('QUEUE', 'PROGRESS', 'READY', 'FAIL');

CREATE TABLE
    "user" (
        "id" SERIAL PRIMARY KEY,
        "github_id" INT NOT NULL,
        "username" VARCHAR NOT NULL,
        "name" VARCHAR,
        "email" VARCHAR UNIQUE NOT NULL,
        "avatar" VARCHAR DEFAULT ''
    );

CREATE TABLE
    "refresh_token" (
        "id" uuid PRIMARY KEY,
        "token" VARCHAR NOT NULL,
        "user_id" INT NOT NULL,
        "expiry" TIMESTAMPTZ NOT NULL DEFAULT 'now()'
    );

CREATE TABLE
    "projects" (
        "id" BIGSERIAL PRIMARY KEY,
        "created_by" INT,
        "name" VARCHAR NOT NULL,
        "github_url" VARCHAR NOT NULL,
        "subdomain" VARCHAR,
        "custom_domain" VARCHAR,
        "created_at" TIMESTAMPTZ NOT NULL DEFAULT 'now()'
    );

CREATE TABLE
    "deployments" (
        "id" BIGSERIAL PRIMARY KEY,
        "project_id" BIGINT,
        "status" deployment_status NOT NULL DEFAULT 'QUEUE'
    );

ALTER TABLE "refresh_token" ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "projects" ADD FOREIGN KEY ("created_by") REFERENCES "user" ("id");

ALTER TABLE "deployments" ADD FOREIGN KEY ("project_id") REFERENCES "projects" ("id");