package registry

import (
	"github.com/juanitodread/ondemand-go-bootcamp/infrastructure/datastore"
	"github.com/juanitodread/ondemand-go-bootcamp/interface/controller"
)

type registry struct {
	pokemonDS datastore.PokemonDS
}

type Registry interface {
	NewAppController() controller.AppController
}

func NewRegistry(pokemonDS datastore.PokemonDS) Registry {
	return &registry{
		pokemonDS: pokemonDS,
	}
}

func (r *registry) NewAppController() controller.AppController {
	return controller.AppController{
		Pokemon: r.NewPokemonController(),
	}
}
