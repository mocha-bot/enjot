package entity

import "github.com/google/uuid"

type ProjectEnvironment struct {
	ID          uuid.UUID `gorm:"column:id;primaryKey"`
	WorkspaceID uuid.UUID `gorm:"column:workspace_id"`
	Project     uuid.UUID `gorm:"column:project_id"`
	Environment uuid.UUID `gorm:"column:environment_id"`
	Config      uuid.UUID `gorm:"column:config_id"`
}
