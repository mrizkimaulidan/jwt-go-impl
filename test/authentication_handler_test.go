package test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mrizkimaulidan/jwt-go-impl/handler"
)

func TestLoginResponseHeaderStatusCodeShouldBe405(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/login", nil)

	handler.Login(recorder, request)

	r := recorder.Result()

	expected := 405
	if r.StatusCode != expected {
		t.Fatalf("response header status code got %d, want %d", recorder.Code, expected)
	}
}

func TestLoginResponseHeaderStatusNameShouldBe405MethodNotAllowed(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/login", nil)

	handler.Login(recorder, request)

	r := recorder.Result()

	expected := "405 Method Not Allowed"
	if r.Status != expected {
		t.Fatalf("response header status name got %s, want %s", recorder.Result().Status, expected)
	}
}

func TestLoginResponseBodyStatusCodeShouldBe405(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/login", nil)

	handler.Login(recorder, request)

	r := recorder.Result()

	var responseBody map[string]any
	json.NewDecoder(r.Body).Decode(&responseBody)

	expected := 405
	intResponseBodyCode := int(responseBody["code"].(float64))
	if intResponseBodyCode != expected {
		t.Fatalf("response body status code got %d, want %d", intResponseBodyCode, expected)
	}
}

func TestLoginResponseBodyStatusNameShouldBeMethodNotAllowed(t *testing.T) {
	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:3000/api/login", nil)

	handler.Login(recorder, request)

	r := recorder.Result()

	var responseBody map[string]any
	json.NewDecoder(r.Body).Decode(&responseBody)

	expected := "METHOD NOT ALLOWED"
	if responseBody["status"] != expected {
		t.Fatalf("response body status got %s, want %s", responseBody["status"], expected)
	}
}

func TestLoginResponseHeaderStatusCodeShouldBe200(t *testing.T) {
	recorder := httptest.NewRecorder()

	requestBody := `{"email":"admin@mail.com","password":"secret"}`
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/login", strings.NewReader(requestBody))

	handler.Login(recorder, request)

	resp := recorder.Result()

	expected := 200
	if resp.StatusCode != expected {
		t.Fatalf("response header status code got %d, want %d", resp.StatusCode, expected)
	}
}

func TestLoginResponseHeaderStatusNameShouldBe200OK(t *testing.T) {
	recorder := httptest.NewRecorder()

	requestBody := `{"email":"admin@mail.com","password":"secret"}`
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/login", strings.NewReader(requestBody))

	handler.Login(recorder, request)

	resp := recorder.Result()

	expected := "200 OK"
	if resp.Status != expected {
		t.Fatalf("response header status name got %s, want %s", resp.Status, expected)
	}
}

func TestLoginResponseBodyStatusCodeShouldBe200(t *testing.T) {
	recorder := httptest.NewRecorder()

	requestBody := `{"email":"admin@mail.com","password":"secret"}`
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/login", strings.NewReader(requestBody))

	handler.Login(recorder, request)

	resp := recorder.Result()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fail()
	}

	var responseBody map[string]any
	json.Unmarshal(bytes, &responseBody)

	expected := 200
	intResponseBodyCode := int(responseBody["code"].(float64))
	if intResponseBodyCode != expected {
		t.Fatalf("response body status code got %d, want %d", intResponseBodyCode, expected)
	}
}

func TestLoginResponseBodyStatusNameShouldBeOK(t *testing.T) {
	recorder := httptest.NewRecorder()

	requestBody := `{"email":"admin@mail.com","password":"secret"}`
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/login", strings.NewReader(requestBody))

	handler.Login(recorder, request)

	resp := recorder.Result()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fail()
	}

	var responseBody map[string]any
	json.Unmarshal(bytes, &responseBody)

	expected := "OK"
	if responseBody["status"] != expected {
		t.Fatalf("response body status code got %s, want %s", responseBody["status"], expected)
	}
}

func TestLoginMalformedJsonBodyRequestHeaderStatusCodeShouldBe400(t *testing.T) {
	recorder := httptest.NewRecorder()

	// malformed the json body
	requestBody := `{"email":"admin@mail.com","password":"secret",}`
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/login", strings.NewReader(requestBody))

	handler.Login(recorder, request)

	resp := recorder.Result()

	expected := 400
	if resp.StatusCode != expected {
		t.Fatalf("response header status code got %d, want %d", resp.StatusCode, expected)
	}
}

