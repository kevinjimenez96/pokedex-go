package pokeapi

import (
	"time"
)

// BaseURL is the root endpoint for the PokeAPI.
const BaseURL = "https://pokeapi.co/api/v2"

// API endpoint paths.
const (
	LocationsPath = "/location-area?limit=20" // List of location areas
	LocationPath  = "/location-area/%s"       // Specific location area by name or ID
	PokemonPath   = "/pokemon/%s"             // Specific Pokémon by name or ID
)

const cacheTTL = 30 * time.Second // Cache TTL for API responses
