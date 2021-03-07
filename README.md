# casbin-http-role-exampe

Simplistic Example of role-based HTTP Authorization with [casbin](https://github.com/casbin/casbin) using [jwt](github.com/dgrijalva/jwt-go) tokens.

Inspired by [zupzup/casbin-http-role-example](zupzup/casbin-http-role-example)

Run with

```bash
go run main.go
```

Which starts a server at `http://localhost:8080` with the following routes:

* `POST /login` - accessible if not logged in
   * takes `name` as a form-data parameter - there is no password
   * Valid Users: 
     * `Admin` ID: `1`, Role: `admin`
     * `Sabine` ID: `2`, Role: `member`
     * `Sepp` ID: `3`, Role: `member`
* `POST /logout` - accessible if logged in
* `GET /member/current` - accessible if logged in as a member
* `GET /member/role` - accessible if logged in as a member
* `GET /admin/stuff` - accessible if logged in as an admin

## Examples using cURL

### /login

```
curl -X "POST" "http://localhost:8080/login" \
     -H 'Content-Type: application/x-www-form-urlencoded' \
     --data-urlencode "name=Sabine"
``` 

### /logout

```
curl "http://localhost:8080/logout" \
     -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTUwODA1MjYsImlhdCI6MTYxNTA3Njg5NiwibmJmIjoxNjE1MDc2ODk2LCJyb2xlIjoiYWRtaW4iLCJ1c2VySUQiOjF9.6qU62-RzCpSqjsUjbFeq1oIlQGDJQpBBm2iCcqtIMwo'
```

### /member/current

```
curl "http://localhost:8080/member/current" \
     -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTUwODA1MjYsImlhdCI6MTYxNTA3Njg5NiwibmJmIjoxNjE1MDc2ODk2LCJyb2xlIjoiYWRtaW4iLCJ1c2VySUQiOjF9.6qU62-RzCpSqjsUjbFeq1oIlQGDJQpBBm2iCcqtIMwo'
```

### /member/role

```
curl "http://localhost:8080/member/role" \
     -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTUwODA1NzYsImlhdCI6MTYxNTA3Njk0NiwibmJmIjoxNjE1MDc2OTQ2LCJyb2xlIjoibWVtYmVyIiwidXNlcklEIjoyfQ.WPKsSvuBRbI7Pdv0GubJRrElcHe244bCtxDUq6nuT2w'
```

### /admin/stuff

```
curl "http://localhost:8080/admin/stuff" \
     -H 'Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MTUwODA1NzYsImlhdCI6MTYxNTA3Njk0NiwibmJmIjoxNjE1MDc2OTQ2LCJyb2xlIjoibWVtYmVyIiwidXNlcklEIjoyfQ.WPKsSvuBRbI7Pdv0GubJRrElcHe244bCtxDUq6nuT2w'
```
