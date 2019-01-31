package config

import (
	"github.com/go-ozzo/ozzo-validation"
)

//
// All available configurations for the micro-service
//
const (
	securityAPIKeyEnvName = "SECURITY_API_KEY"
)

type ConfigurationContainer struct {
	SecurityAPIKey string
}

func (c *ConfigurationContainer) Validate() error {
	return validation.ValidateStruct(c,
		validation.Field(&c.SecurityAPIKey, validation.Required),
	)
}
