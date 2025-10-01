package services

import (
	"github.com/thomzes/field-service-booking-app/common/gcs"
	"github.com/thomzes/field-service-booking-app/repositories"
	fieldService "github.com/thomzes/field-service-booking-app/services/field"
	fieldScheduleService "github.com/thomzes/field-service-booking-app/services/fieldschedule"
	timeService "github.com/thomzes/field-service-booking-app/services/time"
)

type Registry struct {
	repository repositories.IRepositoryRegistry
	gcs        gcs.IGCSClient
}

type IServiceRegistry interface {
	GetField() fieldService.IFieldService
	GetFieldSchedule() fieldScheduleService.IFieldScheduleService
	GetTime() timeService.ITimeService
}

func NewServiceRegistry(repository repositories.IRepositoryRegistry, gcs gcs.IGCSClient) IServiceRegistry {
	return &Registry{repository: repository, gcs: gcs}
}

func (r *Registry) GetField() fieldService.IFieldService {
	return fieldService.NewFieldService(r.repository, r.gcs)
}

func (r *Registry) GetFieldSchedule() fieldScheduleService.IFieldScheduleService {
	return fieldScheduleService.NewFieldScheduleService(r.repository)
}

func (r *Registry) GetTime() timeService.ITimeService {
	return timeService.NewTimeService(r.repository)
}
