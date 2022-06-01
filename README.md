# jwt-go-impl

Implementation of JWT Authorization with Go.

Clone:

```bash
$ git clone https://github.com/mrizkimaulidan/jwt-go-impl.git
```

Download required depedencies:

```bash
$ go mod download
```

Test:

```bash
$ go test -v ./test
=== RUN   TestLoginResponseHeaderStatusCodeShouldBe405
--- PASS: TestLoginResponseHeaderStatusCodeShouldBe405 (0.00s)
=== RUN   TestLoginResponseHeaderStatusNameShouldBe405MethodNotAllowed
--- PASS: TestLoginResponseHeaderStatusNameShouldBe405MethodNotAllowed (0.00s)
=== RUN   TestLoginResponseBodyStatusCodeShouldBe405
--- PASS: TestLoginResponseBodyStatusCodeShouldBe405 (0.00s)
=== RUN   TestLoginResponseBodyStatusNameShouldBeMethodNotAllowed
--- PASS: TestLoginResponseBodyStatusNameShouldBeMethodNotAllowed (0.00s)
=== RUN   TestLoginResponseHeaderStatusCodeShouldBe200
--- PASS: TestLoginResponseHeaderStatusCodeShouldBe200 (1.11s)
=== RUN   TestLoginResponseHeaderStatusNameShouldBe200OK
--- PASS: TestLoginResponseHeaderStatusNameShouldBe200OK (1.41s)
=== RUN   TestLoginResponseBodyStatusCodeShouldBe200
--- PASS: TestLoginResponseBodyStatusCodeShouldBe200 (1.09s)
=== RUN   TestLoginResponseBodyStatusNameShouldBeOK
--- PASS: TestLoginResponseBodyStatusNameShouldBeOK (1.05s)
=== RUN   TestLoginMalformedJsonBodyRequestHeaderStatusCodeShouldBe400
--- PASS: TestLoginMalformedJsonBodyRequestHeaderStatusCodeShouldBe400 (0.00s)
=== RUN   TestLoginMalformedJsonBodyRequestResponseBodyStatusCodeShould400
--- PASS: TestLoginMalformedJsonBodyRequestResponseBodyStatusCodeShould400 (0.00s)
=== RUN   TestLoginMalformedJsonBodyRequestHeaderStatusNameShouldBe400BadRequest
--- PASS: TestLoginMalformedJsonBodyRequestHeaderStatusNameShouldBe400BadRequest (0.00s)
=== RUN   TestLoginMalformedJsonBodyRequestResponseBodyStatusNameShouldBeBadRequest
--- PASS: TestLoginMalformedJsonBodyRequestResponseBodyStatusNameShouldBeBadRequest (0.00s)
=== RUN   TestLoginWrongPasswordResponseHeaderStatusCodeShouldBe401
--- PASS: TestLoginWrongPasswordResponseHeaderStatusCodeShouldBe401 (1.07s)
=== RUN   TestLoginWongPasswordResponseHeaderStatusNameShouldBe401Unauthorized
--- PASS: TestLoginWongPasswordResponseHeaderStatusNameShouldBe401Unauthorized (1.06s)
=== RUN   TestLoginWrongPasswordResponseBodyStatusCodeShouldBe401
--- PASS: TestLoginWrongPasswordResponseBodyStatusCodeShouldBe401 (1.20s)
=== RUN   TestLoginWrongPasswordResponseBodyStatusNameShouldBeUnauthorized
--- PASS: TestLoginWrongPasswordResponseBodyStatusNameShouldBeUnauthorized (1.32s)
=== RUN   TestHomeResponseHeaderStatusCodeShouldBe405
--- PASS: TestHomeResponseHeaderStatusCodeShouldBe405 (0.00s)
=== RUN   TestHomeResponseHeaderStatusNameShouldBe405MethodNotAllowed
--- PASS: TestHomeResponseHeaderStatusNameShouldBe405MethodNotAllowed (0.00s)
=== RUN   TestHomeResponseBodyStatusCodeShouldBe405
--- PASS: TestHomeResponseBodyStatusCodeShouldBe405 (0.00s)
=== RUN   TestHomeResponseBodyStatusNameShouldBeMethodNotAllowed
--- PASS: TestHomeResponseBodyStatusNameShouldBeMethodNotAllowed (0.00s)
=== RUN   TestHomeWithInvalidTokenResponseHeaderStatusCodeShouldBe400
--- PASS: TestHomeWithInvalidTokenResponseHeaderStatusCodeShouldBe400 (0.00s)
=== RUN   TestHomeWithInvalidTokenResponseHeaderStatusNameShouldBe400BadRequest
--- PASS: TestHomeWithInvalidTokenResponseHeaderStatusNameShouldBe400BadRequest (0.00s)
=== RUN   TestHomeWithInvalidTokenResponseBodyStatusCodeShouldBe400
--- PASS: TestHomeWithInvalidTokenResponseBodyStatusCodeShouldBe400 (0.00s)
=== RUN   TestHomeWithInvalidTokenResponseBodyStatusNameShouldBeBadRequest
--- PASS: TestHomeWithInvalidTokenResponseBodyStatusNameShouldBeBadRequest (0.00s)
=== RUN   TestHomeWithValidTokenResponseHeaderStatusCodeShouldBe200
--- PASS: TestHomeWithValidTokenResponseHeaderStatusCodeShouldBe200 (0.00s)
=== RUN   TestHomeWithValidTokenResponseHeaderStatusNameShouldBe200OK
--- PASS: TestHomeWithValidTokenResponseHeaderStatusNameShouldBe200OK (0.00s)
=== RUN   TestHomeWithValidTokenResponseBodyStatusCodeShouldBe200
--- PASS: TestHomeWithValidTokenResponseBodyStatusCodeShouldBe200 (0.00s)
=== RUN   TestHomeWithValidTokenResponseBodyStatusNameShouldBeOK
--- PASS: TestHomeWithValidTokenResponseBodyStatusNameShouldBeOK (0.00s)
PASS
ok      github.com/mrizkimaulidan/jwt-go-impl/test      9.848s
```

