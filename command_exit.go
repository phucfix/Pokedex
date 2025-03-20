package main

import (
    "fmt"
    "os"
)

func commandExit(cfg *config, args ...string) error {
    if len(args) != 0 {
        fmt.Printf("Error: invalid option: %s. This command has no option\n", args[0])
        return nil
    }
    fmt.Printf("Closing the Pokedex... Goodbye!\n")
    os.Exit(0)
    return nil
}
