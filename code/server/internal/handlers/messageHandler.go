package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/thneutral/sdst/code/server/internal/database"
)

func HandleCreateMessage(queries *database.Queries) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		type RequestModel struct {
			UserID uuid.UUID `json:"user_id"`
			ProjectID uuid.UUID `json:"project_id"`
			Body string `json:"body"`
		}

		reqmodel, err := verifyModel[RequestModel](w, r)
		if err != nil {
			log.Println(err.Error())
			return
		}

		_, err = queries.CreateMessage(r.Context(), database.CreateMessageParams{
			ID: uuid.New(),
			UserID:    reqmodel.UserID,
			ProjectID: reqmodel.ProjectID,
			Body: reqmodel.Body,
			PostedAt:  time.Now(),
		})
		if err != nil {
			log.Println(err)
			writeError(w, "Failed to create message", http.StatusInternalServerError)
			return
		}

		writeResponse(w, struct{}{}, http.StatusCreated)
	}
}

func HandleGetMessages(queries *database.Queries) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		projectIdStr := chi.URLParam(r, "projectId")
		projectId, err := uuid.Parse(projectIdStr)
		if err != nil {
			log.Println(err)
			writeError(w, "Failed to parse project id", http.StatusInternalServerError)
			return
		}

		messages, err := queries.GetMessagesByProject(r.Context(), projectId)
		if err != nil {
			log.Println(err)
			writeError(w, "Failed to get messages", http.StatusInternalServerError)
			return
		}

		writeResponse(w, messages, http.StatusCreated)
	}
}