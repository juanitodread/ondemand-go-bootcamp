package datastore

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/juanitodread/ondemand-go-bootcamp/domain/model"
)

type pokemonDS struct {
	csvPath string
}

type PokemonDS interface {
	Read() ([]*model.Pokemon, error)
}

func NewPokemonDS(csvPath string) PokemonDS {
	return &pokemonDS{
		csvPath: csvPath,
	}
}

func (pds *pokemonDS) Read() ([]*model.Pokemon, error) {
	pokemons, err := pds.parse()

	if err != nil {
		return nil, err
	}

	return pokemons, nil
}

func (pds *pokemonDS) parse() ([]*model.Pokemon, error) {
	fmt.Println(pds.csvPath)
	f, err := os.Open(pds.csvPath)

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	defer f.Close()

	csvReader := csv.NewReader(f)

	// skip header csv
	if _, err := csvReader.Read(); err != nil {
		panic(err)
	}

	records, err := csvReader.ReadAll()

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	pokemons := make([]*model.Pokemon, len(records))
	for i, record := range records {
		pokemon := model.Pokemon{}
		pokemon.ID, _ = strconv.Atoi(record[0])
		pokemon.Name = record[1]
		pokemon.Type = record[2]

		pokemons[i] = &pokemon
	}

	return pokemons, nil
}
