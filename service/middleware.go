package service

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

// JWTMiddleware return a middleware for checking if a user is authenticated.
func JWTMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if len(tokenString) == 0 {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode("header not found")
		}
		tokenString = strings.Replace(tokenString, "Bearer ", "", 1)
		claims, err := VerifyToken(tokenString)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(&err)
		}

		userID := claims.(jwt.MapClaims)["name"].(string)
		r.Header.Set("name", userID)
		next.ServeHTTP(w, r)
	})
}

// VerifyToken verify a JWT token based on the environment variable JWT_KEY.
func VerifyToken(tokenString string) (jwt.Claims, error) {
	signingKey := []byte("HUBLOVESFOUNDBADIEZSALIMALSO19951992ECAMLABO20185MIN")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return signingKey, nil
	})

	if token.Valid && err == nil {
		return token.Claims, nil
	}
	return nil, err
}
