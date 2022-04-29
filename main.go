package main

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/juanitodread/ondemand-go-bootcamp/config"
	"github.com/juanitodread/ondemand-go-bootcamp/infrastructure/datastore"
	"github.com/juanitodread/ondemand-go-bootcamp/infrastructure/router"
	"github.com/juanitodread/ondemand-go-bootcamp/registry"
)

func main() {
	config := config.Load()

	pokemonDS, err := datastore.NewPokemonDS(config.PokemonDS.Path)

	// If there's an error getting the DS we can't serve any response
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	r := registry.NewRegistry(pokemonDS)

	g := gin.Default()
	g = router.NewRouter(g, r.NewAppController())

	g.Run(fmt.Sprintf(":%d", config.Server.Port))
}
