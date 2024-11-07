package handlers

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/thneutral/sdst/code/server/internals/dummydb"
)

type AuthenticatedHandler func(w http.ResponseWriter, r *http.Request, user map[string]string)

func Gateway(db *dummydb.DummyDB, handler AuthenticatedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" || !strings.HasPrefix(token, "Bearer ") {
			writeError(w, "Missing or invalid Authorization token", http.StatusUnauthorized)
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")
		dbrespchan := make(chan []map[string]string)
		db.ReadRequest <- dummydb.ReadDBRequest{
			Table:  "users",
			Fields: []string{"username", "password", "token"},
			Data:   dbrespchan,
		}
		dbresp := <-dbrespchan
		var user map[string]string
		for _, data := range dbresp {
			if data["token"] == token {
				user = data
				break
			}
		}
		if user == nil {
			writeError(w, "Unknown token", http.StatusForbidden)
			return
		}
		handler(w, r, user)
	}
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
