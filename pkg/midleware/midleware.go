package midleware

import (
	"context"
	"github.com/dgrijalva/jwt-go"
	"nam_0515/internal/model"
	configs "nam_0515/pkg/config"
	"net/http"
	"strings"
)

type Authentication struct {
	Cfg configs.Token
}

var Auth Authentication

func init() {
	Auth = Authentication{}
}

func (a *Authentication) ValidateRoleUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := extractTokenFromHeader(r)
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims := &model.UserClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(a.Cfg.JwtSecretKey), nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Token is valid, pass the claims to the context for later use
		ctx := context.WithValue(r.Context(), "UserID", claims.UserID)
		ctx = context.WithValue(ctx, "Address", claims.Address)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func extractTokenFromHeader(r *http.Request) string {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	splitToken := strings.Split(authHeader, "Bearer ")
	if len(splitToken) != 2 {
		return ""
	}

	return splitToken[1]
}
