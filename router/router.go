package router

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"lockbin_server/types"
)

func GetRecord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(types.Record{})
}

func CreateRecord(w http.ResponseWriter, r *http.Request) {
	var record types.Record
	err := json.NewDecoder(r.Body).Decode(&record)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	slog.Info("Creating record", slog.String("uuid", record.UUID))
	// Do things

	w.WriteHeader(http.StatusCreated)
	_, err = w.Write([]byte("Record created successfully"))
	if err != nil {
		slog.Error("Can't send response", slog.String("uuid", record.UUID))
	}
}
