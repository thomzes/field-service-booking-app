package clients

import (
	"github.com/thomzes/field-service-booking-app/clients/config"
	clients "github.com/thomzes/field-service-booking-app/clients/user"
	config2 "github.com/thomzes/field-service-booking-app/config"
)

type ClientRegistry struct{}

type IClientRegistry interface {
	GetUser() clients.IUserClient
}

func NewClientRegistry() IClientRegistry {
	return &ClientRegistry{}
}

func (c *ClientRegistry) GetUser() clients.IUserClient {
	return clients.NewUserClient(
		config.NewClientConfig(
			config.WithBaseURL(config2.Config.InternalService.User.Host),
			config.WithSignatureKey(config2.Config.SignatureKey),
		),
	)
}
