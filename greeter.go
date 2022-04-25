package main

import "fmt"

func greeter(greeting string) func(string) string {
	return func(name string) string {
		return fmt.Sprintf("%s, %s", greeting, name)
	}
}

var (
	englishGreeter = greeter("Howdy")
)
