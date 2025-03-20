package main

import (
    "fmt"
    "errors"
)

func commandExplore(cfg *config, args ...string) error {
    if len(args) != 1 {
        return errors.New("you must provide a location name")
    }

    name := args[0]
    locationDetailResp, err := cfg.pokeapiClient.GetLocation(name)
    if err != nil {
        return err
    }

    fmt.Printf("Exploring %s...\n", locationDetailResp.Location.Name)
    fmt.Print("Found pokemon:\n");
    for _, pokemonEncounter := range locationDetailResp.PokemonEncounters {
        fmt.Printf("\t- %s\n", pokemonEncounter.Pokemon.Name)
    }

    return nil
}
