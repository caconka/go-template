package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/caconka/go-template/handlers/dto"
	"github.com/caconka/go-template/utils"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	utils.LogRequest(r).Info("health")
	json.NewEncoder(w).Encode(&dto.Status{Value: "ok"})
}
