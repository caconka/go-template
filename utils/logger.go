package utils

import (
	"net/http"

	log "github.com/sirupsen/logrus"
)

func LogRequest(req *http.Request) *log.Entry {
	return log.WithFields(log.Fields{
		"Path":   req.URL.Path,
		"Method": req.Method,
	})
}
