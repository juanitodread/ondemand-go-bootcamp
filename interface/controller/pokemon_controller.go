package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/juanitodread/ondemand-go-bootcamp/domain/model"
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
	var p []*model.Pokemon
	p, err := pc.pokemonInteractor.GetAll(p)

	if err != nil {
		ctx.AbortWithError(500, err)
	}

	ctx.JSON(200, p)
}
