package authorization

import (
	"errors"
	"log"
	"net/http"

	"github.com/casbin/casbin"
	"github.com/eminetto/casbin-http-role-example/model"
	"github.com/eminetto/casbin-http-role-example/roles"
	"github.com/eminetto/casbin-http-role-example/security"
)

// Authorizer is a middleware for authorization
func Authorizer(e *casbin.Enforcer, users model.Users) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			role := roles.Anonymous
			tokenString := r.Header.Get("Authorization")
			var uid int
			var err error
			if tokenString != "" {
				uid, role, err = parseToken(tokenString)
				if err != nil {
					writeError(http.StatusInternalServerError, "ERROR", w, err)
					return
				}
			}
			// check if the user still exists
			if role != roles.Anonymous {
				exists := users.Exists(uid)
				if !exists {
					writeError(http.StatusForbidden, "FORBIDDEN", w, errors.New("user does not exist"))
					return
				}
			}
			// casbin enforce
			res, err := e.EnforceSafe(string(role), r.URL.Path, r.Method)
			if err != nil {
				writeError(http.StatusInternalServerError, "ERROR", w, err)
				return
			}

			if !res {
				writeError(http.StatusForbidden, "FORBIDDEN", w, errors.New("unauthorized"))
				return
			}
			next.ServeHTTP(w, r)
		}

		return http.HandlerFunc(fn)
	}
}

func parseToken(token string) (int, roles.Role, error) {
	t, err := security.ParseToken(token)
	if err != nil {
		return 0, "", nil
	}
	tData, err := security.GetClaims(t)
	if err != nil {
		return 0, "", nil
	}
	userID := tData["userID"].(float64)
	role := tData["role"].(roles.Role)
	return int(userID), role, nil
}

func writeError(status int, message string, w http.ResponseWriter, err error) {
	log.Print("ERROR: ", err.Error())
	w.WriteHeader(status)
	w.Write([]byte(message))
}
