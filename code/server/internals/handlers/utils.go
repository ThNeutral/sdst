package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"reflect"

	"github.com/google/uuid"
	"github.com/thneutral/sdst/code/server/internal/database"
)

type AuthenticatedHandler func(w http.ResponseWriter, r *http.Request, user database.User)

func Gateway(queries *database.Queries, handler AuthenticatedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			writeError(w, "Missing Authorization header", http.StatusUnauthorized)
			return
		}

		tokenUUID, err := uuid.Parse(token)
		if err != nil {
			writeError(w, "Invalid token", http.StatusBadRequest)
			return
		}

		user, err := queries.GetUserByToken(r.Context(), tokenUUID)
		if err != nil {
			writeError(w, "Failed to find user with given token", http.StatusNotFound)
			return
		}

		handler(w, r, user)
	}
}

func HandlePingGateway(w http.ResponseWriter, r *http.Request, user database.User) {
	var respmodel struct {
		Username string `json:"username"`
		Email    string `json:"email"`
	}
	respmodel.Username = user.Username
	respmodel.Email = user.Email
	writeResponse(w, respmodel, http.StatusOK)
}

func writeResponse(w http.ResponseWriter, payload interface{}, code int) {
	w.WriteHeader(code)
	if payload == nil {
		return
	}
	bytes, err := json.Marshal(payload)
	if err != nil {
		log.Println(err)
		return
	}
	w.Write(bytes)
}

func writeError(w http.ResponseWriter, msg string, code int) {
	type Model struct {
		Message string `json:"message"`
	}
	var m Model
	m.Message = msg
	bytes, _ := json.Marshal(m)
	w.WriteHeader(code)
	w.Write(bytes)
}

func verifyModel[T any](w http.ResponseWriter, r *http.Request) (T, error) {
	var payload T
	data, _ := io.ReadAll(r.Body)
	err := json.Unmarshal(data, &payload)
	if err != nil {
		log.Println(err)
		writeError(w, "Failed to unmarshal fileds", http.StatusBadRequest)
		return payload, err
	}
	val := reflect.ValueOf(payload)
	typ := reflect.TypeOf(payload)
	isError := false
	msg := "Failed to find field(s):"
	for i := 0; i < val.NumField(); i++ {
		fieldValue := val.Field(i)
		if fieldValue.IsZero() && typ.Field(i).Tag.Get("optional") != "true" {
			fieldName := typ.Field(i).Tag.Get("json")
			msg += "\t" + fieldName
			isError = true
		}
	}
	if isError {
		writeError(w, msg, http.StatusBadRequest)
		return payload, errors.New("failed to unmarhsal fields")
	}
	return payload, nil
}
