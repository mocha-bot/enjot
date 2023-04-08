package module

import (
	"context"

	"github.com/google/uuid"
	"github.com/mocha-bot/enjot/core/entity"
	"github.com/mocha-bot/enjot/core/repository"
)

type WorkspaceUsecase interface {
	GetWorkspaces(ctx context.Context, limit int, offset int) (entity.WorkspaceList, error)

	CreateWorkspace(ctx context.Context, name string) (entity.Workspace, error)

	UpdateWorkspace(ctx context.Context, workspaceID uuid.UUID, name string) error

	DeleteWorkspace(ctx context.Context, workspaceID uuid.UUID) error
}

type workspaceUsecase struct {
	workspaceRepo repository.WorkspaceRepository
}

func NewWorkspaceUsecase(workspaceRepo repository.WorkspaceRepository) WorkspaceUsecase {
	return &workspaceUsecase{
		workspaceRepo: workspaceRepo,
	}
}

func (w *workspaceUsecase) CreateWorkspace(ctx context.Context, name string) (entity.Workspace, error) {
	userID := ctx.Value("user_id").(string)

	userUID, err := uuid.Parse(userID)

	if err != nil {
		return entity.Workspace{}, err
	}

	workspace, err := w.workspaceRepo.Create(ctx, name, userUID)

	if err != nil {
		return entity.Workspace{}, err
	}

	return workspace, nil
}

func (w *workspaceUsecase) DeleteWorkspace(ctx context.Context, workspaceID uuid.UUID) error {
	return w.workspaceRepo.Delete(ctx, workspaceID)
}

func (w *workspaceUsecase) GetWorkspaces(
	ctx context.Context,
	limit int,
	offset int,
) (entity.WorkspaceList, error) {
	page := entity.Page{
		Offset: offset,
		Limit:  limit,
	}

	workspaceList, err := w.workspaceRepo.List(ctx, page)

	if err != nil {
		return entity.WorkspaceList{}, err
	}

	return workspaceList, nil
}

func (w *workspaceUsecase) UpdateWorkspace(ctx context.Context, workspaceID uuid.UUID, name string) error {
	userID := ctx.Value("user_id").(string)

	updatedBy, err := uuid.Parse(userID)

	if err != nil {
		return err
	}

	return w.workspaceRepo.Update(ctx, workspaceID, name, updatedBy)
}
