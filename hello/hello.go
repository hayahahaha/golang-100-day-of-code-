package main

import (
	"example.com/greatings"
	"fmt"
	"log"
)

func main() {
	log.SetPrefix("greatings: ")
	log.SetFlags(0)

	// message, err := greatings.Hello("Gladys")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//
	// fmt.Println(message)
	names := []string{"Gladys", "Samatha", "Darrin"}
	messages, err := greatings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(messages)
}
