package controller

import (
	"context"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"strings"
	"time"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorTokens, ok := r.Header["Authorization"]
		if !ok || len(authorTokens) == 0 {
			http.Error(w, "you need to login first", http.StatusUnauthorized)
			return
		}
		jwtToken := strings.Split(authorTokens[0], "Bearer ")[1]
		token, err := jwt.Parse(jwtToken, func(tk *jwt.Token) (interface{}, error) {
			_, ok := tk.Method.(*jwt.SigningMethodHMAC)
			if !ok {
				http.Error(w, "unauthorized", http.StatusUnauthorized)
			}
			return []byte("secretkey"), nil
		})
		if err != nil {
			http.Error(w, fmt.Sprintf("unauthorized due to error %v", err), http.StatusUnauthorized)
			return
		}
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok || !token.Valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		// validate user role
		if claims["userrole"].(float64) == 1 {
			http.Error(w, "need admin role to complete this action", http.StatusUnauthorized)
			return
		}
		// validate claims
		expirationTime := claims["expired"].(float64)
		if expirationTime < float64(time.Now().UTC().Unix()) {
			http.Error(w, "session timeout, please login again", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "props", claims)

		// Do next things
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
