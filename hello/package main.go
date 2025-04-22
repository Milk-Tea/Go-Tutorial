package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main(){
	// Set properties of the predefined Logger, incl
	// the log entry prefix and a flag to disable prinitng
	// the time, source file and line number.

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	//Request a greeting message.
	message, err := greetings.Hello("")

	//If an error was returned, print it to the console and exit program.

	if err != nil {
		log.Fatal(err)
	}

	//if no error was returned then print the returned msg to console.
	fmt.Println(message)
}