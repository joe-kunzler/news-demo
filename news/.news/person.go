package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"
)

func main() {
	// Create a new Person object
	person := &Person{
		Name: "Alice",
		Age:  30,
	}

	// Serialize the Person object to a byte array
	data, err := proto.Marshal(person)
	if err != nil {
		log.Fatalln("Failed to serialize person:", err)
	}

	// Deserialize the byte array back into a Person object
	newPerson := &Person{}
	if err := proto.Unmarshal(data, newPerson); err != nil {
		log.Fatalln("Failed to deserialize person:", err)
	}

	fmt.Println(newPerson.String())
}
