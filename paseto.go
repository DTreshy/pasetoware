package pasetoware

import (
	"strings"

	"github.com/gofiber/fiber/v2"
)

// New PASETO middleware, returns a handler that takes a token in selected lookup param and in case token is valid
// it saves the decrypted token on ctx.Locals, take a look on Config to know more configuration options
func New(authConfigs ...Config) fiber.Handler {
	// Set default authConfig
	config := configDefault(authConfigs...)

	// Return middleware handler
	return func(c *fiber.Ctx) error {
		authHeader := strings.Split(c.Get(fiber.HeaderAuthorization), " ")
		if len(authHeader) != 2 {
			return config.ErrorHandler(c, ErrInvalidToken)
		}

		if authHeader[0] != TokenPrefix {
			return config.ErrorHandler(c, ErrInvalidToken)
		}

		if authHeader[1] == "" {
			return config.ErrorHandler(c, ErrMissingToken)
		}

		parsedToken, err := config.Tokener.ParseToken(authHeader[1])
		if err != nil {
			return config.ErrorHandler(c, err)
		}

		for _, key := range config.PayloadKeys {
			var val string

			if err := parsedToken.Get(key, val); err != nil {
				return config.ErrorHandler(c, err)
			}

			// Store user information from token into context.
			c.Locals(key, val)
		}

		return config.SuccessHandler(c)
	}
}
