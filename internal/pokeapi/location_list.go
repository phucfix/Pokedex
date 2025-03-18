package pokeapi

import (
    "net/http"
    "io"
    "fmt"
    "encoding/json"
)

func (c *Client) ListLocation(pageURL *string) (RespShallowLocations, error) {
    // Make sure the url valid if page URL is empty for first time request
    // or for using mapback to the first page
    url := baseURL + "/location-area"
    if pageURL != nil {
        url = *pageURL
    }

    // If the data for given url in cache exist, return instead of making new request
    if val, ok := c.cache.Get(url); ok {
        locationsResp := RespShallowLocations{}
        err := json.Unmarshal(val, &locationsResp) 
        if err != nil {
            return RespShallowLocations{}, err
        }
        
        return locationsResp, nil
    }

    // Make get request
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return RespShallowLocations{}, err
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return RespShallowLocations{}, err
    }
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return RespShallowLocations{}, err
    }
    resp.Body.Close()

    if resp.StatusCode > 299 {
        return RespShallowLocations{},
               fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
    }

    locationsResp := RespShallowLocations{}
    err = json.Unmarshal(body, &locationsResp)
    if err != nil {
        return RespShallowLocations{}, err
    }

    c.cache.Add(url, body)
    return locationsResp, nil
}
