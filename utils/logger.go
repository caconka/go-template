package utils

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func LogRequest(request *http.Request, message string) {
	log.WithFields(log.Fields{
		"Path":   request.URL.Path,
		"Method": request.Method,
	}).Info(message)
}
