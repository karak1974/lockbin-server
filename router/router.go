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
	uuid := chi.URLParam(r, "id")
	slog.Info("Get request", slog.String("uuid", uuid))

	record, err := database.GetRecord(uuid)
	if err != nil {
		slog.Error("Error getting record", slog.String("error", err.Error()))
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(types.Message{
			Status:  "failed",
			Message: "error getting record",
		})
		return
	}

	var message types.Message
	data, err := json.Marshal(record)
	if err != nil {
		slog.Error("Error marshaling record", slog.String("error", err.Error()))
		message = types.Message{
			Status:  "success",
			Message: "error marshalling record",
		}
	} else {
		message = types.Message{
			Status:  "success",
			Message: string(data),
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}

// CreateRecord create record in the database if every parameter is valid
func CreateRecord(w http.ResponseWriter, r *http.Request) {
	var record types.Record
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		slog.Error("Error decoding body", slog.String("error", err.Error()))
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(types.Message{
			Status:  "failed",
			Message: "invalid request body",
		})
		return
	}

	if err = record.Verify(); err != nil {
		slog.Error("Error verifying record", slog.String("error", err.Error()))
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(types.Message{
			Status:  "failed",
			Message: err.Error(),
		})
		return
	}

	uuid, err := database.CreateRecord(record)
	if err != nil {
		slog.Error("Error creating record", slog.String("error", err.Error()))
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(types.Message{
			Status:  "failed",
			Message: "error creating record",
		})
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(types.Message{
		Status:  "success",
		Message: uuid,
	})
}
