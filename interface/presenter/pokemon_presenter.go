package presenter

import (
	"github.com/juanitodread/ondemand-go-bootcamp/domain/model"
	"github.com/juanitodread/ondemand-go-bootcamp/usecase/presenter"
)

type pokemonPresenter struct{}

func NewPokemonPresenter() presenter.PokemonPresenter {
	return &pokemonPresenter{}
}

func (pp *pokemonPresenter) ResponsePokemons(p []*model.Pokemon) []*model.Pokemon {
	return p
}
