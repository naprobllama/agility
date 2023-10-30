package main

import (
	"projects/agility/swapi"
	"fmt"
)

const SWAPIBaseURL = `https://swapi.dev/`

// 1. Aggregate all character data using pagination calls. 
// 2. Rearrange slice alphabetically using sort and keys
// 3. Setup all other calls and get using wait groups and channels
// Use IDs as keys to order them up when they come back from the WG.
// 
// 4. 
// 
// starship
// home planet
// species

func main() {

	fmt.Print("\nStarting the app: \n")
	client := swapi.InitClient(SWAPIBaseURL)

	fmt.Print("\nInititlaized the client: \n")

	err := client.GetPeopleList(`l`)
	if err != nil {
		fmt.Printf("\nStarting the app: %s\n", err.Error())
		return
	}

	// make an intro package we can call

	// initial swapi client with base URL



	// fmt.Printf("\nclient: response body: %s\n", resBody)

	// make a print package we can call
	

}
