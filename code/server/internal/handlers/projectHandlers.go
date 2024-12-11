package handlers

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/thneutral/sdst/code/server/internal/database"
)

func HandleAddUserToProject(queries *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type RequestModel struct {
			UserID    uuid.UUID `json:"user_id"`
			ProjectID uuid.UUID `json:"project_id"`
			Role      string    `json:"role"`
		}

		reqmodel, err := verifyModel[RequestModel](w, r)
		if err != nil {
			log.Println(err.Error())
			return
		}

		queries.AddUserToProject(r.Context(), database.AddUserToProjectParams{
			ProjectID: reqmodel.ProjectID,
			UserID:    reqmodel.UserID,
			Role:      reqmodel.Role,
		})

		writeResponse(w, nil, http.StatusCreated)
	}
}

func HandlerDeleteUserFromProject(queries *database.Queries) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type RequestModel struct {
			UserID    uuid.UUID `json:"user_id"`
			ProjectID uuid.UUID `json:"project_id"`
		}

		reqmodel, err := verifyModel[RequestModel](w, r)
		if err != nil {
			log.Println(err.Error())
			return
		}

		user, err := queries.GetUserById(r.Context(), database.GetUserByIdParams{
			ProjectID: reqmodel.ProjectID,
			UserID:    reqmodel.UserID,
		})

		if user.Role != "Admin" {
			writeError(w, "You do not have enough rights", http.StatusForbidden)
			return
		}

		if err != nil {
			writeError(w, "Failed to find user with given ID", http.StatusNotFound)
			return
		}

		writeResponse(w, nil, http.StatusNoContent)
	}
}
