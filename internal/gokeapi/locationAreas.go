package gokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"
)

type Config struct {
	Previous, Next *string
}

func NewConfig() *Config {
	next := apiBaseURL + "/location-area/?offset=0&limit=20"

	return &Config{
		nil,
		&next,
	}
}

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
	config        = NewConfig()
)

func GetLocationsCache() *Cache {
	return locationCache
}

func GetLocations() (*LocationArea, error) {

	if config.Next == nil {
		log.Println("There are no more locations")
		return nil, errors.New("no more locations")
	}

	cached, hit := getFromCache(config.Next)

	if hit {
		return cached, nil
	}

	res, err := http.Get(*config.Next)

	if err != nil {
		log.Printf("GET failed with error %v", err)
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	if res.StatusCode > 299 || err != nil {
		log.Printf("Response failed with status %d and body %s", res.StatusCode, body)
		return nil, errors.New("invalid api response")
	}

	locations := &LocationArea{}
	err = locations.UnmarshalResponse(body)

	if err != nil {
		return nil, err
	}

	locationCache.Add(*config.Next, body)
	config.updateUrls(locations.Previous, locations.Next)

	return locations, nil
}

func GetPreviousLocations() (*LocationArea, error) {

	if config.Previous == nil {
		log.Println("Previous URL is nil, there are no previous locations")
		return nil, errors.New("there are no previous locations")
	}

	entry, hit := getFromCache(config.Previous)

	if hit {
		log.Println("Entry retrieved from cache :)")
		config.updateUrls(entry.Previous, entry.Next)
		return entry, nil
	}

	return nil, errors.New("could not find entry in cache")
}

func (loc *LocationArea) UnmarshalResponse(httpBytes []byte) error {

	err := json.Unmarshal(httpBytes, &loc)

	if err != nil {
		log.Printf("Unable to marshal JSON response due to error %v", err)
		return err
	}

	return nil
}

func (c *Config) updateUrls(previous, next *string) {

	if previous == nil && next == nil {
		log.Fatal("Both previous and next are nil, shouldnt happen")
	}

	c.Previous = previous
	c.Next = next
}

func getFromCache(url *string) (locations *LocationArea, hit bool) {

	entry, hit := locationCache.Get(*url)

	if !hit {
		return nil, false
	}

	locations = &LocationArea{}
	err := locations.UnmarshalResponse(entry)

	if err != nil {
		return nil, false
	}

	return locations, true
}
