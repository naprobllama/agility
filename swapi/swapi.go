package swapi

import(

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

func (sc *SWAPIClient) GetPeopleList(searchChar string) []People {
	// Get full list from all pages
	peopleList, _ := getPeopleList(searchChar, sc.BaseURL)

	// Get All Starships
	withStarships := fillInStarships(peopleList)

	// Get All Planets
	withPlanets := fillInPlanet(withStarships)

	// Get Species
	withEverything := fillInSpecies(withPlanets)

	// Alphabetize and return
	return alphabetizePeople(withEverything)
}