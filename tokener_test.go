package pasetoware_test

/*
import (
	"encoding/base64"
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBasicTokenerGenerate(t *testing.T) {
	Tokener := tokener.BasicTokener{}

	for i := 0; i < 1000; i++ {
		token, err := tokener.GenerateToken()
		require.NoError(t, err)
		_, err = Tokener.VerifyAndEncode(token)
		require.NoError(t, err)
	}
}

func TestBasicTokenerVerifyAndEncode(t *testing.T) {
	tests := []struct {
		name          string
		token         string
		expectedToken string
		expectedErr   error
	}{
		{"empty token", "", "", tokener.ErrInvalidToken},
		{"valid too short token", "c29tZXRva2Vu", "", tokener.ErrInvalidToken},
		{"valid too long token", "c29tZXRva2Vuc29tZXRva2Vuc29tZXRva2Vuc29tZXRva2Vuc29tZXRva2Vuc29tZXRva2Vu", "", tokener.ErrInvalidToken},
		{"invalid characaters token", "MTIzNDU2NzgxMjM0NTY3(DEyMzQ)Njc4MTIzNDU2Nzg=", "", base64.CorruptInputError(20)},
		{"valid token", "MTIzNDU2NzgxMjM0NTY3ODEyMzQ1Njc4MTIzNDU2Nzg=", "MTIzNDU2NzgxMjM0NTY3ODEyMzQ1Njc4MTIzNDU2Nzg=", nil},
	}

	tok := tokener.BasicTokener{}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			token, err := tok.VerifyAndEncode(tt.token)
			require.True(t, errors.Is(err, tt.expectedErr))
			require.Equal(t, tt.expectedToken, token)
		})
	}
}*/
