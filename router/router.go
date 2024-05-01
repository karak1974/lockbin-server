package router

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"lockbin_server/database"
	"log/slog"
	"net/http"

	"lockbin_server/types"
)

// GetRecord return stored data if exist
func GetRecord(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	slog.Info("Get request", slog.String("uuid", id))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(types.Record{})
}

// CreateRecord create record in the database if every parameter is valid
func CreateRecord(w http.ResponseWriter, r *http.Request) {
	var record types.Record
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		slog.Error("Error decoding body", slog.String("error", err.Error()))
		message := types.Message{
			Status:  "failed",
			Message: "invalid request body",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(message)
		return
	}

	if err = record.Verify(); err != nil {
		slog.Error("Error verifying record", slog.String("error", err.Error()))
		message := types.Message{
			Status:  "failed",
			Message: err.Error(),
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(message)
		return
	}

	uuid, err := database.CreateRecord(record)
	if err != nil {
		slog.Error("Error creating record", slog.String("error", err.Error()))
		message := types.Message{
			Status:  "failed",
			Message: "error creating record",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(message)
		return
	}

	message := types.Message{
		Status:  "success",
		Message: uuid,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}
