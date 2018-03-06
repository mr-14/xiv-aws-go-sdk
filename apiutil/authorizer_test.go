package apiutil

import (
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func TestValidateJWT(t *testing.T) {
	tokenKey := "1234"
	payload := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"foo": "Foo",
	})

	accessToken, err := payload.SignedString([]byte(tokenKey))
	if err != nil {
		panic(err)
	}

	result, err := ValidateJWT(accessToken, tokenKey)
	if err != nil {
		panic(err)
	}

	assert.Equal(t, result["foo"], "Foo")
}
