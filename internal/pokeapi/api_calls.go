package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) ListLocations(pageURL *string) (PokeMap, error) {
	url := baseURL + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	
	// If url is already found in the cache, use the data in the cache over making a new request to API
	if entry, exists := c.cache.Get(url); exists {
		var locationResp PokeMap
		err := json.Unmarshal(entry, &locationResp)
		if err != nil {
			return PokeMap{}, err
		}
		return locationResp, nil
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

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokeMap{}, err
	}

	locationsResp := PokeMap{}
	err = json.Unmarshal(data, &locationsResp)
	if err != nil {
		return PokeMap{}, err
	}

	// Add url and data gotten from GET request to the cache
	c.cache.Add(url, data)

	return locationsResp, nil
}
