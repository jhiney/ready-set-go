package greetings

import (
	"errors"
	"fmt"
)

//hello returns a greeting for the named person
func Hello(name string) (string, error) {
	//if no name was given, return an error with a message
	if name == "" {
		return "", errors.New("no name")
	}
	//If a name was received, return a value that embeds the name
	//in a message
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
