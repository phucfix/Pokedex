package main

import (
    "fmt"
    "errors"
)

func commandMapf(cfg *config) error {
    locationsResp, err := cfg.pokeapiClient.ListLocation(cfg.nextLocationURL)
    if err != nil {
        return err
    }

    cfg.nextLocationURL = locationsResp.Next
    cfg.prevLocationURL = locationsResp.Previous

    for _, location := range locationsResp.Results {
        fmt.Println(location.Name)
    }

    return nil
}

func commandMapb(cfg *config) error {
    if cfg.prevLocationURL == nil {
        return errors.New("You're on the first page")
    }

    locationsResp, err := cfg.pokeapiClient.ListLocation(cfg.prevLocationURL)
    if err != nil {
        return err
    }

    cfg.nextLocationURL = locationsResp.Next
    cfg.prevLocationURL = locationsResp.Previous

    for _, location := range locationsResp.Results {
        fmt.Println(location.Name)
    }

    return nil
}
