package entity

import "github.com/google/uuid"

type Workspace struct {
	ID     uuid.UUID `gorm:"column:id;primaryKey"`
	Name   string    `gorm:"column:name"`
	UserID uuid.UUID `gorm:"column:user_id"`
	Timestamp
	UpdatedBy uuid.UUID `gorm:"column:updated_by"`
}
