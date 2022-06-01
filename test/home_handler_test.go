package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/mrizkimaulidan/jwt-go-impl/handler"
	"github.com/mrizkimaulidan/jwt-go-impl/helper"
	"github.com/mrizkimaulidan/jwt-go-impl/middleware"
	"github.com/mrizkimaulidan/jwt-go-impl/model"
)

func TestHomeResponseHeaderStatusCodeShouldBe405(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/home", nil)

	handler.Home(recorder, request)

	resp := recorder.Result()

	expected := 405
	if resp.StatusCode != expected {
		t.Fatalf("response header status code got %d, want %d", resp.StatusCode, expected)
	}
}

func TestHomeResponseHeaderStatusNameShouldBe405MethodNotAllowed(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/home", nil)

	handler.Home(recorder, request)

	resp := recorder.Result()

	expected := "405 Method Not Allowed"
	if resp.Status != expected {
		t.Fatalf("response header status code got %s, want %s", resp.Status, expected)
	}
}

func TestHomeResponseBodyStatusCodeShouldBe405(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/home", nil)

	handler.Home(recorder, request)

	resp := recorder.Result()

	var responseBody map[string]any
	json.NewDecoder(resp.Body).Decode(&responseBody)

	expected := 405
	intResponseBodyCode := int(responseBody["code"].(float64))
	if intResponseBodyCode != expected {
		t.Fatalf("response body status code got %d, want %d", intResponseBodyCode, expected)
	}
}

func TestHomeResponseBodyStatusNameShouldBeMethodNotAllowed(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/home", nil)

	handler.Home(recorder, request)

	resp := recorder.Result()

	var responseBody map[string]any
	json.NewDecoder(resp.Body).Decode(&responseBody)

	expected := "METHOD NOT ALLOWED"
	if responseBody["status"] != expected {
		t.Fatalf("response body status name got %s, want %s", responseBody["status"], expected)
	}
}

func TestHomeWithInvalidTokenResponseHeaderStatusCodeShouldBe400(t *testing.T) {
	// invalid token make request failed
	tokenHeader := fmt.Sprintf("Bearer %s", "invalid token lol")
	homeHandler := func(w http.ResponseWriter, r *http.Request) {}

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/home", nil)
	request.Header.Set("Authorization", tokenHeader)

	recorder := httptest.NewRecorder()

	homeHandler(recorder, request)

	// using the auth middleware when hitting the endpoint
	next := middleware.AuthMiddleware(http.HandlerFunc(handler.Home))
	next.ServeHTTP(recorder, request)

	resp := recorder.Result()

	expected := 400
	if resp.StatusCode != expected {
		t.Fatalf("response header status code got %d, want %d", resp.StatusCode, expected)
	}
}

func TestHomeWithInvalidTokenResponseHeaderStatusNameShouldBe400BadRequest(t *testing.T) {
	// invalid token make request failed
	tokenHeader := fmt.Sprintf("Bearer %s", "invalid token lol")
	homeHandler := func(w http.ResponseWriter, r *http.Request) {}

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/home", nil)
	request.Header.Set("Authorization", tokenHeader)

	recorder := httptest.NewRecorder()

	homeHandler(recorder, request)

	// using the auth middleware when hitting the home endpoint
	next := middleware.AuthMiddleware(http.HandlerFunc(handler.Home))
	next.ServeHTTP(recorder, request)

	resp := recorder.Result()

	expected := "400 Bad Request"
	if resp.Status != expected {
		t.Fatalf("response header status name got %s, want %s", resp.Status, expected)
	}
}

func TestHomeWithInvalidTokenResponseBodyStatusCodeShouldBe400(t *testing.T) {
	// invalid token make request failed
	tokenHeader := fmt.Sprintf("Bearer %s", "invalid token lol")
	homeHandler := func(w http.ResponseWriter, r *http.Request) {}

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/home", nil)
	request.Header.Set("Authorization", tokenHeader)

	recorder := httptest.NewRecorder()

	homeHandler(recorder, request)

	// using the auth middleware when hitting the home endpoint
	next := middleware.AuthMiddleware(http.HandlerFunc(handler.Home))
	next.ServeHTTP(recorder, request)

	resp := recorder.Result()

	var responseBody map[string]any
	json.NewDecoder(resp.Body).Decode(&responseBody)

	expected := 400
	intResponseBodyCode := int(responseBody["code"].(float64))
	if intResponseBodyCode != expected {
		t.Fatalf("response body status code got %d, want %d", intResponseBodyCode, expected)
	}
}

