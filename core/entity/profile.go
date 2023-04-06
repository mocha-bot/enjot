package entity

import "github.com/google/uuid"

type Profile struct {
	ID       uuid.UUID `gorm:"column:id;primaryKey"`
	FullName string    `gorm:"column:full_name"`
	UserID   uuid.UUID `gorm:"column:user_id"`
	Timestamp
	UpdatedBy uuid.UUID `gorm:"column:updated_by"`
}
