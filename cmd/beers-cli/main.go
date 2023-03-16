package main

import (
	"flag"

	beerscli "github.com/dgarciabravo13/go-beers-api/internal"
	"github.com/dgarciabravo13/go-beers-api/internal/cli"
	"github.com/dgarciabravo13/go-beers-api/internal/storage/api"
	"github.com/dgarciabravo13/go-beers-api/internal/storage/csv"
	"github.com/spf13/cobra"
)

func main() {

	csvData := flag.Bool("csv", false, "load data from csv")
	flag.Parse()

	var repo beerscli.BeerRepo
	repo = csv.NewRepository()

	if !*csvData {
		repo = api.NewOntarioRepository()
	}

	rootCmd := &cobra.Command{Use: "beers-cli"}
	rootCmd.AddCommand(cli.InitBeersCmd(repo))
	rootCmd.Execute()
}
