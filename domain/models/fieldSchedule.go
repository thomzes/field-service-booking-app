package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/thomzes/field-service-booking-app/constants"
)

type FieldSchedule struct {
	ID        uint                          `gorm:"primaryKey;autoIncrement"`
	UUID      uuid.UUID                     `gorm:"type:uuid;not null"`
	FieldID   uint                          `gorm:"type:int;not null"`
	TimeID    uint                          `gorm:"type:int;not null"`
	Date      time.Time                     `gorm:"type:date; not null"`
	Status    constants.FieldScheduleStatus `gorm:"type:int;not null"`
	CreatedAt *time.Time
	UpdateAt  *time.Time
	DeletedAt *time.Time
	Field     Field `gorm:"foreignKey:field_id;references:id;constraint:onUpdate:CASCADE, onDelete:CASCADE"`
	Time      Time  `gorm:"foreignKey:time_id;references:id;constraint:onUpdate:CASCADE, onDelete:CASCADE"`
}
