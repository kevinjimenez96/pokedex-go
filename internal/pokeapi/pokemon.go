package pokeapi

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/kevinjimenez96/pokedex/internal/pokecache"
)

var pokemonClient = PokemonClient{
	client: Client{
		Cache: pokecache.NewCache(time.Duration(cacheTTL)),
	},
}

func GetPokemon(name string) (Pokemon, error) {
	data, err := pokemonClient.client.GetCached(fmt.Sprintf(BaseURL+PokemonPath, name))

	if err != nil {
		return Pokemon{}, err
	}

	var pokemon Pokemon

	err = json.Unmarshal(data, &pokemon)
	if err != nil {
		return Pokemon{}, fmt.Errorf("error unmarshaling response: %w", err)
	}

	return pokemon, nil
}
