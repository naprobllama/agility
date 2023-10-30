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

func (sc *SWAPIClient) GetPeopleList(searchChar string) error {
	searchChar = `l`
	// Get full list from all pages
	err := getPeopleList(searchChar, sc.BaseURL)
	return err
	
	// Alphabetize() the list

	// GetAll starships
	// Add all starship data

	// Get homeworld data
	// add into the list

	// Get species data
	// add into the list

	// return the list 
}