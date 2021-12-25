package auth

import (
	"bar/configs"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func Verify(token string) (username string, err error) {
	publicKey, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(configs.JWTPublicKey))

	tok, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}

		return publicKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("validate: %w", err)
	}

	claims, ok := tok.Claims.(jwt.MapClaims)
	if !ok || !tok.Valid {
		return "", fmt.Errorf("validate: invalid")
	}

	if name, ok := claims["username"]; ok {
		if d, ok := name.(string); ok {
			return d, nil
		}
	}
	return "", fmt.Errorf("not data")
}

func VerifyFromHeader(header http.Header) (username string, err error) {
	httpAuth := header.Get("Authorization")
	if !strings.HasPrefix(httpAuth, "Bearer ") {
		return "", fmt.Errorf("token not found")
	}

	token := strings.Fields(httpAuth)[1]

	return Verify(token)
}
