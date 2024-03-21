package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Very basic REPL, press exit to quit")
	fmt.Print(">> ")
	for {

		scanner.Scan()
		text := scanner.Text()
		fmt.Println("Echoing: ", text)

		if text == "exit" {
			fmt.Println("Exiting...")
			break
		}
	}

}
