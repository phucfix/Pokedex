package main

import (
    "math/rand"
    "fmt"
    "errors"
)

func commandCatch(cfg *config, args ...string) error {
    if len(args) != 1 {
        return errors.New("you must provide a pokemon name")
    }

    pokemonName := args[0]

    // Check if pokemon already in dex
    if _, ok := cfg.caughtPokemon[pokemonName]; ok {
        fmt.Printf("You already have %s in your dex\n", pokemonName)
    }

    pokemonInfoResp, err := cfg.pokeapiClient.GetPokemonInfo(pokemonName)
    if err != nil {
        return err
    }
    
    
    result := rand.Intn(pokemonInfoResp.BaseExperience)

    fmt.Printf("Throwing a Pokeball at %s...\n", pokemonName)
    if result > 40{
        fmt.Printf("%s escaped!\n", pokemonName)
        return nil
    }
    fmt.Printf("%s was caught!\n", pokemonName)
    fmt.Println("You may now inspect it with the inspect command.")

    // Add pokemon to dex
    cfg.caughtPokemon[pokemonName] = pokemonInfoResp

    return nil
}
