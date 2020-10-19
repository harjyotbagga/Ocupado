package utils

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/harjyotbagga/meeting-scheduler-api/models"
)

func InvalidRequest(w http.ResponseWriter, statCode int, message string) {
	w.Header().Set("Content-Type", "application/json")
	switch statCode {
	case 400:
		w.WriteHeader(http.StatusBadRequest)
	case 403:
		w.WriteHeader(http.StatusForbidden)
	case 404:
		w.WriteHeader(http.StatusNotFound)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
	err := models.Error{
		StatusCode:   statCode,
		ErrorMessage: message,
	}
	json.NewEncoder(w).Encode(err)
}

func ExtractParam(r *http.Request) string {
	p := strings.Split(r.URL.Path, "/")
	return p[len(p)-1]
}

func InternalError(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	err := models.Error{
		StatusCode:   500,
		ErrorMessage: "An Internal Server Error has occoured. Gotta check the logs!",
	}
	json.NewEncoder(w).Encode(err)
}
