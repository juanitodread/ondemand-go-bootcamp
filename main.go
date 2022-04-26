package main

import (
	"github.com/gin-gonic/gin"
	"github.com/juanitodread/ondemand-go-bootcamp/infraestructure/router"
	"github.com/juanitodread/ondemand-go-bootcamp/registry"
)

func main() {
	r := registry.NewRegistry()

	g := gin.Default()
	g = router.NewRouter(g, r.NewAppController())

	g.Run()
}
