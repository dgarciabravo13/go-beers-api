package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	beerscli "github.com/dgarciabravo13/go-beers-api/internal"
)

const (
	defaultEndpoint = "/v2/beers"
	apiURL          = "https://api.punkapi.com"
)

type beerRepo struct {
	url string
}

// NewApiRepository to get beers
func NewApiRepository() beerscli.BeerRepo {
	return &beerRepo{url: apiURL}
}

func (b *beerRepo) GetBeers() (beers []beerscli.Beer, err error) {
	response, err := http.Get(fmt.Sprintf("%v%v", b.url, defaultEndpoint))
	if err != nil {
		return nil, err
	}

	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &beers)
	if err != nil {
		return nil, err
	}
	return
}
