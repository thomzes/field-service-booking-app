package controllers

import (
	fieldController "github.com/thomzes/field-service-booking-app/controllers/field"
	fieldScheduleController "github.com/thomzes/field-service-booking-app/controllers/fieldschedule"
	timeController "github.com/thomzes/field-service-booking-app/controllers/time"
	"github.com/thomzes/field-service-booking-app/services"
)

type Registry struct {
	service services.IServiceRegistry
}

type IControllerRegistry interface {
	GetField() fieldController.IFieldController
	GetFieldSchedule() fieldScheduleController.IFieldScheduleController
	GetTime() timeController.ITimeController
}

func NewRegistryController(service services.IServiceRegistry) IControllerRegistry {
	return &Registry{service: service}
}

func (r *Registry) GetField() fieldController.IFieldController {
	return fieldController.NewFieldController(r.service)
}

func (r *Registry) GetFieldSchedule() fieldScheduleController.IFieldScheduleController {
	return fieldScheduleController.NewFieldScheduleController(r.service)
}

func (r *Registry) GetTime() timeController.ITimeController {
	return timeController.NewTimeController(r.service)
}
