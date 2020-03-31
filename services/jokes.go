package services

import (
	"github.com/caconka/go-template/models"
)

type (
	JokeService interface {
		GetRandomJoke() (*models.Joke, error)
		GetJokeByID(id string) (*models.Joke, error)
	}

	JokeProvider interface {
		GetRandomJoke() (*models.Joke, error)
		GetJokeByID(id string) (*models.Joke, error)
	}
)

type jokeService struct {
	provider JokeProvider
}

func (j *jokeService) GetRandomJoke() (*models.Joke, error) {
	return j.provider.GetRandomJoke()
}

func (j *jokeService) GetJokeByID(id string) (*models.Joke, error) {
	return j.provider.GetJokeByID(id)
}

func NewJokeService(provider JokeProvider) JokeService {
	return &jokeService{
		provider,
	}
}
