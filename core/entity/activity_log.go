package entity

import "github.com/google/uuid"

type ActivityLog struct {
	ID          uuid.UUID `gorm:"column:id;primaryKey"`
	Action      string    `gorm:"column:action"`
	Fields      string    `gorm:"column:fields"`
	UserID      uuid.UUID `gorm:"column:user_id"`
	WorkspaceID uuid.UUID `gorm:"column:workspace_id"`
	Project     uuid.UUID `gorm:"column:project_id"`
	Environment uuid.UUID `gorm:"column:environment_id"`
	Config      uuid.UUID `gorm:"column:config_id"`
	Timestamp
}
