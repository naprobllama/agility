package swapi // make into character package? 

import(
	"net/http"
	"io"
	"encoding/json"
	// "fmt"
	"strconv"
	"sync"
	"sort"
)	

type PeopleResponse struct {
	Count int `json:"count"`
	Next *string `json:"next"`
	Previous *string `json:"previous"`
	Results []People `json:"results"`
}

type People struct {
	Name string `json:"name"`
	HairColor string `json:"hair_color"`
	StarshipURLs []string `json:"starships"`
	Starships []Starship `json:"starships_data"`
	Homeworld string `json:"homeworld"`
	Planet Planet
	SpeciesURLs []string `json:"species"`
	Species []Species // Should only be one species but just in case someone is a hybrid. Like Spock. :) 
}

type Starship struct {
	PersonName string
	Name string `json:"name"`
	Model string `json:"model"`
	StarshipClass string `json:"starship_class"`
	Manufacturer string `json:"manufacturer"`
	Crew string `json:"crew"`
	CargoCapacity string `json:"cargo_capacity"`
	URL string `json:"url"`
}

type Planet struct {
	PersonName string
	Name string `json:"name"`
	Climate string `json:"climate"`
	Population string `json:"population"`
	Diameter string `json:"diameter"`
	Gravity string `json:"gravity"`
}

type Species struct {
	PersonName string
	Name string `json:"name"`
	Language string `json:"language"`
	AverageLifespan string `json:"average_lifespan"`
}

// GetPeopleList() gets all the people from all api pages
func getPeopleList(searchChar string, baseURL string) ([]People, error) {

	people := make([]People, 0)

	var initialData PeopleResponse

	initialURL := baseURL+`api/people/?search=` + searchChar
	pageURL := baseURL+`api/people/?search=` + searchChar + `&page=`
	initialResp, err := get(initialURL)
	if err != nil {
		return people, err
	}
	
	if err := json.Unmarshal([]byte(initialResp), &initialData); err != nil {
		return people, err
	}

	if initialData.Next == nil {
		return initialData.Results, nil
	}

	for _, person := range initialData.Results {
		people = append(people, person)
	}

	pageTotal := ((initialData.Count - 1) / len(initialData.Results)) + 1
	pageChan := make(chan []People, pageTotal)
	var wg sync.WaitGroup

	for i := 1; i < pageTotal; i++ {
		url := pageURL + strconv.Itoa(i+1)

		wg.Add(1)

		go func() {
			defer wg.Done()
			var data PeopleResponse

			ok := false
			for !ok { // Because my home internet provider drops every 7th request or so

				response, _ := get(url)
				err := json.Unmarshal([]byte(response), &data)
				if err != nil {
					// keeping for debugging because my internet is sketchy
					// fmt.Printf("\nJSON PARSING FAILED: %s ", err.Error())
				} else {
					ok = true
				}
			}
			pageChan <- data.Results
		}()
	}
	
	wg.Wait()
	close(pageChan)

	for list := range pageChan {
		for _, person := range list {
			people = append(people, person)
		}
	}

	return people, nil
}

// TODO: remove from this package into a http package
func get(url string) ([]byte, error) {

    resp, err := http.Get(url)
    if err != nil {
        return make([]byte,0), err
    }
	defer resp.Body.Close()
	
	return io.ReadAll(resp.Body)
}

func alphabetizePeople(people []People) []People {

	sort.Slice(people, func(i, j int) bool {
		return people[i].Name < people[j].Name
	})

	return people
}

func fillInStarships(people []People) []People {
	var wg sync.WaitGroup
	starshipChan := make(chan Starship, 500) // todo make this dynamic

	for _, person := range people {
		for _, shipURL := range person.StarshipURLs {

			wg.Add(1)
			go func(peeps People, url string) {
				defer wg.Done()

				var ship Starship

				ok := false
				for !ok { // Because my home internet provider drops every 7th request or so
					
					response, _ := get(url)
					err := json.Unmarshal([]byte(response), &ship)
					if err != nil {
						// keeping for debugging because my internet is sketchy
						// fmt.Printf("\nJSON PARSING FAILED: %s ", err.Error())
					} else {
						ok = true
					}
				}

				ship.PersonName = peeps.Name

				starshipChan <- ship
			}(person, shipURL)
		}
	}

	wg.Wait()
	close(starshipChan)

	for ship := range starshipChan { 
		for i := 0; i < len(people); i++ {
			if ship.PersonName == people[i].Name {
				people[i].Starships = append(people[i].Starships,ship) // save data to slice
			}
		}
	}

	return people // Todo: improve error handling
}


func fillInSpecies(people []People) []People {
	var wg sync.WaitGroup
	speciesChan := make(chan Species, 500) // todo make this dynamic

	for _, person := range people {
		for _, speciesURL := range person.SpeciesURLs {

			wg.Add(1)
			go func(peeps People, url string) {
				defer wg.Done()

				var species Species

				ok := false
				for !ok { // Because my home internet provider drops every 7th request or so
					
					response, _ := get(url)
					err := json.Unmarshal([]byte(response), &species)
					if err != nil {
						// keeping for debugging because my internet is sketchy
						// fmt.Printf("\nJSON PARSING FAILED: %s ", err.Error())
					} else {
						ok = true
					}
				}

				species.PersonName = peeps.Name

				speciesChan <- species
			}(person, speciesURL)
		}
	}

	wg.Wait()
	close(speciesChan)

	for species := range speciesChan { 
		for i := 0; i < len(people); i++ {
			if species.PersonName == people[i].Name {
				people[i].Species = append(people[i].Species,species) // save data to slice
			}
		}
	}

	return people // Todo: improve error handling
}

func fillInPlanet(people []People) []People {
	var wg sync.WaitGroup
	homeworldChan := make(chan Planet, 150) // todo make this dynamic

	for _, person := range people {
		wg.Add(1)
		go func(peep People) {
			defer wg.Done()

			var planet Planet

			ok := false
			for !ok { // Because my home internet provider drops every 7th request or so
				response, _ := get(peep.Homeworld)
				err := json.Unmarshal([]byte(response), &planet)
				if err != nil {
					// keeping for debugging because my internet is sketchy
					// fmt.Printf("\nJSON PARSING FAILED: %s ", err.Error())
				} else {
					ok = true
				}
			}

			planet.PersonName = peep.Name // TODO: make the url the key

			homeworldChan <- planet
		}(person)
	}

	wg.Wait()
	close(homeworldChan)

	for planet := range homeworldChan { 
		for i := 0; i < len(people); i++ {
			if planet.PersonName == people[i].Name {
				people[i].Planet = planet
			}
		}
	}

	return people // Todo: improve error handling
}