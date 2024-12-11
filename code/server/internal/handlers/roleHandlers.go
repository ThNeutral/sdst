package handlers

import (
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/thneutral/sdst/code/server/internal/database"
)

func HandlerUpdateUserRole(queries *database.Queries) http.HandlerFunc {
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

		user, err := queries.GetUserById(r.Context(), database.GetUserByIdParams{
			UserID:    reqmodel.UserID,
			ProjectID: reqmodel.ProjectID,
		})

		if err != nil {
			writeError(w, "User with given id is not found", http.StatusNotFound)
			return
		}

		queries.UpdateUserRole(r.Context(), database.UpdateUserRoleParams{
			Role: user.Role,
		})

	}
}
