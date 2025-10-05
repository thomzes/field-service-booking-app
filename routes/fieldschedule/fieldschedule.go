package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thomzes/field-service-booking-app/clients"
	"github.com/thomzes/field-service-booking-app/constants"
	"github.com/thomzes/field-service-booking-app/controllers"
	"github.com/thomzes/field-service-booking-app/middlewares"
)

type FieldScheduleRoute struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
	client     clients.IClientRegistry
}

type IFieldScheduleRoute interface {
	Run()
}

func NewFieldScheduleRoute(controller controllers.IControllerRegistry, group *gin.RouterGroup, client clients.IClientRegistry) IFieldScheduleRoute {
	return &FieldScheduleRoute{
		controller: controller,
		group:      group,
		client:     client,
	}
}

func (fs *FieldScheduleRoute) Run() {
	group := fs.group.Group("/field/schedule")
	group.GET("lists/:uuid", middlewares.AuthenticateWithoutToken(), fs.controller.GetFieldSchedule().GetAllByFieldIDAndDate)
	group.PATCH("/status", middlewares.AuthenticateWithoutToken(), fs.controller.GetFieldSchedule().UpdateStatus)
	group.Use(middlewares.Authenticate())
	group.GET("/pagination", middlewares.CheckRole([]string{
		constants.Admin,
		constants.Customer,
	}, fs.client),
		fs.controller.GetFieldSchedule().GetAllWithPagination)
	group.GET("/:uuid", middlewares.CheckRole([]string{
		constants.Admin,
		constants.Customer,
	}, fs.client),
		fs.controller.GetFieldSchedule().GetByUUID)
	group.POST("", middlewares.CheckRole([]string{
		constants.Admin,
	}, fs.client),
		fs.controller.GetFieldSchedule().Create)
	group.POST("/one-month", middlewares.CheckRole([]string{
		constants.Admin,
	}, fs.client),
		fs.controller.GetFieldSchedule().GenerateScheduleForOneMonth)
	group.PUT("/:uuid", middlewares.CheckRole([]string{
		constants.Admin,
	}, fs.client),
		fs.controller.GetFieldSchedule().Update)
	group.DELETE("/:uuid", middlewares.CheckRole([]string{
		constants.Admin,
	}, fs.client),
		fs.controller.GetFieldSchedule().Delete)
}
