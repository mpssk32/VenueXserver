package jwt

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(
	userID string,
	role string,
) (string, error) {

	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp": time.Now().Add(
			24 * time.Hour,
		).Unix(),
	}

	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte(os.Getenv("JWT_SECRET")),
	)
}

func ValidateJWT(
	tokenString string,
) (*jwt.Token, error) {

	return jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {

			return []byte(
				os.Getenv("JWT_SECRET"),
			), nil
		},
	)
}