package main

import (
	"fmt"
	"log"

	"ganesh.com/greetings"
)

func main() {
	log.SetPrefix("greeting: ")
	log.SetFlags(0)

	names := []string{"ganesh", "sanjay", "lakshmi"}

	Message, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(Message)
}
