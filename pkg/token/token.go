package token

import (
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrTokenNotValid = fmt.Errorf("token method is not valid")
)

// Generate
//
// Generating a new token (which can be an `idToken`,`accessToken` or `refreshToken`)
// that will later be used by the client for authorization and permission to access the API.
//
// For the required claims, you can refer to [here](https://github.com/golang-jwt/jwt/blob/v4/claims.go#L24-L45)
// These are the standard claims data that need to be filled.
//
// All the claims are mandatory except for `nbf“ and `jti“.
func Generate(mapClaims jwt.MapClaims, key []byte) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)

	tokenString, err := token.SignedString(key)

	return tokenString, err
}

// Verify
//
// Given token as a string, we need to verify it to check is that token is valid or not
// and whether the token is already expire or not
func Verify(mapClaims jwt.MapClaims, key []byte, tokenString string) (bool, error) {
	token, err := jwt.ParseWithClaims(tokenString, mapClaims, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)

		if !ok {
			return nil, ErrTokenNotValid
		}

		return key, nil
	})

	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, nil
	}

	return true, err
}
