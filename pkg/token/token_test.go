package token

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/assert"
)

func TestGenerate(t *testing.T) {
	t.Run("generate token should be success", func(t *testing.T) {
		mapClaims := jwt.MapClaims{
			"iss": "https://example.com",
			"sub": "1234567890",
			"aud": "https://api.example.com",
			"iat": jwt.NewNumericDate(time.Now()),
			"exp": jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		}

		key := []byte("testmysecretsafely")

		tokenString, err := Generate(mapClaims, key)

		assert.NoError(t, err)

		t.Log("token:", tokenString)
	})
}

func TestVerify(t *testing.T) {
	type args struct {
		mapClaims   jwt.MapClaims
		key         []byte
		tokenString string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "verify token should be success",
			args: args{
				mapClaims:   make(jwt.MapClaims),
				key:         []byte("testmysecretsafely"),
				tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJodHRwczovL2FwaS5leGFtcGxlLmNvbSIsImV4cCI6MTY4MDkwNzAwNywiaWF0IjoxNjgwODIwNjA3LCJpc3MiOiJodHRwczovL2V4YW1wbGUuY29tIiwic3ViIjoiMTIzNDU2Nzg5MCJ9.eeR6lw2widsaQ5HZxl-U7YvLjuMWPZPDpIzn7bBT7m8",
			},
			want:    true,
			wantErr: false,
		},
		{
			name: "verify token should be failed token expired",
			args: args{
				mapClaims:   make(jwt.MapClaims),
				key:         []byte("testmysecretsafely"),
				tokenString: "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJodHRwczovL2FwaS5leGFtcGxlLmNvbSIsImV4cCI6MTY4MDgyMDA5NiwiaWF0IjoxNjgwODIwMDk2LCJpc3MiOiJodHRwczovL2V4YW1wbGUuY29tIiwic3ViIjoiMTIzNDU2Nzg5MCJ9.dUbSnnGk2VUFawDFWgp0KxH8p_RbFyxfRrmZiPWwiPg",
			},
			want:    false,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Verify(tt.args.mapClaims, tt.args.key, tt.args.tokenString)
			if (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Verify() = %v, want %v", got, tt.want)
			}
		})
	}
}
