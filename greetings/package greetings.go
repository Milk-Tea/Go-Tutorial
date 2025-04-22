package greetings 

import (
	"errors"
	"fmt"
	"math/rand"
)

//Hello returns a greeting for the named person. 
func Hello(name string) (string, error) {
	//If no name was given, return an error with a message. 

	if name == ""{
		return "", errors.New("No name was given")
	}

	//Create a msg using a random format.
	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

//randomFormat returns on of a set of greeting msgs.
//returned msg is selected at random.

func randomFormat() string {
	//A slice of message format. 

	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Wassup big homie, %v!!",
	}

	//Return a randomly selected message format by specifying
	// a random index for the slice of formats. 

	return formats[rand.Intn(len(formats))]
}