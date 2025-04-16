package pokeapi

import (
	"encoding/json"
	"net/http"
	"io"
)

func (c *Client) ListLocations(pageURL *string) (PokeMap, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeMap{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokeMap{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokeMap{}, err
	}

	locationsResp := PokeMap{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return PokeMap{}, err
	}

	return locationsResp, nil
}
