package config

type Config struct {
	PokemonDS struct {
		Path string
	}
	Server struct {
		Port int
	}
}

func Load() Config {
	config := Config{}
	config.PokemonDS.Path = "assets/pokemons.csv"
	config.Server.Port = 8080

	return config
}
