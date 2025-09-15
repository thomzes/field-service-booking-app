package models

import (
	"time"

	"github.com/google/uuid"
)

type Time struct {
	ID        uint      `gorm:"primaryKey;autoIncrement"`
	UUID      uuid.UUID `gorm:"type:uuid;not null"`
	StartTime string    `gorm:"type:time without timezone;not null"`
	EndTime   string    `gorm:"type:time without timezone; not null"`
	CreatedAt *time.Time
	UpdateAt  *time.Time
}
