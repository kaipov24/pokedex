package main

import (
	"fmt"
)

func commandExplore(cfg *config, args ...string) error {
	fmt.Println("Exploring")
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous

	// for _, loc := range locationsResp.Results {
	// 	fmt.Println(loc.Name)
	// }

	return nil
}
