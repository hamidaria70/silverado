package main

import (
	"fmt"
	"log"
	"practices/greetings"
)

func main() {
	names := []string{"Hamid", "Jack", "Joe"}
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	final, err := greetings.Hellos(names)
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println(final)
}