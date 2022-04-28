package repository

import (
	"github.com/juanitodread/ondemand-go-bootcamp/domain/model"
	"github.com/juanitodread/ondemand-go-bootcamp/infrastructure/datastore"
)

type pokemonRepository struct {
	pokemonDS datastore.PokemonDS
}

type PokemonRepository interface {
	FindAll(p []*model.Pokemon) ([]*model.Pokemon, error)
}

func NewPokemonRepository(pokemonDS datastore.PokemonDS) PokemonRepository {
	return &pokemonRepository{
		pokemonDS: pokemonDS,
	}
}

func (pr *pokemonRepository) FindAll(p []*model.Pokemon) ([]*model.Pokemon, error) {
	pokemons, _ := pr.pokemonDS.Read()

	return pokemons, nil
}
