package entity

import "time"

type Timestamp struct {
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
	DeletedAt time.Time `gorm:"-"`
}

func CreatedAt() Timestamp {
	now := time.Now().UTC()

	return Timestamp{
		CreatedAt: now,
		UpdatedAt: now,
	}
}
