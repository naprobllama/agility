package main

import (
	"projects/agility/swapi"
	"fmt"
	"bufio"
	"os"
)

const SWAPIBaseURL = `https://swapi.dev/`

func main() {

	client := swapi.InitClient(SWAPIBaseURL)

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter Star Wars Character Search: ")
	scanner.Scan()
	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}

	people := client.GetPeopleList(scanner.Text())

	// print the data
	for _, person := range people {
		fmt.Printf("\nName: %s",person.Name)
		fmt.Printf("\nHair Color: %s",person.HairColor)
		if len(person.Starships) > 0 {
			fmt.Print("\nStarships:")
		}
		for _, ship := range person.Starships {
			fmt.Printf("\n     Ship Name: %s", ship.Name)
			fmt.Printf("\n     Ship Model: %s", ship.Model)
			fmt.Printf("\n     Ship Class: %s", ship.StarshipClass)
			fmt.Printf("\n     Ship Cargo Capacity: %s\n", ship.CargoCapacity)
		}

		fmt.Print("\nHomeworld:")
		fmt.Printf("\n     Name: %s", person.Planet.Name)
		fmt.Printf("\n     Climate: %s", person.Planet.Climate)
		fmt.Printf("\n     Population: %s", person.Planet.Population)
		fmt.Printf("\n     Diameter: %s", person.Planet.Diameter)
		fmt.Printf("\n     Gravity: %s\n", person.Planet.Gravity)

		if len(person.Species) > 0 {
			fmt.Print("\nSpecies:")
		}

		for _, species := range person.Species {
			fmt.Printf("\n     Name: %s", species.Name)
			fmt.Printf("\n     Language: %s", species.Language)
			fmt.Printf("\n     Average Life Span: %s\n", species.AverageLifespan)
		}

		fmt.Print("**********************************\n")
	}

	fmt.Print("\n\n\nMAY THE FORCE BE WITH YOU...\n\n\n")

}
