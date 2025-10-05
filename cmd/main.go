package cmd

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"strings"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
	"github.com/thomzes/field-service-booking-app/clients"
	"github.com/thomzes/field-service-booking-app/common/gcs"
	"github.com/thomzes/field-service-booking-app/common/response"
	"github.com/thomzes/field-service-booking-app/config"
	"github.com/thomzes/field-service-booking-app/constants"
	"github.com/thomzes/field-service-booking-app/controllers"
	"github.com/thomzes/field-service-booking-app/domain/models"
	"github.com/thomzes/field-service-booking-app/middlewares"
	"github.com/thomzes/field-service-booking-app/repositories"
	"github.com/thomzes/field-service-booking-app/routes"
	"github.com/thomzes/field-service-booking-app/services"
)

var command = &cobra.Command{
	Use:   "serve",
	Short: "Start the server",
	Run: func(cmd *cobra.Command, args []string) {
		_ = godotenv.Load()
		config.Init()
		db, err := config.InitDatabase()
		if err != nil {
			panic(err)
		}

		loc, err := time.LoadLocation("Asia/Jakarta")
		if err != nil {
			panic(err)
		}
		time.Local = loc

		err = db.AutoMigrate(
			&models.Field{},
			&models.FieldSchedule{},
			&models.Time{},
		)
		if err != nil {
			panic(err)
		}

		gcs := initGCS()
		client := clients.NewClientRegistry()
		repository := repositories.NewRepositoryRegistry(db)
		service := services.NewServiceRegistry(repository, gcs)
		controller := controllers.NewControllerRegistry(service)

		router := gin.Default()
		router.Use(middlewares.HandlePanic())
		router.NoRoute(func(ctx *gin.Context) {
			ctx.JSON(http.StatusNotFound, response.Response{
				Status:  constants.Error,
				Message: fmt.Sprintf("Path %s", http.StatusText(http.StatusNotFound)),
			})
		})
		router.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, response.Response{
				Status:  constants.Success,
				Message: "Welcome to field service",
			})
		})
		// handle CORS
		router.Use(func(ctx *gin.Context) {
			ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
			ctx.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT")
			ctx.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, x-sevice-name, x-api-key, x-request-at")
			ctx.Next()
		})

		// handle rate limiter
		lmt := tollbooth.NewLimiter(
			config.Config.RateLimiterMaxRequest,
			&limiter.ExpirableOptions{
				DefaultExpirationTTL: time.Duration(config.Config.RateLimiterTimeSecond) * time.Second,
			})
		router.Use(middlewares.RateLimiter(lmt))

		group := router.Group("api/v1")
		route := routes.NewRouteRegistry(controller, group, client)
		route.Serve()

		port := fmt.Sprintf(":%d", config.Config.Port)
		router.Run(port)
	},
}

func Run() {
	err := command.Execute()
	if err != nil {
		panic(err)
	}
}

func initGCS() gcs.IGCSClient {
	decode, err := base64.StdEncoding.DecodeString(config.Config.GCSPrivateKey)
	if err != nil {
		panic(err)
	}

	stringPrivateKey := strings.ReplaceAll(string(decode), "\\n", "\n")
	gcsServiceAccount := gcs.ServiceAccountKeyJSON{
		Type:                    config.Config.GCSType,
		ProjectID:               config.Config.GCSProjectID,
		PrivateKeyID:            config.Config.GCSPrivateKeyID,
		PrivateKey:              stringPrivateKey,
		ClientEmail:             config.Config.GCSClientEmail,
		ClientID:                config.Config.GCSClientID,
		AuthURI:                 config.Config.GCSAuthURI,
		TokenURI:                config.Config.GCSTokenURI,
		AuthProviderX509CertURL: config.Config.GCSAuthProviderX508CertURL,
		ClientX509CertURL:       config.Config.GCSClientX509CertURL,
		UniverseDomain:          config.Config.GCSUniverseDomain,
	}
	gcsClient := gcs.NewGCSClient(gcsServiceAccount, config.Config.GCSBucketName)
	return gcsClient
}
