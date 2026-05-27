package jwt

import (
	"context"
	"net/http"
	

	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(next http.Handler) http.Handler {

	return http.HandlerFunc(func(
		w http.ResponseWriter,
		r *http.Request,
	) {

		authHeader := r.Header.Get(
			"Authorization",
		)

		if authHeader == "" {

			http.Error(
				w,
				"missing token",
				http.StatusUnauthorized,
			)

			return
		}

		tokenString := authHeader[len("Bearer "):]

		token, err := ValidateJWT(
			tokenString,
		)

		if err != nil || !token.Valid {

			http.Error(
				w,
				"invalid token",
				http.StatusUnauthorized,
			)

			return
		}

		claims := token.Claims.(jwt.MapClaims)

		ctx := context.WithValue(
			r.Context(),
			"user_id",
			claims["user_id"],
		)

		ctx = context.WithValue(
			ctx,
			"role",
			claims["role"],
		)

		next.ServeHTTP(
			w,
			r.WithContext(ctx),
		)
	})
}
