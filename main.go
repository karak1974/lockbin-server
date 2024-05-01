package main

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi"

	"lockbin_server/router"
)

func main() {
	r := chi.NewRouter()

	r.Get("/record/{id:[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}}", router.GetRecord)
	r.Post("/record", router.CreateRecord)

	slog.Info("Starting Lockbin server")
	if err := http.ListenAndServe(":8080", r); err != nil {
		slog.Error("Couldn't start Lockbin server", slog.String("error", err.Error()))
	}
}
