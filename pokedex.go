package main

import (
    "github.com/phucfix/pokedexcli/internal/pokeapi"
)

func createPokedex() map[string]pokeapi.Pokemon {
    return make(map[string]pokeapi.Pokemon)
}
