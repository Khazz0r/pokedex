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

func (c *Client) ListPokemon(areaName string) (PokeLocation, error) {
	url := baseURL + "/location-area/" + areaName

	// If url of location is already in cache, use the cache instead of making a new request
	if entry, exists := c.cache.Get(url); exists {
		var locationResp PokeLocation
		err := json.Unmarshal(entry, &locationResp)
		if err != nil {
			return PokeLocation{}, err
		}
		return locationResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokeLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokeLocation{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokeLocation{}, err
	}

	locationResp := PokeLocation{}
	err = json.Unmarshal(data, &locationResp)
	if err != nil {
		return PokeLocation{}, err
	}

	// Add details of location to the cache
	c.cache.Add(url, data)

	return locationResp, nil
}

func (c *Client) GetPokemon(pokemonName string) (Pokemon, error) {
	url := baseURL + "/pokemon/" + pokemonName

	// If url of pokemon is already in cache, use the cache instead of making a new request
	if entry, exists := c.cache.Get(url); exists {
		var pokemonResponse Pokemon
		err := json.Unmarshal(entry, &pokemonResponse)
		if err != nil {
			return Pokemon{}, err
		}
		return pokemonResponse, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Pokemon{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return Pokemon{}, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return Pokemon{}, err
	}

	pokemonResp := Pokemon{}
	err = json.Unmarshal(data, &pokemonResp)
	if err != nil {
		return Pokemon{}, err
	}

	// Add pokemon and its details to the cache
	c.cache.Add(url, data)

	return pokemonResp, nil
}
