package main

import (
	"fmt"

	"example.com/greetings"
)

func main(){
	//Get a greeting msg and print it. 

	message := greetings.Hello("Stranger")
	fmt.Println(message)
}