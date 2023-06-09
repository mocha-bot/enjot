package repository

import (
	"context"
	"crypto/rand"
	"time"

	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/mocha-bot/enjot/core/entity"
	repoIntf "github.com/mocha-bot/enjot/core/repository"
)

const (
	DEFAULT_LIMIT = 10
)

type workspaceRepository struct {
	db    *gorm.DB
	table string
}

func NewWorkspaceRepository(db *gorm.DB) repoIntf.WorkspaceRepository {
	return &workspaceRepository{
		db:    db,
		table: "workspace",
	}
}

func (w *workspaceRepository) Create(ctx context.Context, name string, userID uuid.UUID) (entity.Workspace, error) {
	id, err := uuid.NewRandomFromReader(rand.Reader)

	if err != nil {
		return entity.Workspace{}, err
	}

	workspaceEntity := entity.Workspace{
		ID:     id,
		Name:   name,
		UserID: userID,
	}

	err = w.db.Table(w.table).
		Create(&workspaceEntity).
		Error

	if err != nil {
		return entity.Workspace{}, err
	}

	return workspaceEntity, nil
}

func (w *workspaceRepository) Delete(ctx context.Context, workspaceID uuid.UUID) error {
	return w.db.Table(w.table).
		Where("id = ?", workspaceID).
		Update("deleted_at", time.Now().UTC()).
		Error
}

func (w *workspaceRepository) List(
	ctx context.Context,
	page entity.Page,
) (entity.WorkspaceList, error) {
	stmt := w.db.Table(w.table).
		Where("deleted_at IS NULL").
		Offset(page.Offset)

	if page.Limit == 0 {
		page.Limit = DEFAULT_LIMIT
	}

	stmt.Limit(page.Limit)

	rows, err := stmt.Rows()

	if err == gorm.ErrRecordNotFound {
		return entity.WorkspaceList{
			Page: entity.Page{},
		}, nil
	}

	if err != nil {
		return entity.WorkspaceList{}, err
	}

	defer rows.Close()

	var result entity.WorkspaceList

	for rows.Next() {
		var workspace entity.Workspace

		err := w.db.ScanRows(rows, &workspace)

		if err != nil {
			continue
		}

		result.Workspaces = append(result.Workspaces, workspace)
	}

	result.Page = page

	return result, nil
}

func (w *workspaceRepository) Update(
	ctx context.Context,
	workspaceID uuid.UUID,
	name string,
	updatedBy uuid.UUID,
) error {
	return w.db.Table(w.table).
		Where("id = ?", workspaceID).
		Updates(map[string]interface{}{
			"name":       name,
			"updated_by": updatedBy,
		}).
		Error
}
