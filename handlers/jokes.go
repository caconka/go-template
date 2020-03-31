package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/caconka/go-template/handlers/adapters"
	"github.com/caconka/go-template/models"
	"github.com/caconka/go-template/services"
	"github.com/caconka/go-template/utils"

	"github.com/julienschmidt/httprouter"
)

const (
	paramJokeID = "id"
)

type JokeHandler struct {
	Service services.JokeService
}

func NewJokeHandler(s services.JokeService) *JokeHandler {
	return &JokeHandler{
		Service: s,
	}
}

func (h *JokeHandler) GetRandomJoke(w http.ResponseWriter, r *http.Request) {
	if joke, err := h.Service.GetRandomJoke(); err != nil {
		utils.LogRequest(r).Error(err.Error())

		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&models.CustomError{Message: err.Error()})
	} else {
		jokeDto := adapters.ConvertJokeToDto(joke)
		utils.LogRequest(r).Info(fmt.Sprintf("joke id %s", jokeDto.ID))

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(jokeDto)
	}
}

func (h *JokeHandler) GetJokeByID(w http.ResponseWriter, r *http.Request) {
	jokeID := httprouter.ParamsFromContext(r.Context()).ByName(paramJokeID)

	joke, err := h.Service.GetJokeByID(jokeID)

	if err != nil {
		utils.LogRequest(r).Error(fmt.Sprintf(err.Error()))

		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(&models.CustomError{Message: err.Error()})
	} else {
		jokeDto := adapters.ConvertJokeToDto(joke)
		utils.LogRequest(r).Info(fmt.Sprintf("joke id %s", jokeDto.ID))

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(jokeDto)
	}
}
