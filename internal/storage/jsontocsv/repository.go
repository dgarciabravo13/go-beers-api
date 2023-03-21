package jsontocsv

import (
	"encoding/csv"
	"fmt"
	"os"

	beerscli "github.com/dgarciabravo13/go-beers-api/internal"
)

func JsonToCsv(beers []beerscli.Beer) {
	file, err := os.Create("beers.csv")

	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)
	writer.Write([]string{"ID", "Name", "Tagline", "FirstBrewed", "Description"})

	for _, beer := range beers {
		err := writer.Write([]string{
			fmt.Sprintf("%d", beer.ID),
			beer.Name,
			beer.Tagline,
			beer.FirstBrewed,
			beer.Description,
		})
		if err != nil {
			panic(err)
		}
	}

	writer.Flush()

}
