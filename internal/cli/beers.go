package cli

import (
	"fmt"
	"strconv"

	beerscli "github.com/dgarciabravo13/go-beers-api/internal"
	"github.com/dgarciabravo13/go-beers-api/internal/storage/jsontocsv"
	"github.com/spf13/cobra"
)

// CobraFn function definion of run cobra command
type CobraFn func(cmd *cobra.Command, args []string)

const idFlag = "id"

// InitBeersCmd initialize beers command
func InitBeersCmd(repository beerscli.BeerRepo) *cobra.Command {
	beersCmd := &cobra.Command{
		Use:   "beers",
		Short: "Print data about beers",
		Run:   runBeersFn(repository),
	}

	beersCmd.Flags().StringP(idFlag, "i", "", "id of the beer")

	return beersCmd
}

func runBeersFn(repository beerscli.BeerRepo) CobraFn {
	return func(cmd *cobra.Command, args []string) {
		beers, _ := repository.GetBeers()

		jsontocsv.JsonToCsv(beers)

		id, _ := cmd.Flags().GetString(idFlag)

		if id != "" {
			i, _ := strconv.Atoi(id)
			for _, beer := range beers {
				if beer.ID == i {
					fmt.Println(beer)
					return
				}
			}
		} else {
			fmt.Println(beers)
		}
	}
}
