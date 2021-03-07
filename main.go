package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/casbin/casbin"
	"github.com/eminetto/casbin-http-role-example/authorization"
	"github.com/eminetto/casbin-http-role-example/model"
	"github.com/eminetto/casbin-http-role-example/security"
	"github.com/gorilla/mux"
)

func main() {
	// setup casbin auth rules
	authEnforcer, err := casbin.NewEnforcerSafe("./auth_model.conf", "./policy.csv")
	if err != nil {
		log.Fatal(err)
	}

	users := createUsers()

	// setup routes
	r := mux.NewRouter()
	r.HandleFunc("/login", loginHandler(users)).Methods("POST")
	r.HandleFunc("/logout", logoutHandler())
	r.HandleFunc("/member/current", currentMemberHandler())
	r.HandleFunc("/member/role", memberRoleHandler())
	r.HandleFunc("/admin/stuff", adminHandler())
	r.Use(authorization.Authorizer(authEnforcer, users))

	log.Print("Server started on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func createUsers() model.Users {
	users := model.Users{}
	users = append(users, model.User{ID: 1, Name: "Admin", Role: "admin"})
	users = append(users, model.User{ID: 2, Name: "Sabine", Role: "member"})
	users = append(users, model.User{ID: 3, Name: "Sepp", Role: "member"})
	return users
}

func loginHandler(users model.Users) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		name := r.FormValue("name")
		user, err := users.FindByName(name)
		if err != nil {
			writeError(http.StatusBadRequest, "WRONG_CREDENTIALS", w, err)
			return
		}
		var result struct {
			Token string `json:"token"`
		}
		result.Token, err = security.NewToken(user.ID, user.Role)
		if err != nil {
			writeError(http.StatusInternalServerError, "ERROR", w, err)
			return
		}

		if err := json.NewEncoder(w).Encode(result); err != nil {
			writeError(http.StatusInternalServerError, "ERROR", w, err)
			return
		}
	})
}

func logoutHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeSuccess("SUCCESS", w)
	})
}

func currentMemberHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeSuccess("I'm an Member!", w)
	})
}

func memberRoleHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeSuccess("I'm an Member!", w)
	})
}

func adminHandler() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		writeSuccess("I'm an Admin!", w)
	})
}

func writeError(status int, message string, w http.ResponseWriter, err error) {
	log.Print("ERROR: ", err.Error())
	w.WriteHeader(status)
	w.Write([]byte(message))
}

func writeSuccess(message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}
