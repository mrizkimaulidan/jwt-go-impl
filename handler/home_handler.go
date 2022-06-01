package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mrizkimaulidan/jwt-go-impl/helper"
	"github.com/mrizkimaulidan/jwt-go-impl/model"
)

// Home handler, it should be protected with Auth Middleware
func Home(w http.ResponseWriter, r *http.Request) {
	// if method not equal GET
	if r.Method != "GET" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(model.Response{
			Code:   http.StatusMethodNotAllowed,
			Status: "METHOD NOT ALLOWED",
		})
		return
	}

	userInformationCtx := r.Context().Value(model.ContextKeyUserInformation).(*helper.UserClaims)

	json.NewEncoder(w).Encode(model.Response{
		Code:   http.StatusOK,
		Status: "OK",
		Data: map[string]any{
			"id":        userInformationCtx.Id,
			"name":      userInformationCtx.Name,
			"email":     userInformationCtx.Email,
			"expiresAt": userInformationCtx.ExpiresAt,
		},
	})
}