Run:

```bash
$ go run main.go
2022/06/01 12:29:23 server running at localhost:3000
```


API testing file request are provided on root folder with name `test.http` file.

User for login

```
User 1 :
Email : admin@mail.com
Password : secret

User 2 :
Email : mrizkimaulidan@mail.com
Password : secret
```
**API Home Endpoint**

**Request:**
```json
POST /api HTTP/1.1
Content-Type: application/json
```

**Response:**
```json
HTTP/1.1 200 OK
Date: Wed, 01 Jun 2022 04:41:54 GMT
Content-Length: 49
Content-Type: text/plain; charset=utf-8
Connection: close

{
  "code": 200,
  "status": "OK",
  "data": "API home endpoint"
}
```

**Login Endpoint**

**Request:**
```json
POST /api/login HTTP/1.1
Accept: application/json
Content-Type: application/json

{
    "email": "admin@mail.com",
    "password": "secret"
}
```

**Response:**
```json
HTTP/1.1 200 OK
Date: Wed, 01 Jun 2022 04:43:32 GMT
Content-Length: 281
Content-Type: text/plain; charset=utf-8
Connection: close

{
  "code": 200,
  "status": "OK",
  "data": {
    "email": "admin@mail.com",
    "id": 1,
    "name": "Administrator",
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTQwNjA0MTIsImlkIjoxLCJuYW1lIjoiQWRtaW5pc3RyYXRvciIsImVtYWlsIjoiYWRtaW5AbWFpbC5jb20ifQ.8R8pPaJ-Dz3gp1RoEJJ-939-gRfWbAZ-UaOT0v1BbY8"
  }
}
```

**Protected Home Endpoint**

**Request:**
```json
GET /api/home HTTP/1.1
Content-Type: application/json
Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2NTQwNjA0MTIsImlkIjoxLCJuYW1lIjoiQWRtaW5pc3RyYXRvciIsImVtYWlsIjoiYWRtaW5AbWFpbC5jb20ifQ.8R8pPaJ-Dz3gp1RoEJJ-939-gRfWbAZ-UaOT0v1BbY8
```

**Response:**
```json
HTTP/1.1 200 OK
Date: Wed, 01 Jun 2022 04:44:51 GMT
Content-Length: 114
Content-Type: text/plain; charset=utf-8
Connection: close

{
  "code": 200,
  "status": "OK",
  "data": {
    "email": "admin@mail.com",
    "expiresAt": 1654060412,
    "id": 1,
    "name": "Administrator"
  }
}
```