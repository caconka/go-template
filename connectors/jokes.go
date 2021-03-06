package connectors

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/caconka/go-template/connectors/adapters"
	tpModels "github.com/caconka/go-template/connectors/models"
	"github.com/caconka/go-template/connectors/utils"
	"github.com/caconka/go-template/models"
	"github.com/caconka/go-template/services"
)

const (
	endpoint       = "https://icanhazdadjoke.com"
	httpTimeoutSec = 60
)

type jokeProvider struct {
	httpClient  http.Client
	apiEndpoint string
}

func NewJokeProvider() services.JokeProvider {
	return &jokeProvider{
		httpClient: http.Client{
			Timeout: httpTimeoutSec * time.Second,
		},
		apiEndpoint: endpoint,
	}
}

func (provider *jokeProvider) GetRandomJoke() (res *models.Joke, err error) {

	defer func() {
		if err != nil {
			log.Error(err)
		}
	}()

	if req, err := utils.NewHTTPRequest(http.MethodGet, provider.apiEndpoint, nil, nil); err == nil {
		req.Header.Add("Accept", "application/json")
		resp, err := utils.MakeHTTPQuery(req)

		if err != nil {
			return nil, err
		}

		joke := new(tpModels.Joke)
		json.NewDecoder(resp.Body).Decode(&joke)
		return adapters.ConvertTpJokeToJoke(joke), nil
	}

	return
}

func (provider *jokeProvider) GetJokeByID(id string) (res *models.Joke, err error) {

	defer func() {
		if err != nil {
			log.Error(err)
		}
	}()

	req, err := utils.NewHTTPRequest(http.MethodGet, buildEndpointURL(provider.apiEndpoint, id), nil, nil)

	if err == nil {
		req.Header.Add("Accept", "application/json")
		resp, err := utils.MakeHTTPQuery(req)

		if err != nil {
			return nil, err
		}

		joke := new(tpModels.Joke)
		json.NewDecoder(resp.Body).Decode(&joke)

		switch joke.Status {
		case http.StatusOK, http.StatusCreated, http.StatusAccepted:
			return &models.Joke{
				ID:   joke.ID,
				Joke: joke.Joke,
			}, nil

		default:
			err := fmt.Errorf("Error code %d while retrieving joke id %s", joke.Status, id)
			return &models.Joke{ID: id, Joke: ""}, err
		}

	}

	return
}

func buildEndpointURL(basePath, id string) string {
	return fmt.Sprintf("%s/j/%s", basePath, id)
}
