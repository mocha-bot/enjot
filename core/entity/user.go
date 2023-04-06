package entity

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"column:id;primaryKey"`
	Code     string    `gorm:"column:code"`
	Email    string    `gorm:"column:email"`
	Password string    `gorm:"column:password"`
	Timestamp
	UpdatedBy uuid.UUID `gorm:"column:updated_by"`
}
