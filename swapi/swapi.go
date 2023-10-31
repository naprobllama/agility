package swapi

import(
	"fmt"
)	

type SWAPIClient struct {
	BaseURL string
}

func InitClient(baseURL string) (*SWAPIClient) {
	return &SWAPIClient {
		BaseURL: baseURL,
	}
}

func (sc *SWAPIClient) getBaseURL(url string) string {
	return sc.BaseURL
}

func (sc *SWAPIClient) GetPeopleList(searchChar string) error {
	searchChar = `l`
	// Get full list from all pages
	peopleList, err := getPeopleList(searchChar, sc.BaseURL)

	for _, v := range peopleList{
		fmt.Print("\nTHESE CAME BACK FROM THE CHANNEL: %s\n ", v.Name)
	}

	// Alphabetize the list
	// orderedPeeps := alphabetizePeople(peopleList)

	// GetAll starships
	withStarships := fillInStarships(peopleList)
	// for debug printing only
	for _, person := range withStarships {
		fmt.Printf("\nPerson >>>> : %s\n", person.Name)
		for _, ship := range person.Starships {
			fmt.Printf("Starship >>>> : %s\n", ship)
		}
	}

	withPlanets := fillInPlanet(withStarships)

	// Get homeworld data
	for _, person := range withPlanets {
		fmt.Printf("\nPerson >>>> : %s\n", person.Name)
		fmt.Printf("\nPlanet >>>> : %s\n", person.Planet)
	}

	withEverything := fillInSpecies(withPlanets)
	for _, person := range withEverything {
		fmt.Printf("\nPerson >>>> : %s\n", person.Name)
		fmt.Printf("\nSpecies >>>> : %s\n", person.Species)
	}

	alphabetized := alphabetizePeople(withEverything)
	for _, person := range alphabetized {
		fmt.Printf("\nPerson >>>> : %s\n", person.Name)
		fmt.Printf("\nSpecies >>>> : %s\n", person.Species)
	}

	// Get species data
	// add into the list

	// return the list 
	return err
}