package pokeapi

import (
	"fmt"
	"io"
	"net/http"

	"github.com/kevinjimenez96/pokedex/internal/pokecache"
)

type Client struct {
	Cache pokecache.Cache
}

func (c *Client) Get(url string) ([]byte, error) {
	res, err := http.Get(url)

	if err != nil {
		return nil, fmt.Errorf("error making request: %w", err)
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response: %w", err)
	}

	return data, nil
}

func (c *Client) GetCached(url string) ([]byte, error) {
 	if data, ok := c.Cache.Get(url); ok {
		return data, nil
	} else {
		data, err := c.Get(url)
		if err != nil {
			return nil, err
		}

		c.Cache.Add(url, data)

		return data, nil
	}
}
