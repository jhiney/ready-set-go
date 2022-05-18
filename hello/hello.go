package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	//set properties of the predefined Logger, including
	//the log entry prefix and a flag to disable printing
	//the time, source file, and line number
	log.SetPrefix("greedings: ")
	log.SetFlags(0)

	//request a greetings message
	message, err := greetings.Hello("")

	//If an error was returned, print it to the console and
	//exit the program

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
