package pasetoware_test

import (
	"testing"
	"time"

	"aidanwoods.dev/go-paseto"
	"github.com/DTreshy/pasetoware"
	"github.com/stretchr/testify/require"
)

func TestTokener_GenerateToken(t *testing.T) {
	tokener := &pasetoware.Tokener{
		TokenPrivateKey: paseto.NewV4AsymmetricSecretKey(),
	}

	tests := []struct {
		name       string
		m          map[string]string
		expiration time.Duration
	}{
		{
			name: "Test with valid input",
			m: map[string]string{
				"key1": "value1",
				"key2": "value2",
			},
			expiration: time.Minute * 5, // token will expire in 5 minutes
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token := tokener.GenerateToken(tt.m, tt.expiration)
			require.NotEmpty(t, token, "The token should not be empty")
		})
	}
}
