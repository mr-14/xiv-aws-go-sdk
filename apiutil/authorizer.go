package apiutil

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/mr-14/xiv-aws-go-sdk/errorutil"
)

// ValidateJWT validates Json web token
func ValidateJWT(tokenString string, tokenKey string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(tokenKey), nil
	})
	if err == nil {
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			return claims, nil
		}
	}

	return nil, &errorutil.HTTPError{
		Status: http.StatusUnauthorized,
		Form:   &errorutil.FormError{Message: "error.jwt.invalid"},
	}
}
