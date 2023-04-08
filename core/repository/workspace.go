package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/mocha-bot/enjot/core/entity"
)

type WorkspaceRepository interface {
	Create(ctx context.Context, name string, userID uuid.UUID) (entity.Workspace, error)

	List(ctx context.Context, page entity.Page) (entity.WorkspaceList, error)

	Update(ctx context.Context, workspaceID uuid.UUID, name string, updatedBy uuid.UUID) error

	Delete(ctx context.Context, workspaceID uuid.UUID) error
}
