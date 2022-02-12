package libraries

import (
	"errors"
	"ms-fetch/domain/constants/messages"

	"github.com/golang-jwt/jwt"
)

type JWTLibrary struct {
	SecretKey string
}

func NewJWTLibrary(secretKey string) JWTLibrary {
	return JWTLibrary{
		SecretKey: secretKey,
	}
}

func (lib JWTLibrary) ValidateToken(encodedToken string) (jwt.MapClaims, bool) {

	// parse token
	token, err := jwt.Parse(encodedToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return encodedToken, errors.New(messages.TokenIsNotValidMessage)
		}
		return []byte(lib.SecretKey), nil
	})
	if err != nil {
		return nil, false
	}

	// get payload
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, true
	}

	return nil, false
}
