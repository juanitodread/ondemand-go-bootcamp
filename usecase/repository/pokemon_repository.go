package repository

import "github.com/juanitodread/ondemand-go-bootcamp/domain/model"

type PokemonRepository interface {
	FindAll(p []*model.Pokemon) ([]*model.Pokemon, error)
}
