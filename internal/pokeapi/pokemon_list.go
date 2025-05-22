package pokeapi

import (
	"encoding/json"
	"io"
	"net/http"
)

func (c *Client) GetLocations(area string) (RespFullLocation, error) {
	url := baseURL + "/location-area/" + area

	if val, ok := c.cache.Get(url); ok {
		locationsResp := RespFullLocation{}
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespFullLocation{}, err
		}

		return locationsResp, nil
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return RespFullLocation{}, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return RespFullLocation{}, err
	}
	defer resp.Body.Close()

	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RespFullLocation{}, err
	}

	locationsResp := RespFullLocation{}
	err = json.Unmarshal(dat, &locationsResp)
	if err != nil {
		return RespFullLocation{}, err
	}

	c.cache.Add(url, dat)
	return locationsResp, nil
}
