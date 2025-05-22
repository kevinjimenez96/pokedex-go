package pokeapi

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/kevinjimenez96/pokedex/internal/pokecache"
)

var l = LocationsClient{
	next:     "",
	previous: "",
	client: Client{
		Cache: pokecache.NewCache(time.Duration(5 * time.Second)),
	},
}

func GetNextLocations() ([]LocationArea, error) {
	if l.next == "" {
		l.next = baseUrl + locationsPath
	}

	return getLocations(l.next)
}

func GetPreviousLocations() ([]LocationArea, error) {
	if l.previous == "" {
		return nil, fmt.Errorf("you're on the first page")
	}

	return getLocations(l.previous)
}

func GetLocation(location string) (LocationResponse, error) {
	data, err := l.client.GetCached(fmt.Sprintf(baseUrl+locationPath, location))
	if err != nil {
		return LocationResponse{}, err
	}

	var locationResponse LocationResponse
	err = json.Unmarshal(data, &locationResponse)
	if err != nil {
		return LocationResponse{}, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return locationResponse, nil
}

func getLocations(url string) ([]LocationArea, error) {
	data, err := l.client.GetCached(url)

	if err != nil {
		return nil, err
	}

	var locationsResponse LocationsResponse

	err = json.Unmarshal(data, &locationsResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	l.next = locationsResponse.Next
	l.previous = locationsResponse.Previous

	return locationsResponse.Results, nil
}
