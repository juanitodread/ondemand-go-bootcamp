package datastore

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strings"
)

type pokemonDS struct {
	csvPath string
}

type PokemonDS interface {
	Read() ([][]string, error)
}

func NewPokemonDS(csvPath string) (PokemonDS, error) {
	if csvPath == "" {
		return nil, errors.New("invalid CSV path: Path can't be empty")
	}

	if !strings.HasSuffix(csvPath, ".csv") {
		return nil, fmt.Errorf("invalid CSV path [%s]: Path must have .csv extension", csvPath)
	}

	return &pokemonDS{
		csvPath: csvPath,
	}, nil
}

func (pds *pokemonDS) Read() ([][]string, error) {
	f, err := os.Open(pds.csvPath)

	if err != nil {
		return nil, fmt.Errorf("error: CSV file not found at this path: [%s]", pds.csvPath)
	}

	defer f.Close()

	csvReader := csv.NewReader(f)

	// skip header csv
	if _, err := csvReader.Read(); err != nil {
		return nil, fmt.Errorf("error: CSV file missing header: [%s]", pds.csvPath)
	}

	records, err := csvReader.ReadAll()

	// EOF is not an error in ReadAll() so in this case we're handling
	// unexpected error (OS, different columns length, etc)
	if err != nil {
		return nil, fmt.Errorf("error: unexpected error=[%s], path=[%s]", err.Error(), pds.csvPath)
	}

	return records, nil
}
