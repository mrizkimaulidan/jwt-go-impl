package helper

import (
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/mrizkimaulidan/jwt-go-impl/model"
)

var SECRETKEY = []byte("secret-key")

// Custom user claims
type UserClaims struct {
	jwt.StandardClaims
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// Generate JWT for authorization
func GenerateJWT(user model.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &UserClaims{
		Id:    user.Id,
		Name:  user.Name,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(30 * time.Minute).Unix(), // add token expiration
		},
	})

	tokenString, err := token.SignedString(SECRETKEY)

	return tokenString, err
}
