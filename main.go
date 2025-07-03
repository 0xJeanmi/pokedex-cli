package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		input := cleanInput(scanner.Text())

		if len(input) == 0 {
			fmt.Println("Incorrect command. Write 'help' to see the list of commands.")
			continue
		}

		commandName := input[0]

		params := input[1:]

		command, exists := getCommands()[commandName]

		if exists {
			err := command.callback(params)

			if err != nil {
				fmt.Println(err)
			}

			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}
