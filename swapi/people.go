package swapi // make into character package? 

import(
	"fmt"
	"net/http"
	"io"
	"encoding/json"
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
	Starships []string `json:"starships"`
	Homeworld string `json:"homeworld"`
	Species []string `json:"species"`
}

type Starship struct {
	Name string `json:"name"`
	Model string `json:"model"`
	StarshipClass string `json:"starship_class"`
	Manufacturer string `json:"manufacturer"`
	Crew string `json:"crew"`
	CargoCapacity string `json:"cargo_capacity"`
}

type Planet struct {
	Name string `json:"name"`
	Climate string `json:"climate"`
	Population string `json:"population"`
	Diameter string `json:"diameter"`
	Gravity string `json:"gravity"`
}

type Species struct {
	Name string `json:"name"`
	Language string `json:"language"`
	AverageLifespan string `json:"average_lifespan"`
}

// GetPeopleList() gets all the people from all api pages
func getPeopleList(searchChar string, baseURL string) ([]People, error) {
	people := make([]People, 0)

	url := baseURL+`api/people/?search=` + searchChar
	var next *string
	start := "start"
	next = &start
	
	for next != nil {

		var data PeopleResponse

		response, err := get(url)
		if err != nil {
			return err
		}

		if err := json.Unmarshal([]byte(response), &data); err != nil {
			return err
		}

		for _, v := range data.Results {
			people = append(people, v)
		}

		
		next = data.Next
		if next != nil {
			url = *data.Next
		}
	}

	return people, nil
}

//TODO: remove from this package into a http package
func get(url string) ([]byte, error) {

    resp, err := http.Get(url)
    if err != nil {
        return make([]byte,0), err
    }
	defer resp.Body.Close()
	
	return io.ReadAll(resp.Body)
}







/*


	
	// get page total and count


	// get additional pages
	pageTotal := 1 // (total_items - 1) / page_size + 1
	for page := 0; page < pageTotal; page++{
		

		pages = (total_items - 1) / page_size + 1
	}

	pages := 0
	for data.Next != nil { // <-- doesn't handle the lambda case yet
		go GetPage(pageChan)
		pages++
	}

	for i := 0; i < pages;{
        list, ok := <-pageChan
        if ok {
			// append list to the full list etc
			i++ // not failure resistant but will work
        } 
    }


	 := get()
	
	
	if err := json.Unmarshal([]byte(respBody), &data); err != nil {
		fmt.Printf("Failed to unmarshal", err)
		os.Exit(1)
	}


    resp, err := http.Get("https://gobyexample.com")
    if err != nil {
        panic(err)
	}
	
	defer resp.Body.Close()
	
	respBody, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("client: could not read response body: %s\n", err)
		os.Exit(1)
	}

    fmt.Println("Response status:", resp.Status)


}

func getPage() {

}

// GetCharcterList() gets all the characters from the API and returns them
func (c *Character) AlphabetizeCharacterList( /* wrap sort pamas here) {



}

*/
