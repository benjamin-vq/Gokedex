package gokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"
)

var (
	exploreCache = NewCache(120 * time.Second)
)

func GetExploreAreasCache() *Cache {
	return exploreCache
}

func GetExploreAreas(area *string) (*ExploreArea, error) {
	exploreUrl := apiBaseURL + "/location-area/" + *area

	if cached, hit := getAreaFromCache(&exploreUrl); hit {
		fmt.Println("Retrieved area from cache :)")
		return cached, nil
	}

	res, err := http.Get(exploreUrl)

	if err != nil {
		return nil, err
	}

	body, err := io.ReadAll(res.Body)
	defer res.Body.Close()

	// TODO: Appropiate error messages for incorrect areas and such
	if res.StatusCode > 299 || err != nil {
		return nil, errors.New("invalid api response")
	}

	exploreArea := &ExploreArea{}
	err = json.Unmarshal(body, exploreArea)

	if err != nil {
		return nil, errors.New("error unmarshalling JSON")
	}

	exploreCache.Add(exploreUrl, body)

	return exploreArea, nil
}

func getAreaFromCache(url *string) (areas *ExploreArea, hit bool) {

	entry, hit := exploreCache.Get(*url)

	if !hit {
		return nil, false
	}

	areas = &ExploreArea{}
	err := json.Unmarshal(entry, areas)

	// Should never happen, if it was cached it means it was previously
	// unmarshalled
	if err != nil {
		return nil, false
	}

	return areas, true
}
