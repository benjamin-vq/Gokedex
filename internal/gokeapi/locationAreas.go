package gokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

// https://mholt.github.io/json-to-go/2
type LocationArea struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

var (
	locationCache = NewCache(120 * time.Second)
)

func GetLocationsCache() *Cache {
	return locationCache
}

func GetLocations(url *string) (*LocationArea, error) {

	locationUrl := apiBaseURL + "/location-area?offset=0&limit=20"
	if url != nil {
		locationUrl = *url
	}

	cached, hit := getLocFromCache(&locationUrl)

	if hit {
		return cached, nil
	}

	res, err := http.Get(locationUrl)

	if err != nil {
		log.Printf("GET failed with error %v\n", err)
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 || err != nil {
		log.Printf("Response failed with status %d and body %s\n", res.StatusCode, body)
		return nil, errors.New("invalid api response")
	}

	locations := &LocationArea{}
	err = json.Unmarshal(body, locations)

	if err != nil {
		return nil, err
	}

	locationCache.Add(locationUrl, body)

	return locations, nil
}

func getLocFromCache(url *string) (locations *LocationArea, hit bool) {

	entry, hit := locationCache.Get(*url)

	if !hit {
		return nil, false
	}

	locations = &LocationArea{}
	err := json.Unmarshal(entry, locations)

	// Should never happen, if it was cached it means it was previously
	// unmarshalled
	if err != nil {
		return nil, false
	}

	return locations, true
}
