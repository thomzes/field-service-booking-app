package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thomzes/field-service-booking-app/clients"
	"github.com/thomzes/field-service-booking-app/constants"
	"github.com/thomzes/field-service-booking-app/controllers"
	"github.com/thomzes/field-service-booking-app/middlewares"
)

type TimeRoute struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
	client     clients.IClientRegistry
}

type ITimeRoute interface {
	Run()
}

func NewTimeRoute(controller controllers.IControllerRegistry, group *gin.RouterGroup, client clients.IClientRegistry) ITimeRoute {
	return &TimeRoute{controller: controller, group: group, client: client}
}

func (t *TimeRoute) Run() {
	group := t.group.Group("/time")
	group.Use(middlewares.Authenticate())
	group.GET("", middlewares.CheckRole([]string{
		constants.Admin,
	}, t.client),
		t.controller.GetTime().GetAll)
	group.GET("/:uuid", middlewares.CheckRole([]string{
		constants.Admin,
	}, t.client),
		t.controller.GetTime().GetByUUID)
	group.POST("", middlewares.CheckRole([]string{
		constants.Admin,
	}, t.client),
		t.controller.GetTime().Create)
}
