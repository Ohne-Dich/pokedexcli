package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetPokemon(pokemon string) (PokemonFull, error) {
	url := baseURL + "/pokemon/" + pokemon

	if val, ok := c.cache.Get(url); ok {
		locationsResp := PokemonFull{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return PokemonFull{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonFull{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return PokemonFull{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return PokemonFull{}, err
	}

	locationsResp := PokemonFull{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return PokemonFull{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}
