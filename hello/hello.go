package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Leonardo")
	fmt.Println(message)
}
