package repository

import (
	"fmt"

	"github.com/juanitodread/ondemand-go-bootcamp/domain/model"
)

type pokemonRepository struct{}

type PokemonRepository interface {
	FindAll(p []*model.Pokemon) ([]*model.Pokemon, error)
}

func NewPokemonRepository() PokemonRepository {
	return &pokemonRepository{}
}

func (pr *pokemonRepository) FindAll(p []*model.Pokemon) ([]*model.Pokemon, error) {
	// returning in memory data
	pokemons := make([]*model.Pokemon, 5)
	for i := 0; i < 5; i++ {
		ID := i + 1
		pokemon := model.Pokemon{
			ID:   ID,
			Name: fmt.Sprintf("pokemons-%d", ID),
		}
		pokemons[i] = &pokemon
	}

	return pokemons, nil
}
