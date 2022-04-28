package registry

import (
	"github.com/juanitodread/ondemand-go-bootcamp/interface/controller"
	"github.com/juanitodread/ondemand-go-bootcamp/interface/presenter"
	"github.com/juanitodread/ondemand-go-bootcamp/interface/repository"
	ii "github.com/juanitodread/ondemand-go-bootcamp/usecase/interactor"
	ip "github.com/juanitodread/ondemand-go-bootcamp/usecase/presenter"
	ir "github.com/juanitodread/ondemand-go-bootcamp/usecase/repository"
)

func (r *registry) NewPokemonController() controller.PokemonController {
	return controller.NewPokemonController(r.NewPokemonInteractor())
}

func (r *registry) NewPokemonInteractor() ii.PokemonInteractor {
	return ii.NewPokemonInteractor(r.NewPokemonRepository(), r.NewPokemonPresenter())
}

func (r *registry) NewPokemonRepository() ir.PokemonRepository {
	return repository.NewPokemonRepository(r.pokemonDS)
}

func (r *registry) NewPokemonPresenter() ip.PokemonPresenter {
	return presenter.NewPokemonPresenter()
}
