package adapters

import (
	tpModels "github.com/caconka/go-template/connectors/models"
	"github.com/caconka/go-template/models"
)

func ConvertTpJokeToJoke(tpJoke *tpModels.Joke) *models.Joke {
	return &models.Joke{
		ID:   tpJoke.ID,
		Joke: tpJoke.Joke,
	}
}
