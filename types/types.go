package types

import (
	"errors"
	"log/slog"
	"time"
)

type Record struct {
	MasterKey  string `json:"master_key,omitempty"`
	UnlockTime int64  `json:"unlock_time"`
	DeleteTime int64  `json:"delete_time"`
	Message    string `json:"message,omitempty"`
}

func (r Record) Verify() error {
	if len(r.MasterKey) < 4 {
		slog.Error("Too short master key", slog.Int("minimum", 4), slog.Int("given", len(r.MasterKey)))
		return errors.New("master key is too short, must be at least 4 character")
	}
	if len(r.Message) < 1 {
		slog.Error("Too short message", slog.Int("minimum", 1), slog.Int("give", len(r.MasterKey)))
		return errors.New("message is too short, must be at least 1 character")
	}
	if r.UnlockTime <= time.Now().Unix() {
		slog.Error("Unlock time can't be before now", slog.Time("now", time.Now()), slog.Time("unlock", time.Unix(r.UnlockTime, 0)))
		return errors.New("unlock time can't be before now")
	}
	if r.DeleteTime <= r.UnlockTime {
		slog.Error("Delete time can't be before unlock time", slog.Time("delete", time.Unix(r.DeleteTime, 0)), slog.Time("unlock", time.Unix(r.UnlockTime, 0)))
		return errors.New("delete time can't be before unlock time")
	}
	if r.DeleteTime <= time.Now().Unix() {
		slog.Error("Delete time can't be before now", slog.Time("now", time.Now()), slog.Time("delete", time.Unix(r.DeleteTime, 0)))
		return errors.New("delete time can't be before now")
	}
	return nil
}

type Message struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
