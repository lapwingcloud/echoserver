package httpechoserver

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"
)

func newLogger(logformat string) *slog.Logger {
	switch logformat {
	case "json":
		return slog.New(slog.NewJSONHandler(os.Stderr, nil))
	default:
		return slog.Default()
	}
}

func writeJSON(w http.ResponseWriter, data any) error {
	resp, err := json.Marshal(data)
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resp)
	if err != nil {
		return err
	}
	return nil
}

func writeError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)
	_, _ = w.Write([]byte(fmt.Sprintf(
		`{"timestamp":"%s","error":"%v"}`,
		time.Now().Format(time.RFC3339),
		err,
	)))
}