func TestLoginMalformedJsonBodyRequestResponseBodyStatusCodeShould400(t *testing.T) {
	recorder := httptest.NewRecorder()

	// malformed the json body
	requestBody := `{"email":"admin@mail.com","password":"secret",}`
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/login", strings.NewReader(requestBody))

	handler.Login(recorder, request)

	resp := recorder.Result()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fail()
	}

	var responseBody map[string]any
	json.Unmarshal(bytes, &responseBody)

	expected := 400
	intResponseBodyCode := int(responseBody["code"].(float64))
	if intResponseBodyCode != expected {
		t.Fatalf("response body status code got %d, want %d", intResponseBodyCode, expected)
	}
}

func TestLoginMalformedJsonBodyRequestHeaderStatusNameShouldBe400BadRequest(t *testing.T) {
	recorder := httptest.NewRecorder()

	// malformed the json body
	requestBody := `{"email":"admin@mail.com","password":"secret",}`
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/login", strings.NewReader(requestBody))

	handler.Login(recorder, request)

	resp := recorder.Result()
	expected := "400 Bad Request"
	if resp.Status != expected {
		t.Fatalf("response header status name got %s, want %s", resp.Status, expected)
	}
}

func TestLoginMalformedJsonBodyRequestResponseBodyStatusNameShouldBeBadRequest(t *testing.T) {
	recorder := httptest.NewRecorder()

	// malformed the json body
	requestBody := `{"email":"admin@mail.com","password":"secret",}`
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/login", strings.NewReader(requestBody))

	handler.Login(recorder, request)

	resp := recorder.Result()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fail()
	}

	var responseBody map[string]any
	json.Unmarshal(bytes, &responseBody)

	expected := "BAD REQUEST"
	if responseBody["status"] != expected {
		t.Fatalf("response body status name got %s, want %s", responseBody["status"], expected)
	}
}

func TestLoginWrongPasswordResponseHeaderStatusCodeShouldBe401(t *testing.T) {
	recorder := httptest.NewRecorder()

	// the password are wrong
	requestBody := `{"email":"admin@mail.com","password":"wrongpassword"}`
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/login", strings.NewReader(requestBody))

	handler.Login(recorder, request)

	resp := recorder.Result()

	expected := 401
	if resp.StatusCode != expected {
		t.Fatalf("response body status name got %d, want %d", resp.StatusCode, expected)
	}
}

func TestLoginWongPasswordResponseHeaderStatusNameShouldBe401Unauthorized(t *testing.T) {
	recorder := httptest.NewRecorder()

	// the password are wrong
	requestBody := `{"email":"admin@mail.com","password":"wrongpassword"}`
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/login", strings.NewReader(requestBody))

	handler.Login(recorder, request)

	resp := recorder.Result()

	expected := "401 Unauthorized"
	if resp.Status != expected {
		t.Fatalf("response body status name got %s, want %s", resp.Status, expected)
	}
}

func TestLoginWrongPasswordResponseBodyStatusCodeShouldBe401(t *testing.T) {
	recorder := httptest.NewRecorder()

	// the password are wrong
	requestBody := `{"email":"admin@mail.com","password":"wrongpassword"}`
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/login", strings.NewReader(requestBody))

	handler.Login(recorder, request)

	resp := recorder.Result()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fail()
	}

	var responseBody map[string]any
	json.Unmarshal(bytes, &responseBody)

	expected := 401
	intResponseBodyCode := int(responseBody["code"].(float64))
	if intResponseBodyCode != expected {
		t.Fatalf("response body status name got %d, want %d", intResponseBodyCode, expected)
	}
}

func TestLoginWrongPasswordResponseBodyStatusNameShouldBeUnauthorized(t *testing.T) {
	recorder := httptest.NewRecorder()

	// the password are wrong
	requestBody := `{"email":"admin@mail.com","password":"wrongpassword"}`
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/login", strings.NewReader(requestBody))

	handler.Login(recorder, request)

	resp := recorder.Result()

	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fail()
	}

	var responseBody map[string]any
	json.Unmarshal(bytes, &responseBody)

	expected := "UNAUTHORIZED"
	if responseBody["status"] != expected {
		t.Fatalf("response body status name got %s, want %s", responseBody["status"], expected)
	}
}
