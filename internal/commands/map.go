package commands

import (
	"fmt"

	"github.com/scgitstuff/pokedex/internal/pokewrap"
)

var mapPage = -1

func commandMap(cmdArgs []string) error {
	area, err := pokewrap.GetLocationAreaPage(mapPage + 1)
	if err != nil {
		return err
	}

	mapPage++

	for i := range area.Results {
		fmt.Println(area.Results[i].Name)
	}

	return nil
}

func commandMapb(cmdArgs []string) error {
	if mapPage < 1 {
		return nil
	}

	area, err := pokewrap.GetLocationAreaPage(mapPage - 1)
	if err != nil {
		return err
	}

	mapPage--

	for i := range area.Results {
		fmt.Println(area.Results[i].Name)
	}

	return nil
}
