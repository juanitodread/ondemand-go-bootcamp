package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/juanitodread/ondemand-go-bootcamp/usecase/interactor"
)

type pokemonController struct {
	pokemonInteractor interactor.PokemonInteractor
}

type PokemonController interface {
	GetPokemons(ctx *gin.Context)
}

func NewPokemonController(pi interactor.PokemonInteractor) PokemonController {
	return &pokemonController{pi}
}

func (pc *pokemonController) GetPokemons(ctx *gin.Context) {
	p, err := pc.pokemonInteractor.GetAll()

	if err != nil {
		ctx.JSON(500, err.Error())
	} else {
		ctx.JSON(200, p)
	}
}
