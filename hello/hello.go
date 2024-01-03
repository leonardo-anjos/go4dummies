package main

import (
	"fmt"
	"log"

	"golang.org/x/example/hello/reverse"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Leonardo", "Quézia", "Léo"}

	messages, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
	fmt.Println(reverse.String("Hello"), reverse.Int(24601))
}
