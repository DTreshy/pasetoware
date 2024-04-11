package pasetoware

import (
	"strings"

	"github.com/gofiber/fiber/v3"
)

// New PASETO middleware, returns a handler that takes a token in selected lookup param and in case token is valid
// it saves the decrypted token on ctx.Locals, take a look on Config to know more configuration options
func New(authConfigs ...Config) fiber.Handler {
	// Set default authConfig
	config := configDefault(authConfigs...)

	// Return middleware handler
	return func(c fiber.Ctx) error {
		token, err := getToken(c)
		if err != nil {
			return config.ErrorHandler(c, err)
		}

		parsedToken, err := config.Tokener.ParseToken(token)
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

func getToken(c fiber.Ctx) (string, error) {
	authHeader := strings.SplitN(c.Get(fiber.HeaderAuthorization), " ", 2)
	if len(authHeader) != 2 {
		return "", ErrInvalidToken
	}

	if authHeader[0] != TokenPrefix {
		return "", ErrInvalidToken
	}

	if authHeader[1] == "" {
		return "", ErrMissingToken
	}

	return authHeader[1], nil
}
