package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/juanitodread/ondemand-go-bootcamp/config"
	"github.com/juanitodread/ondemand-go-bootcamp/infrastructure/datastore"
	"github.com/juanitodread/ondemand-go-bootcamp/infrastructure/router"
	"github.com/juanitodread/ondemand-go-bootcamp/registry"
)

func main() {
	config := config.Load()

	pokemonDS := datastore.NewPokemonDS(config.PokemonDS.Path)

	r := registry.NewRegistry(pokemonDS)

	g := gin.Default()
	g = router.NewRouter(g, r.NewAppController())

	g.Run(fmt.Sprintf(":%d", config.Server.Port))
}
