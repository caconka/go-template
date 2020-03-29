package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/caconka/go-template/handlers/dto"
	"github.com/caconka/go-template/utils"
)

const (
	contentType     = "Content-Type"
	contentTypeJSON = "application/json"
)

func HealthCheck(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set(contentType, contentTypeJSON)

	if body, err := json.Marshal(&dto.Status{Value: "ok"}); err == nil {
		rw.WriteHeader(http.StatusOK)
		fmt.Fprintf(rw, string(body))
	} else {
		rw.WriteHeader(http.StatusInternalServerError)
	}

	logger.LogRequest(r, "health")
}
