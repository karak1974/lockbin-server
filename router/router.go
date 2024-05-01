package router

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"lockbin_server/database"
	"log/slog"
	"net/http"

	"lockbin_server/types"
)

func GetRecord(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	slog.Info("Get request", slog.String("uuid", id))

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(types.Record{})
}

func CreateRecord(w http.ResponseWriter, r *http.Request) {
	var record types.Record
	var message types.Message
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		slog.Error("Error decoding body", slog.Any("error", err))
		message = types.Message{
			Status:  "failed",
			Message: "invalid request body",
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(message)
		return
	}

	slog.Info("Creating record")
	uuid, err := database.CreateRecord(record)
	if err != nil {
		slog.Error("Error creating record", slog.Any("error", err))
		message = types.Message{
			Status:  "failed",
			Message: "error creating record",
		}
	} else {
		message = types.Message{
			Status:  "success",
			Message: uuid,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(message)
}
