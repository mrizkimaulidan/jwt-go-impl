package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mrizkimaulidan/jwt-go-impl/helper"
	"github.com/mrizkimaulidan/jwt-go-impl/model"
)

// Login handler for REST API
func Login(w http.ResponseWriter, r *http.Request) {
	// if method not equal POST
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(model.Response{
			Code:   http.StatusMethodNotAllowed,
			Status: "METHOD NOT ALLOWED",
		})
		return
	}

	var request map[string]any
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(model.Response{
			Code:   http.StatusBadRequest,
			Status: "BAD REQUEST",
			Data:   "malformed json body",
		})
		return
	}

	for _, user := range model.Users {
		// check email first
		if user.Email == request["email"] {
			// if email found, check the password
			if helper.CheckPassword(user.Password, request["password"].(string)) {
				// generate JWT for authorization
				token, err := helper.GenerateJWT(user)
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					json.NewEncoder(w).Encode(model.Response{
						Code:   http.StatusInternalServerError,
						Status: "INTERNAL SERVER ERROR",
						Data:   err.Error(),
					})
					return
				}

				json.NewEncoder(w).Encode(model.Response{
					Code:   http.StatusOK,
					Status: "OK",
					Data: map[string]any{
						"id":    user.Id,
						"name":  user.Name,
						"email": user.Email,
						"token": token,
					},
				})
				return
			} else {
				// password are wrong
				w.WriteHeader(http.StatusUnauthorized)
				json.NewEncoder(w).Encode(model.Response{
					Code:   http.StatusUnauthorized,
					Status: "UNAUTHORIZED",
					Data:   "password are wrong",
				})
				return
			}
		}
	}
}
