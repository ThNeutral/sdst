package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/thneutral/sdst/code/server/internal/database"
	"golang.org/x/crypto/bcrypt"
)

func HandleCreateUser(queries *database.Queries) http.HandlerFunc {
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

		user, err := queries.CreateUser(r.Context(), database.CreateUserParams{
			UserID:    uuid.New(),
			Token:     uuid.New(),
			Username:  reqmodel.Username,
			Email:     reqmodel.Email,
			Password:  string(password),
			CreatedAt: time.Now(),
			LastLogin: time.Now(),
		})
		if err != nil {
			log.Println(err)
			writeError(w, "Failed to create user", http.StatusInternalServerError)
			return
		}

		var respmodel struct {
			Token string `json:"token"`
		}
		respmodel.Token = user.Token.String()
		writeResponse(w, respmodel, http.StatusCreated)
	}
}

func HandleLoginByEmail(queries *database.Queries) http.HandlerFunc {
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

		user, err := queries.GetUserByEmail(r.Context(), reqmodel.Email)
		if err != nil {
			writeError(w, "Failed to find user with given email", http.StatusNotFound)
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqmodel.Password)); err != nil {
			writeError(w, "Incorrect password", http.StatusForbidden)
			return
		}

		queries.UpdateLastLoginTime(r.Context(), database.UpdateLastLoginTimeParams{
			UserID:    user.UserID,
			LastLogin: time.Now(),
		})

		var respmodel struct {
			Token string `json:"token"`
		}
		respmodel.Token = user.Token.String()
		writeResponse(w, respmodel, http.StatusCreated)
	}
}

func HandleLoginByUsername(queries *database.Queries) http.HandlerFunc {
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

		user, err := queries.GetUserByUsername(r.Context(), reqmodel.Username)
		if err != nil {
			writeError(w, "Failed to find user with given email", http.StatusNotFound)
			return
		}

		if err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqmodel.Password)); err != nil {
			writeError(w, "Incorrect password", http.StatusForbidden)
			return
		}

		queries.UpdateLastLoginTime(r.Context(), database.UpdateLastLoginTimeParams{
			UserID:    user.UserID,
			LastLogin: time.Now(),
		})

		var respmodel struct {
			Token string `json:"token"`
		}
		respmodel.Token = user.Token.String()
		writeResponse(w, respmodel, http.StatusCreated)
	}
}
