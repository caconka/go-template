package adapters

import (
	"github.com/caconka/go-template/handlers/dto"
	"github.com/caconka/go-template/models"
)

func ConvertJokeToDto(joke *models.Joke) *dto.Joke {
	return &dto.Joke{
		ID:   joke.ID,
		Joke: joke.Joke,
	}
}
