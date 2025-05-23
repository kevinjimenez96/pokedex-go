package pokeapi

type PokemonClient struct {
	client Client
}

type LocationsClient struct {
	next     string
	previous string
	client   Client
}

type LocationsResponse struct {
	Count    int            `json:"count"`
	Next     string         `json:"next"`
	Previous string         `json:"previous"`
	Results  []LocationArea `json:"results"`
}

type LocationArea struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type LocationResponse struct {
	GameIndex         int                 `json:"game_index"`
	ID                int                 `json:"id"`
	Location          Location            `json:"location"`
	Name              string              `json:"name"`
	PokemonEncounters []PokemonEncounters `json:"pokemon_encounters"`
}

type Location struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonE struct {
	Name string `json:"name"`
	URL  string `json:"url"`
}

type PokemonEncounters struct {
	Pokemon PokemonE `json:"pokemon"`
}

type Pokemon struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	Stats          []Stat `json:"stats"`
	Types          []Type `json:"types"`
}

type Stat struct {
	Stat struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"stat"`
	BaseStat int `json:"base_stat"`
}

type Type struct {
	Type struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"type"`
}
