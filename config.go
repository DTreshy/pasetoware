package pasetoware

import (
	"errors"

	"github.com/gofiber/fiber/v2"
)

var (
	ErrExpiredToken   = errors.New("token has expired")
	ErrMissingToken   = errors.New("missing PASETO token")
	ErrInvalidToken   = errors.New("invalid PASETO token")
	ErrDataUnmarshal  = errors.New("can't unmarshal token data to Payload type")
	TokenAudience     = "gofiber.gophers"
	TokenSubject      = "user-token"
	TokenField        = "token"
	TokenPrefix       = "Bearer"
	DefaultContextKey = "payload"
)

// Config defines the config for PASETO middleware
type Config struct {
	// Filter defines a function to skip middleware.
	// Optional. Default: nil
	Next func(*fiber.Ctx) bool

	// SuccessHandler defines a function which is executed for a valid token.
	// Optional. Default: c.Next()
	SuccessHandler fiber.Handler

	// ErrorHandler defines a function which is executed for an invalid token.
	// It may be used to define a custom PASETO error.
	// Optional. Default: 401 Invalid or expired PASETO
	ErrorHandler fiber.ErrorHandler

	Tokener *Tokener

	// ContextKey to store user information from the token into context.
	// Optional. Default: DefaultContextKey.
	ContextKey string

	PayloadKeys []string
}

// ConfigDefault is the default config
var ConfigDefault = Config{
	Next:           nil,
	SuccessHandler: nil,
	ErrorHandler:   nil,
	Tokener:        NewTokener(),
	ContextKey:     DefaultContextKey,
}

func defaultErrorHandler(c *fiber.Ctx, err error) error {
	return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
}

// Helper function to set default values
func configDefault(authConfigs ...Config) Config {
	// Return default authConfigs if nothing provided
	config := ConfigDefault
	if len(authConfigs) > 0 {
		// Override default authConfigs
		config = authConfigs[0]
	}

	// Set default values
	if config.Next == nil {
		config.Next = ConfigDefault.Next
	}

	if config.SuccessHandler == nil {
		config.SuccessHandler = func(c *fiber.Ctx) error {
			return c.Next()
		}
	}

	if config.ErrorHandler == nil {
		config.ErrorHandler = defaultErrorHandler
	}

	return config
}
