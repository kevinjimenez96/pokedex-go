package pokeapi

import (
	"time"

	"github.com/kevinjimenez96/pokedex/internal/pokecache"
)

var l = PokemonClient{
	client: Client{
		Cache: pokecache.NewCache(time.Duration(5 * time.Second)),
	},
}
