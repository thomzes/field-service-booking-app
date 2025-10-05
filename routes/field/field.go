package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/thomzes/field-service-booking-app/clients"
	"github.com/thomzes/field-service-booking-app/constants"
	"github.com/thomzes/field-service-booking-app/controllers"
	"github.com/thomzes/field-service-booking-app/middlewares"
)

type FieldRoute struct {
	controller controllers.IControllerRegistry
	group      *gin.RouterGroup
	client     clients.IClientRegistry
}

type IFieldRoute interface {
	Run()
}

func NewFieldRoute(controller controllers.IControllerRegistry, group *gin.RouterGroup, client clients.IClientRegistry) IFieldRoute {
	return &FieldRoute{
		controller: controller,
		group:      group,
		client:     client,
	}
}

func (f *FieldRoute) Run() {
	group := f.group.Group("/field")
	group.GET("", middlewares.AuthenticateWithoutToken(), f.controller.GetField().GetAllWithoutPagination)
	group.GET(":uuid", middlewares.AuthenticateWithoutToken(), f.controller.GetField().GetByUUID)
	group.Use(middlewares.Authenticate())
	group.GET("/pagination", middlewares.CheckRole([]string{
		constants.Admin,
		constants.Customer,
	}, f.client),
		f.controller.GetField().GetAllWithPagination)
	group.POST("", middlewares.CheckRole([]string{
		constants.Admin,
	}, f.client),
		f.controller.GetField().Create)
	group.PUT("/:uuid", middlewares.CheckRole([]string{
		constants.Admin,
	}, f.client),
		f.controller.GetField().Update)
	group.DELETE("/:uuid", middlewares.CheckRole([]string{
		constants.Admin,
	}, f.client),
		f.controller.GetField().Delete)
}
