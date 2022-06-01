package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/mrizkimaulidan/jwt-go-impl/handler"
	"github.com/mrizkimaulidan/jwt-go-impl/middleware"
	"github.com/mrizkimaulidan/jwt-go-impl/model"
)

func main() {
	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(model.Response{
			Code:   http.StatusOK,
			Status: "OK",
			Data:   "API home endpoint",
		})
	})

	http.HandleFunc("/api/login", handler.Login)

	// protected handler
	http.Handle("/api/home", middleware.AuthMiddleware(http.HandlerFunc(handler.Home)))

	log.Println("server running at localhost:3000")
	log.Fatal(http.ListenAndServe("localhost:3000", nil))
}
