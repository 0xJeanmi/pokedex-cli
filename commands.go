package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"

	pokecache "github.com/0xJeanmi/pokedexcli/internal"
)

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	commands := getCommands()

	fmt.Println("Welcome to the Pokedex!")
	fmt.Print("Usage: \n\n")

	for c := range commands {
		fmt.Println(commands[c].name + ": " + commands[c].description)
	}
	return nil
}

var currentMapPage int = 0

func displayLocationAreas() error {
	params := requestParams{
		method:   "get",
		endpoint: "location-area/",
		body:     nil,
		current:  currentMapPage,
	}

	response, err := httpRequest(params)
	if err != nil {
		return fmt.Errorf("error getting location areas: %v", err)
	}

	results, ok := response["results"].([]interface{})
	if !ok || len(results) == 0 {
		return fmt.Errorf("no location areas found")
	}

	for _, el := range results {
		if location, ok := el.(map[string]interface{}); ok {
			if name, ok := location["name"].(string); ok {
				fmt.Println("🌏 " + name)
			}
		}
	}

	return nil
}

func commandMap() error {
	err := displayLocationAreas()
	if err != nil {
		fmt.Println(err)
		return err
	}

	currentMapPage += 20
	return nil
}

func commandMapb() error {
	if currentMapPage == 0 {
		fmt.Println("You are already on the first page. Use 'map' to advance.")
		return nil
	}

	currentMapPage -= 40

	err := displayLocationAreas()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}

func commandExplore(params []string) error {

	if len(params) == 0 || len(params) > 1 {
		return fmt.Errorf("incorrect arguments. Use 'explore <area>' to explore an area")
	}

	area := params[0]

	httpParams := requestParams{
		method:   "get",
		endpoint: "location-area/" + area,
		body:     nil,
		current:  0,
	}

	response, err := httpRequest(httpParams)
	if err != nil {
		return fmt.Errorf("error getting Pokemon in area: %v", err)
	}

	results, ok := response["pokemon_encounters"].([]interface{})
	if !ok || len(results) == 0 {
		return fmt.Errorf("no Pokemon found in this area")
	}

	for _, el := range results {
		if resultsMap, ok := el.(map[string]interface{}); ok {
			if pokemonsMap, ok := resultsMap["pokemon"].(map[string]interface{}); ok {
				if name, ok := pokemonsMap["name"].(string); ok {
					fmt.Println("🐾 " + name)
				}
			}
		}
	}

	return nil
}

var pokedex = pokecache.CreateNewPokedex()

func commandCatch(params []string) error {

	if len(params) == 0 || len(params) > 1 {
		return fmt.Errorf("incorrect arguments. Use 'catch <pokemon_name>' to catch a Pokemon")
	}

	pokemon := params[0]

	fmt.Printf("Throwing a Poke Ball at %s...\n", pokemon)

	httpParams := requestParams{
		method:   "get",
		endpoint: "pokemon/" + pokemon,
		body:     nil,
		current:  0,
	}

	response, err := httpRequest(httpParams)
	if err != nil {
		return fmt.Errorf("error getting Pokemon data: %v", err)
	}

	name := response["name"].(string)
	baseExperience := response["base_experience"].(float64)

	min := baseExperience / 3
	max := baseExperience

	probabilityToCatch := min + rand.Float64()*(max-min)

	if probabilityToCatch >= baseExperience/2 {
		pokedex.CapturePokemon(name, int(baseExperience))
		fmt.Println("Congratulations! You caught: " + name)
	} else {
		fmt.Println(name + " escaped! Try again!")
	}

	return nil
}

func commandInspect(params []string) error {

	if len(params) == 0 || len(params) > 1 {
		return fmt.Errorf("incorrect arguments. Use 'inspect <pokemon_name>' to show information about a Pokemon in your Pokedex")
	}

	pokemon := params[0]
	p, ok := pokedex.GetPokemon(pokemon)

	if !ok {
		fmt.Printf("%s is not in your Pokedex\n", pokemon)
	} else {
		fmt.Println("Name: " + p.Name)
		fmt.Println("Experience: " + strconv.Itoa(p.Xp))
	}
	return nil
}

func commandPokedex() error {
	pokemons, ok := pokedex.GetPokedex()

	if !ok {
		fmt.Println("Your Pokedex is empty",)
	} else {
		for _, p := range pokemons {
			fmt.Println("Name: " + p.Name)
		}
	}

	return nil
}