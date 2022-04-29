package datastore_test

import (
	"testing"

	"github.com/juanitodread/ondemand-go-bootcamp/infrastructure/datastore"
	tu "github.com/juanitodread/ondemand-go-bootcamp/testutil"
)

func TestNewPokemonDSEmptyPath(t *testing.T) {
	pokemonDS, err := datastore.NewPokemonDS("")

	tu.AssertNil(t, pokemonDS)
	tu.AssertEquals(t, "invalid CSV path: Path can't be empty", err.Error())
}

func TestNewPokemonDSInvalidCSVExtension(t *testing.T) {
	pokemonDS, err := datastore.NewPokemonDS("pokemons.txt")

	tu.AssertNil(t, pokemonDS)
	tu.AssertEquals(t, "invalid CSV path [pokemons.txt]: Path must have .csv extension", err.Error())
}

func TestPokemonDSReadFileNotExists(t *testing.T) {
	pokemonDS, _ := datastore.NewPokemonDS("not-exists-pokemons.csv")

	pokemonRecords, err := pokemonDS.Read()

	tu.AssertEquals(t, "error: CSV file not found at this path: [not-exists-pokemons.csv]", err.Error())
	tu.AssertNil(t, pokemonRecords)
}

func TestPokemonDSReadFileHasNoHeader(t *testing.T) {
	pokemonDS, _ := datastore.NewPokemonDS("../../assets/empty-pokemons.csv")

	pokemonRecords, err := pokemonDS.Read()

	tu.AssertEquals(t, "error: CSV file missing header: [../../assets/empty-pokemons.csv]", err.Error())
	tu.AssertNil(t, pokemonRecords)
}
