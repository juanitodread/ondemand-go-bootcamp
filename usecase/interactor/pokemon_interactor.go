package interactor

import (
	"github.com/juanitodread/ondemand-go-bootcamp/domain/model"
	"github.com/juanitodread/ondemand-go-bootcamp/usecase/presenter"
	"github.com/juanitodread/ondemand-go-bootcamp/usecase/repository"
)

type pokemonInteractor struct {
	PokemonRepository repository.PokemonRepository
	PokemonPresenter  presenter.PokemonPresenter
}

type PokemonInteractor interface {
	GetAll(p []*model.Pokemon) ([]*model.Pokemon, error)
}

func NewPokemonInteractor(r repository.PokemonRepository, p presenter.PokemonPresenter) PokemonInteractor {
	return &pokemonInteractor{r, p}
}

func (pi *pokemonInteractor) GetAll(p []*model.Pokemon) ([]*model.Pokemon, error) {
	p, err := pi.PokemonRepository.FindAll(p)
	if err != nil {
		return nil, err
	}

	return pi.PokemonPresenter.ResponsePokemons(p), nil
}
