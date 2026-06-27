package pokeapi

import (
	"testing"
)

func TestGetArea(t *testing.T) {
	const apiUrl string = "https://pokeapi.co/api/v2/location-area/"

	areas, err := GetArea(apiUrl)
	if err != nil {
		t.Fatalf("error: %v getting locations from api", err)
	}

	for _, location := range areas.Results {
		if len(location.LocationName) <= 0 {
			t.Errorf("length of location name was 0. Returned: %v\n", location)
		}
	}
}
