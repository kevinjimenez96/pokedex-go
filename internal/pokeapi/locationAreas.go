package pokeapi

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/kevinjimenez96/pokedex/internal/pokecache"
)

var locationsClient = LocationsClient{
	next:     "",
	previous: "",
	client: Client{
		Cache: pokecache.NewCache(time.Duration(cacheTTL)),
	},
}

func GetNextLocations() ([]LocationArea, error) {
	if locationsClient.next == "" {
		locationsClient.next = BaseURL + LocationsPath
	}

	return getLocations(locationsClient.next)
}

func GetPreviousLocations() ([]LocationArea, error) {
	if locationsClient.previous == "" {
		return nil, fmt.Errorf("you're on the first page")
	}

	return getLocations(locationsClient.previous)
}

func GetLocation(location string) (LocationResponse, error) {
	data, err := locationsClient.client.GetCached(fmt.Sprintf(BaseURL+LocationPath, location))
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
	data, err := locationsClient.client.GetCached(url)

	if err != nil {
		return nil, err
	}

	var locationsResponse LocationsResponse

	err = json.Unmarshal(data, &locationsResponse)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling response: %w", err)
	}

	locationsClient.next = locationsResponse.Next
	locationsClient.previous = locationsResponse.Previous

	return locationsResponse.Results, nil
}
