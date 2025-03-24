package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationResponse struct {
	Results []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(cfg *config) error {
	const baseURL = "https://pokeapi.co/api/v2/location-area/"
	resp, err := http.Get(baseURL)
	if err != nil {
		return fmt.Errorf("failed to get location areas: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return nil
	}

	var locations LocationResponse
	err = json.Unmarshal(body, &locations)
	if err != nil {
		fmt.Println(err)
	}
	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
