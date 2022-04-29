package repository

import "github.com/juanitodread/ondemand-go-bootcamp/domain/model"

type PokemonRepository interface {
	FindAll() ([]*model.Pokemon, error)
}
