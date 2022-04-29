package repository

import (
	"fmt"
	"strconv"

	"github.com/juanitodread/ondemand-go-bootcamp/domain/model"
	"github.com/juanitodread/ondemand-go-bootcamp/infrastructure/datastore"
)

type pokemonRepository struct {
	pokemonDS datastore.PokemonDS
}

type PokemonRepository interface {
	FindAll() ([]*model.Pokemon, error)
}

func NewPokemonRepository(pokemonDS datastore.PokemonDS) PokemonRepository {
	return &pokemonRepository{
		pokemonDS: pokemonDS,
	}
}

func (pr *pokemonRepository) FindAll() ([]*model.Pokemon, error) {
	pokemonRecords, err := pr.pokemonDS.Read()

	if err != nil {
		return nil, err
	}

	var pokemons []*model.Pokemon

	for i, record := range pokemonRecords {
		pokemon, err := pr.toPokemon(record)

		if err != nil {
			fmt.Printf("[warning]: datasource record: line=[%d] content=[%s] can't be converted to Pokemon. Record will be ignored\n", i, record)
		} else {
			pokemons = append(pokemons, pokemon)
		}
	}

	return pokemons, nil
}

func (pr *pokemonRepository) toPokemon(record []string) (*model.Pokemon, error) {
	if len(record) < 3 {
		err := fmt.Errorf("invalid record fields: record=[%s]", record)
		fmt.Println(err)
		return nil, err
	}

	ID, err := strconv.Atoi(record[0])
	name := record[1]
	ptype := record[2]

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &model.Pokemon{
		ID:   ID,
		Name: name,
		Type: ptype,
	}, nil
}
