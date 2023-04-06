package entity

import "github.com/google/uuid"

type Environment struct {
	ID   uuid.UUID `gorm:"column:id;primaryKey"`
	Name string    `gorm:"column:name"`
	Timestamp
	UpdatedBy uuid.UUID `gorm:"column:updated_by"`
}
