package presenter

import "github.com/juanitodread/ondemand-go-bootcamp/domain/model"

type PokemonPresenter interface {
	ResponsePokemons(p []*model.Pokemon) []*model.Pokemon
}
