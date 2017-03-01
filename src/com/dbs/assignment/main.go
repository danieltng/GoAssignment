package main

import (
	//"com/dbs/model"
	"com/dbs/service"
	"fmt"
	"log"
)

func handleInput() (query, location string) {
	fmt.Print("Enter search query: ")
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err)
		return
	}
	query = input
	fmt.Println("You have entered: " + input)

	fmt.Print("Enter location: ")
	_, err = fmt.Scanln(&input)
	if err != nil {
		log.Fatal(err)
		return
	}
	location = input
	fmt.Println(location)
	
	fmt.Println("Your search query is: " + query + "with location: " + location)
	
	return
}

func main() {
	query,location := handleInput()

	service.SearchByCriteria(query, location)
	//service.SearchByCriteria("Starbucks", "Singapore")
	fmt.Println("Program exits")
}
