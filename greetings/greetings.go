package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty input value")
	}

	message := fmt.Sprintf(randomGreeting(), name)
	return message, nil
}

func randomGreeting() string {
	formats := []string{
		"Hi, %v. Welcome to go lang world!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
