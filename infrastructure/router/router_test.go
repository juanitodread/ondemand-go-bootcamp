package router_test

import (
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/juanitodread/ondemand-go-bootcamp/infrastructure/router"
	"github.com/juanitodread/ondemand-go-bootcamp/interface/controller"
	tu "github.com/juanitodread/ondemand-go-bootcamp/testutil"
)

type PokemonControllerMock struct {
	controller.PokemonController
}

func (pcm PokemonControllerMock) GetPokemons(ctx *gin.Context) {}

func TestRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	appMock := controller.AppController{}
	appMock.Pokemon = PokemonControllerMock{}

	router := router.NewRouter(r, appMock)
	routes := router.Routes()

	tu.AssertEquals(t, 1, len(routes))
	tu.AssertEquals(t, "GET", routes[0].Method)
	tu.AssertEquals(t, "/pokemons", routes[0].Path)
}
