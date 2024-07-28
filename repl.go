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

		fmt.Print(">> ")
		scanner.Scan()
		text := scanner.Text()
		fmt.Println("Echoing: ", text)

		if text == "exit" {
			fmt.Println("Exiting...")
			break
		}
	}
}

func cleanInput(str string) []string {
	// lowercase the strings
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words

}
