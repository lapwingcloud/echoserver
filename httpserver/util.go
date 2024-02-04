package httpserver

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func writeJSON(w http.ResponseWriter, data any) error {
	resp, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	resp = append(resp, byte('\n'))
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
		time.Now().Format(time.RFC3339Nano),
		err,
	)))
}
