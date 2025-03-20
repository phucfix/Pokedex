package pokeapi

import (
    "encoding/json"
    "fmt"
    "net/http"
    "io"
)

func (c *Client) GetLocation(locationName string) (LocationDetails, error) {
    url := baseURL + "/location-area/" + locationName

    // If the data for given url in cache exist, return the data instead of making new request
    if val, ok := c.cache.Get(url); ok {
        locationDetailsResp := LocationDetails{}
        err := json.Unmarshal(val, &locationDetailsResp)
        if err != nil {
            return LocationDetails{}, err
        }

        return locationDetailsResp, nil
    }

    // Make get request if no cache for the data exist
    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        return LocationDetails{}, err
    }

    resp, err := c.httpClient.Do(req)
    if err != nil {
        return LocationDetails{}, err
    }
    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return LocationDetails{}, err
    }
    resp.Body.Close()

    if resp.StatusCode > 299 {
        return LocationDetails{},
               fmt.Errorf("Response failed with status code: %d and\nbody: %s\n", resp.StatusCode, body)
    }

    locationDetailResp := LocationDetails{}
    err = json.Unmarshal(body, &locationDetailResp)
    if err != nil {
        return LocationDetails{}, err
    }

    c.cache.Add(url, body)
    return locationDetailResp, nil
}