func TestHomeWithInvalidTokenResponseBodyStatusNameShouldBeBadRequest(t *testing.T) {
	// invalid token make request failed
	tokenHeader := fmt.Sprintf("Bearer %s", "invalid token lol")
	homeHandler := func(w http.ResponseWriter, r *http.Request) {}

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/home", nil)
	request.Header.Set("Authorization", tokenHeader)

	recorder := httptest.NewRecorder()

	homeHandler(recorder, request)

	// using the auth middleware when hitting the home endpoint
	next := middleware.AuthMiddleware(http.HandlerFunc(handler.Home))
	next.ServeHTTP(recorder, request)

	resp := recorder.Result()

	var responseBody map[string]any
	json.NewDecoder(resp.Body).Decode(&responseBody)

	expected := "BAD REQUEST"
	if responseBody["status"] != expected {
		t.Fatalf("response body status name got %s, want %s", responseBody["status"], expected)
	}
}

func TestHomeWithValidTokenResponseHeaderStatusCodeShouldBe200(t *testing.T) {
	user := model.Users[0]
	//actual token
	token, err := helper.GenerateJWT(user)
	if err != nil {
		t.Fatal(err)
	}

	tokenHeader := fmt.Sprintf("Bearer %s", token)
	homeHandler := func(w http.ResponseWriter, r *http.Request) {}

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/home", nil)
	request.Header.Set("Authorization", tokenHeader)

	recorder := httptest.NewRecorder()

	homeHandler(recorder, request)

	// using the auth middleware when hitting the home endpoint
	next := middleware.AuthMiddleware(http.HandlerFunc(handler.Home))
	next.ServeHTTP(recorder, request)

	resp := recorder.Result()

	var responseBody map[string]any
	json.NewDecoder(resp.Body).Decode(&responseBody)

	expected := 200
	intResponseBodyCode := int(responseBody["code"].(float64))
	if intResponseBodyCode != expected {
		t.Fatalf("response header status code got %d, want %d", intResponseBodyCode, expected)
	}
}

func TestHomeWithValidTokenResponseHeaderStatusNameShouldBe200OK(t *testing.T) {
	user := model.Users[0]
	//actual token
	token, err := helper.GenerateJWT(user)
	if err != nil {
		t.Fatal(err)
	}

	tokenHeader := fmt.Sprintf("Bearer %s", token)
	homeHandler := func(w http.ResponseWriter, r *http.Request) {}

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/home", nil)
	request.Header.Set("Authorization", tokenHeader)

	recorder := httptest.NewRecorder()

	homeHandler(recorder, request)

	// using the auth middleware when hitting the home endpoint
	next := middleware.AuthMiddleware(http.HandlerFunc(handler.Home))
	next.ServeHTTP(recorder, request)

	resp := recorder.Result()

	var responseBody map[string]any
	json.NewDecoder(resp.Body).Decode(&responseBody)

	expected := "200 OK"
	if resp.Status != expected {
		t.Fatalf("response header status name got %s, want %s", responseBody["status"], expected)
	}
}

func TestHomeWithValidTokenResponseBodyStatusCodeShouldBe200(t *testing.T) {
	user := model.Users[0]
	//actual token
	token, err := helper.GenerateJWT(user)
	if err != nil {
		t.Fatal(err)
	}

	tokenHeader := fmt.Sprintf("Bearer %s", token)
	homeHandler := func(w http.ResponseWriter, r *http.Request) {}

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/home", nil)
	request.Header.Set("Authorization", tokenHeader)

	recorder := httptest.NewRecorder()

	homeHandler(recorder, request)

	// using the auth middleware when hitting the home endpoint
	next := middleware.AuthMiddleware(http.HandlerFunc(handler.Home))
	next.ServeHTTP(recorder, request)

	resp := recorder.Result()

	var responseBody map[string]any
	json.NewDecoder(resp.Body).Decode(&responseBody)

	expected := 200
	intResponseBodyCode := int(responseBody["code"].(float64))
	if intResponseBodyCode != expected {
		t.Fatalf("response body status name got %d, want %d", intResponseBodyCode, expected)
	}
}

func TestHomeWithValidTokenResponseBodyStatusNameShouldBeOK(t *testing.T) {
	user := model.Users[0]
	//actual token
	token, err := helper.GenerateJWT(user)
	if err != nil {
		t.Fatal(err)
	}

	tokenHeader := fmt.Sprintf("Bearer %s", token)
	homeHandler := func(w http.ResponseWriter, r *http.Request) {}

	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/home", nil)
	request.Header.Set("Authorization", tokenHeader)

	recorder := httptest.NewRecorder()

	homeHandler(recorder, request)

	// using the auth middleware when hitting the home endpoint
	next := middleware.AuthMiddleware(http.HandlerFunc(handler.Home))
	next.ServeHTTP(recorder, request)

	resp := recorder.Result()

	var responseBody map[string]any
	json.NewDecoder(resp.Body).Decode(&responseBody)

	expected := "OK"
	if responseBody["status"] != expected {
		t.Fatalf("response body status name got %s, want %s", responseBody["status"], expected)
	}
}
