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

func HealthCheck(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set(contentType, contentTypeJSON)

	if body, err := json.Marshal(&dto.Status{Value: "ok"}); err == nil {
		rw.WriteHeader(http.StatusOK)
		utils.LogRequest(req).Info("health")
		fmt.Fprintf(rw, string(body))
	} else {
		utils.LogRequest(req).Error(err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
	}

}
