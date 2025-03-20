package main

import (
    "fmt"
    "errors"
)

func commandMapf(cfg *config, args ...string) error {
    if len(args) != 0 {
        fmt.Printf("Error: invalid option: %s. This command has no option\n", args[0])
        return nil
    }

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

func commandMapb(cfg *config, args ...string) error {
    if len(args) != 0 {
        fmt.Printf("Error: invalid option: %s. This command has no option\n", args[0])
        return nil
    }

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
