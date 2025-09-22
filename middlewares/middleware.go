package middlewares

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"strings"

	"github.com/didip/tollbooth"
	"github.com/didip/tollbooth/limiter"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/thomzes/field-service-booking-app/clients"
	"github.com/thomzes/field-service-booking-app/common/response"
	"github.com/thomzes/field-service-booking-app/config"
	"github.com/thomzes/field-service-booking-app/constants"
	errConstant "github.com/thomzes/field-service-booking-app/constants/error"
)

func HandlePanic() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logrus.Errorf("Recover from panic: %v", err)
				ctx.JSON(http.StatusInternalServerError, response.Response{
					Status:  constants.Error,
					Message: errConstant.ErrInternalServerError.Error(),
				})
				ctx.Abort()
			}
		}()

		ctx.Next()
	}
}

func RateLimiter(lmt *limiter.Limiter) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := tollbooth.LimitByRequest(lmt, ctx.Writer, ctx.Request)
		if err != nil {
			ctx.JSON(http.StatusTooManyRequests, response.Response{
				Status:  constants.Error,
				Message: errConstant.ErrToManyRequest.Error(),
			})
			ctx.Abort()
		}
		ctx.Next()
	}
}

func extractBearerToken(token string) string {
	arrayToken := strings.Split(token, " ")

	if len(arrayToken) == 2 {
		return arrayToken[1]
	}

	return ""
}

func responseUnauthorize(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusUnauthorized, response.Response{
		Status:  constants.Error,
		Message: message,
	})
	ctx.Abort()
}

func validateAPIKey(ctx *gin.Context) error {
	apiKey := ctx.GetHeader(constants.XApiKey)
	requestAt := ctx.GetHeader(constants.XRequestAt)
	serviceName := ctx.GetHeader(constants.XServiceName)
	signatureKey := config.Config.SignatureKey

	validateKey := fmt.Sprintf("%s:%s:%s", serviceName, signatureKey, requestAt)
	hash := sha256.New()
	hash.Write([]byte(validateKey))
	resultHash := hex.EncodeToString(hash.Sum(nil))

	if apiKey != resultHash {
		return errConstant.ErrUnauthorize
	}

	return nil
}

func contains(roles []string, role string) bool {
	for _, r := range roles {
		if r == role {
			return true
		}
	}
	return false
}

func CheckRole(roles []string, client clients.IClientRegistry) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		user, err := client.GetUser().GetUserByToken(ctx.Request.Context())
		if err != nil {
			responseUnauthorize(ctx, errConstant.ErrUnauthorize.Error())
			return
		}

		if !contains(roles, user.Role) {
			responseUnauthorize(ctx, errConstant.ErrUnauthorize.Error())
			return
		}

		ctx.Next()
	}
}

func Authenticate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var err error
		token := ctx.GetHeader(constants.Authorization)
		if token == "" {
			responseUnauthorize(ctx, errConstant.ErrUnauthorize.Error())
			return
		}

		err = validateAPIKey(ctx)
		if err != nil {
			responseUnauthorize(ctx, err.Error())
			return
		}

		ctx.Next()
	}
}

func AuthenticateWithoutToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		err := validateAPIKey(ctx)
		if err != nil {
			responseUnauthorize(ctx, err.Error())
			return
		}

		ctx.Next()
	}
}
