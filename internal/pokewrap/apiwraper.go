package pokewrap

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/scgitstuff/pokedex/internal/pokecache"
)

const (
	baseURL  = "https://pokeapi.co/api/v2"
	pageSize = 20
)

var cache = pokecache.NewCache(5 * time.Second)

func getBytes(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 299 {
		return []byte{}, fmt.Errorf("response failed with status code: %d", resp.StatusCode)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return data, nil
}

// mapPage: 0 based multiplier for offset
func GetLocationAreaPage(mapPage int) (LocationAreaPage, error) {
	// https://pokeapi.co/api/v2/location-area?limit=20&offset=0

	url := fmt.Sprintf("%s/location-area?limit=%v&offset=%v", baseURL, pageSize, pageSize*mapPage)

	var data []byte
	var hit bool
	var err error
	data, hit = cache.Get(url)
	if !hit {
		data, err = getBytes(url)
		if err != nil {
			return LocationAreaPage{}, err
		}

		cache.Add(url, data)
	}
	// fmt.Println(PrettyJSON(string(data)))

	area := LocationAreaPage{}
	if err := json.Unmarshal(data, &area); err != nil {
		return LocationAreaPage{}, err
	}

	return area, nil
}

// name or ID will work
func GetLocationAreaEncounter(loc string) (LocationAreaEncounter, error) {
	// https://pokeapi.co/api/v2/location-area/sunyshore-city-area

	url := fmt.Sprintf("%s/location-area/%s", baseURL, loc)

	var data []byte
	var hit bool
	var err error
	data, hit = cache.Get(url)
	if !hit {
		data, err = getBytes(url)
		if err != nil {
			return LocationAreaEncounter{}, err
		}

		cache.Add(url, data)
	}
	// fmt.Println(PrettyJSON(string(data)))

	encounter := LocationAreaEncounter{}
	if err := json.Unmarshal(data, &encounter); err != nil {
		return LocationAreaEncounter{}, err
	}

	return encounter, nil
}

// https://pokeapi.co/api/v2/pokemon/{id or name}/
func GetPokemon(name string) (Pokemon, error) {
	// https://pokeapi.co/api/v2/pokemon/pikachu/

	url := fmt.Sprintf("%s/pokemon/%s", baseURL, name)

	var data []byte
	var hit bool
	var err error
	data, hit = cache.Get(url)
	if !hit {
		data, err = getBytes(url)
		if err != nil {
			return Pokemon{}, err
		}

		cache.Add(url, data)
	}
	// fmt.Println(PrettyJSON(string(data)))

	vermin := Pokemon{}
	if err := json.Unmarshal(data, &vermin); err != nil {
		return Pokemon{}, err
	}

	return vermin, nil
}

func PrettyJSON(str string) (string, error) {
	var bytesBuffer bytes.Buffer
	if err := json.Indent(&bytesBuffer, []byte(str), "", "    "); err != nil {
		return "", err
	}

	return bytesBuffer.String(), nil
}
