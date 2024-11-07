package handlers

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/thneutral/sdst/code/server/internals/dummydb"
	"golang.org/x/crypto/bcrypt"
)

func HandleCreateUser(db *dummydb.DummyDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type RequestModel struct {
			Username string `json:"username"`
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		reqmodel, err := verifyModel[RequestModel](w, r)
		if err != nil {
			log.Println(err.Error())
			return
		}
		password, err := bcrypt.GenerateFromPassword([]byte(reqmodel.Password), bcrypt.DefaultCost)
		if err != nil {
			writeError(w, "Failed to hash password", http.StatusInternalServerError)
			return
		}
		token := uuid.NewString()
		db.WriteRequest <- dummydb.WriteDBRequest{
			Table: "users",
			Row: map[string]string{
				"id":       uuid.NewString(),
				"username": reqmodel.Username,
				"email":    reqmodel.Email,
				"password": string(password),
				"token":    token,
			},
		}
		writeResponse(w, struct {
			Token string `json:"token"`
		}{Token: token}, http.StatusCreated)
	}
}

func HandleLoginByEmail(db *dummydb.DummyDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type RequestModel struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		reqmodel, err := verifyModel[RequestModel](w, r)
		if err != nil {
			log.Println(err.Error())
			return
		}
		dbrespchan := make(chan []map[string]string)
		db.ReadRequest <- dummydb.ReadDBRequest{
			Table:  "users",
			Fields: []string{"email", "password", "token"},
			Data:   dbrespchan,
		}
		dbresp := <-dbrespchan

		var user map[string]string
		for _, row := range dbresp {
			if row["email"] == reqmodel.Email {
				user = row
				break
			}
		}

		if user == nil {
			writeError(w, "Did not find user with such email", http.StatusNotFound)
			return
		}

		passwordHashed := user["password"]
		token := user["token"]

		if err := bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(reqmodel.Password)); err != nil {
			writeError(w, "Invalid password", http.StatusForbidden)
			return
		}

		writeResponse(w, struct {
			Token string `json:"token"`
		}{Token: token}, http.StatusOK)
	}
}

func HandleLoginByUsername(db *dummydb.DummyDB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type RequestModel struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}
		reqmodel, err := verifyModel[RequestModel](w, r)
		if err != nil {
			log.Println(err.Error())
			return
		}
		dbrespchan := make(chan []map[string]string)
		db.ReadRequest <- dummydb.ReadDBRequest{
			Table:  "users",
			Fields: []string{"username", "password", "token"},
			Data:   dbrespchan,
		}
		dbresp := <-dbrespchan

		var user map[string]string
		for _, row := range dbresp {
			if row["username"] == reqmodel.Username {
				user = row
				break
			}
		}

		if user == nil {
			writeError(w, "Did not find user with such username", http.StatusNotFound)
			return
		}

		passwordHashed := user["password"]
		token := user["token"]

		if err := bcrypt.CompareHashAndPassword([]byte(passwordHashed), []byte(reqmodel.Password)); err != nil {
			writeError(w, "Invalid password", http.StatusForbidden)
			return
		}

		writeResponse(w, struct {
			Token string `json:"token"`
		}{Token: token}, http.StatusOK)
	}
}

func HandlePingGateway(w http.ResponseWriter, r *http.Request, user map[string]string) {
	writeResponse(w, struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}{Username: user["username"], Email: user["email"]}, http.StatusOK)
}
