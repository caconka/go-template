package handlers

import (
	"fmt"
	"net/http"

	"github.com/caconka/go-template/handlers/adapters"
	"github.com/caconka/go-template/services"
	"github.com/caconka/go-template/utils"
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
	utils.LogRequest(req, fmt.Sprintf("joke id %s", jokeDto.ID))
	rw.WriteHeader(http.StatusOK)

	fmt.Fprintf(rw, "%s", jokeDto.Joke)
}
