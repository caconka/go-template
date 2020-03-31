package handlers

import (
	"fmt"
	"net/http"

	"github.com/caconka/go-template/handlers/adapters"
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

func (h *JokeHandler) GetRandomJoke(rw http.ResponseWriter, req *http.Request) {
	joke, _ := h.Service.GetRandomJoke()

	jokeDto := adapters.ConvertJokeToDto(joke)
	utils.LogRequest(req).Info(fmt.Sprintf("joke id %s", jokeDto.ID))
	rw.WriteHeader(http.StatusOK)

	fmt.Fprintf(rw, "%s", jokeDto.Joke)
}

func (h *JokeHandler) GetJokeByID(rw http.ResponseWriter, req *http.Request) {
	jokeID := httprouter.ParamsFromContext(req.Context()).ByName(paramJokeID)

	joke, err := h.Service.GetJokeByID(jokeID)

	if err != nil {
		utils.LogRequest(req).Error(fmt.Sprintf(err.Error()))
		rw.WriteHeader(http.StatusNotFound)
	} else {
		jokeDto := adapters.ConvertJokeToDto(joke)
		utils.LogRequest(req).Info(fmt.Sprintf("joke id %s", jokeDto.ID))
		rw.WriteHeader(http.StatusOK)

		fmt.Fprintf(rw, "%s", jokeDto.Joke)
	}
}
