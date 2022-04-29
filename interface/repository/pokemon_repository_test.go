package repository_test

import (
	"errors"
	"testing"

	"github.com/juanitodread/ondemand-go-bootcamp/domain/model"
	"github.com/juanitodread/ondemand-go-bootcamp/interface/repository"
	tu "github.com/juanitodread/ondemand-go-bootcamp/testutil"
)

type PokemonDSMock struct{}

func (pdsMock PokemonDSMock) Read() ([][]string, error) {
	return [][]string{
		{"1", "picachu", "electric", "prop4", "prop5"},
		{"2", "bulbasaur", "grass", "prop4", "prop5"},
	}, nil
}

type PokemonDSMockError struct{}

func (pdsMock PokemonDSMockError) Read() ([][]string, error) {
	return nil, errors.New("ds-error")
}

type PokemonDSMockEmptyRecords struct{}

func (pdsMock PokemonDSMockEmptyRecords) Read() ([][]string, error) {
	return [][]string{}, nil
}

type PokemonDSMockInvalidRecordColumn struct{}

func (pdsMock PokemonDSMockInvalidRecordColumn) Read() ([][]string, error) {
	return [][]string{
		{"1", "picachu", "electric", "prop4", "prop5"},
		{"invalid-record"},
	}, nil
}

type PokemonDSMockInvalidRecordValue struct{}

func (pdsMock PokemonDSMockInvalidRecordValue) Read() ([][]string, error) {
	return [][]string{
		{"1", "picachu", "electric", "prop4", "prop5"},
		{"invalid-id", "skipped", "invalid", "prop4", "prop5"},
	}, nil
}

func TestNewPokemonRepository(t *testing.T) {
	pr := repository.NewPokemonRepository(nil)

	tu.AssertNotNil(t, pr)
}

func TestPokemonRepositoryFindAll(t *testing.T) {
	pDSMock := PokemonDSMock{}

	pr := repository.NewPokemonRepository(pDSMock)

	pokemons, err := pr.FindAll()

	tu.AssertEquals(t, 2, len(pokemons))
	tu.AssertEquals(t, model.Pokemon{ID: 1, Name: "picachu", Type: "electric"}, *pokemons[0])
	tu.AssertEquals(t, model.Pokemon{ID: 2, Name: "bulbasaur", Type: "grass"}, *pokemons[1])
	tu.AssertNil(t, err)
}

func TestPokemonRepositoryFindAllDSError(t *testing.T) {
	pDSMock := PokemonDSMockError{}

	pr := repository.NewPokemonRepository(pDSMock)

	pokemons, err := pr.FindAll()

	tu.AssertEquals(t, "ds-error", err.Error())
	tu.AssertNil(t, pokemons)
}

func TestPokemonRepositoryFindAllDSEmptyRecords(t *testing.T) {
	pDSMock := PokemonDSMockEmptyRecords{}

	pr := repository.NewPokemonRepository(pDSMock)

	pokemons, err := pr.FindAll()

	tu.AssertEquals(t, 0, len(pokemons))
	tu.AssertNil(t, err)
}

func TestPokemonRepositoryFindAllDSInvalidRecordColumn(t *testing.T) {
	pDSMock := PokemonDSMockInvalidRecordColumn{}

	pr := repository.NewPokemonRepository(pDSMock)

	pokemons, err := pr.FindAll()

	tu.AssertEquals(t, 1, len(pokemons))
	tu.AssertEquals(t, model.Pokemon{ID: 1, Name: "picachu", Type: "electric"}, *pokemons[0])
	tu.AssertNil(t, err)
}

func TestPokemonRepositoryFindAllDSInvalidRecordValue(t *testing.T) {
	pDSMock := PokemonDSMockInvalidRecordValue{}

	pr := repository.NewPokemonRepository(pDSMock)

	pokemons, err := pr.FindAll()

	tu.AssertEquals(t, 1, len(pokemons))
	tu.AssertEquals(t, model.Pokemon{ID: 1, Name: "picachu", Type: "electric"}, *pokemons[0])
	tu.AssertNil(t, err)
}
