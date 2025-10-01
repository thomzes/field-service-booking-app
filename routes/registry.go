package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thomzes/field-service-booking-app/clients"
	"github.com/thomzes/field-service-booking-app/controllers"
	fieldRoute "github.com/thomzes/field-service-booking-app/routes/field"
	fieldScheduleRoute "github.com/thomzes/field-service-booking-app/routes/fieldschedule"
	timeRoute "github.com/thomzes/field-service-booking-app/routes/time"
)

type Registry struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
	client     clients.IClientRegistry
}

type IRouteRegistry interface {
	Serve()
}

func NewRouteRegistry(controller controllers.IControllerRegistry, group *gin.RouterGroup, client clients.IClientRegistry) IRouteRegistry {
	return &Registry{controller: controller, group: group, client: client}
}

func (r *Registry) fieldRoute() fieldRoute.IFieldRoute {
	return fieldRoute.NewFieldRoute(r.controller, r.group, r.client)
}

func (r *Registry) fieldScheduleRoute() fieldScheduleRoute.IFieldScheduleRoute {
	return fieldScheduleRoute.NewFieldScheduleRoute(r.controller, r.group, r.client)
}

func (r *Registry) timeRoute() timeRoute.ITimeRoute {
	return timeRoute.NewTimeRoute(r.controller, r.group, r.client)
}

func (r *Registry) Serve() {
	r.fieldRoute().Run()
	r.fieldScheduleRoute().Run()
	r.timeRoute().Run()
}
