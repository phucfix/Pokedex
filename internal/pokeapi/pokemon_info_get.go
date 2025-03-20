package pokeapi

import (
    "fmt"
    "encoding/json"
    "net/http"
    "io"
)

func (c *Client) GetPokemonInfo(pokemonName string) (Pokemon, error) {
    url := baseURL + "/pokemon/" + pokemonName

    // If data for giving url already in cache, return the data instead of making new request
    if val, ok := c.cache.Get(url); ok {
        PokemonResp := Pokemon{}
        err := json.Unmarshal(val, &PokemonResp)
        if err != nil {
            return Pokemon{}, err
        }

        return PokemonResp, nil
    }

    // Make request if no cache for the data exist
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return Pokemon{}, err
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return Pokemon{}, err
    }
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return Pokemon{}, err
    }
    resp.Body.Close()

    if resp.StatusCode > 299 {
        return Pokemon{},
               fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
    }

    PokemonResp := Pokemon{}
    err = json.Unmarshal(body, &PokemonResp)
    if err != nil {
        return Pokemon{}, err
    }

    c.cache.Add(url, body)
    return PokemonResp, nil
}
