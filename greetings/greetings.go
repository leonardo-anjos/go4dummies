package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty input value")
	}

	message := fmt.Sprintf("Hi, %v. Welcome to go lang wolrd!", name)
	return message, nil
}
