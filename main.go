package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Author: Jake Smith"
	fmt.Println("App: Example")
	fmt.Println("Version: 1.0.0")

	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')

	// TODO: Allow a user to select a different greeting
	fmt.Println(englishGreeter(name))
}
