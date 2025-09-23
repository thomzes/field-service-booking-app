package repositories

import (
	fieldRepo "github.com/thomzes/field-service-booking-app/repositories/field"
	fieldScheduleRepo "github.com/thomzes/field-service-booking-app/repositories/fieldschedule"
	timeScheduleRepo "github.com/thomzes/field-service-booking-app/repositories/time"
	"gorm.io/gorm"
)

type Registry struct {
	db *gorm.DB
}

type IRepositoryRegistry interface {
	GetField() fieldRepo.IFieldRepository
	GetFieldSchedule() fieldScheduleRepo.IFieldScheduleRepository
	GetTime() timeScheduleRepo.ITimeRepository
}

func NewRepositoryRegistry(db *gorm.DB) IRepositoryRegistry {
	return &Registry{db: db}
}

func (r *Registry) GetField() fieldRepo.IFieldRepository {
	return fieldRepo.NewFieldRepository(r.db)
}

func (r *Registry) GetFieldSchedule() fieldScheduleRepo.IFieldScheduleRepository {
	return fieldScheduleRepo.NewFieldScheduleRepository(r.db)
}

func (r *Registry) GetTime() timeScheduleRepo.ITimeRepository {
	return timeScheduleRepo.NewTimeRepository(r.db)
}
