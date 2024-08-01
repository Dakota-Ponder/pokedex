package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)




func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Very basic REPL, press exit to quit")
	for {

		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		fmt.Println("Echoing: ", words)

		// check if first element of the slice is the word exit 
		if words[0] == "exit" {
			commandExit()
		}
	}
}




func cleanInput(str string) []string {
	// lowercase the strings
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words

}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

// function that holds a map that returns commands 
func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name: "help",
			description: "Displays a help message",
			// callback: commandHelp,
			}, 
		"exit": {
			name: "exit",
			description: "Exit the PokeDex",
			// callback: commandExit,
		},
	}

}
