// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CreateDeployment(ctx context.Context, arg CreateDeploymentParams) (Deployment, error)
	CreateProject(ctx context.Context, arg CreateProjectParams) (Project, error)
	CreateRefreshToken(ctx context.Context, arg CreateRefreshTokenParams) (RefreshToken, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteProjects(ctx context.Context, id int64) error
	DeleteUser(ctx context.Context, id int32) error
	GetDeploymentByProjectId(ctx context.Context, projectID pgtype.Int8) (Deployment, error)
	GetProjectById(ctx context.Context, id int64) (Project, error)
	GetProjects(ctx context.Context, createdBy pgtype.Int4) (Project, error)
	GetUser(ctx context.Context, id int32) (User, error)
	Update(ctx context.Context, arg UpdateParams) (User, error)
	UpdateProjects(ctx context.Context, arg UpdateProjectsParams) (Project, error)
}

var _ Querier = (*Queries)(nil)
