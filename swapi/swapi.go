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

	withEverything := fillInSpecies(withPlanets)

	return alphabetizePeople(withEverything)
}