package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
	"github.com/mrizkimaulidan/jwt-go-impl/helper"
	"github.com/mrizkimaulidan/jwt-go-impl/model"
)

// JWT authentication middleware. Checking the token on header
// is valid or not. Send user information on request context
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")

		if !strings.Contains(authorizationHeader, "Bearer") {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Response{
				Code:   http.StatusBadRequest,
				Status: "BAD REQUEST",
				Data:   "invalid token",
			})
			return
		}

		// get the jwt token on header
		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)

		token, err := jwt.ParseWithClaims(tokenString, &helper.UserClaims{}, func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("signing method failed")
			}

			return helper.SECRETKEY, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(model.Response{
				Code:   http.StatusBadRequest,
				Status: "BAD REQUEST",
				Data:   err.Error(),
			})
			return
		}

		if claims, ok := token.Claims.(*helper.UserClaims); ok && token.Valid {
			// create context on http request
			ctx := context.WithValue(context.Background(), model.ContextKeyUserInformation, claims)
			r = r.WithContext(ctx)

			next.ServeHTTP(w, r)
		}
	})
}
