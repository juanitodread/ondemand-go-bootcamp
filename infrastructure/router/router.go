package router

import (
	"github.com/gin-gonic/gin"
	"github.com/juanitodread/ondemand-go-bootcamp/interface/controller"
)

func NewRouter(r *gin.Engine, c controller.AppController) *gin.Engine {
	r.GET("/pokemons", c.Pokemon.GetPokemons)

	return r
}
