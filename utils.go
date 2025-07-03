package main

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	pokecache "github.com/0xJeanmi/pokedexcli/internal"
)

func cleanInput(s string) []string {
	return strings.Fields(strings.ToLower(s))
}

type cliCommand struct {
	name        string
	description string
	callback    func(params []string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback: func(params []string) error {
				return commandExit()
			},
		},
		"help": {
			name:        "help",
			description: "Display a help message",
			callback: func(params []string) error {
				return commandHelp()
			},
		},
		"map": {
			name:        "map",
			description: "Display the next 20 location areas in the Pokemon world",
			callback: func(params []string) error {
				return commandMap()
			},
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous 20 location areas in the Pokemon world",
			callback: func(params []string) error {
				return commandMapb()
			},
		},
		"explore": {
			name:        "explore",
			description: "Display a list of Pokemon in a selected area",
			callback: func(params []string) error {
				return commandExplore(params)
			},
		},
		"catch": {
			name:        "catch",
			description: "Try to catch a Pokemon",
			callback: func(params []string) error {
				return commandCatch(params)
			},
		},
		"inspect": {
			name:        "inspect",
			description: "Display information about a Pokemon in your Pokedex",
			callback: func(params []string) error {
				return commandInspect(params)
			},
		},
		"pokedex": {
			name:		 "pokedex",
			description: "Display Pokemons in your Pokedex",
			callback: func(params []string) error {
				return commandPokedex()
			},
		},
	}
}

type requestParams struct {
	method   string
	endpoint string
	body     io.Reader
	current  int
}

var cache = pokecache.NewCache(5 * time.Minute)

func httpRequest(params requestParams) (map[string]interface{}, error) {
	method := params.method
	endpoint := "https://pokeapi.co/api/v2/" + params.endpoint + "?offset=" + strconv.Itoa(params.current) + "&limit=20"

	if cachedData, found := cache.Get(endpoint); found {
		var response map[string]interface{}
		err := json.Unmarshal(cachedData, &response)
		if err == nil {
			return response, nil
		}
	}

	var res *http.Response
	var err error

	if strings.ToLower(method) == "get" {
		res, err = http.Get(endpoint)
	} else {
		res, err = http.Post(endpoint, "application/json", params.body)
	}

	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := map[string]interface{}{}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}

	cache.Add(endpoint, body)

	return response, nil
}
