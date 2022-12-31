package router

import (
	"github.com/georgexx009-wizeline/ondemand-go-bootcamp/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(r *gin.Engine, c *controller.Controller) *gin.Engine {
	v1Group := r.Group("/api/v1")

	v1Group.GET("/read-external-api", c.ReadExternalApi)
	v1Group.GET("/read-pokemon-csv", c.ReadPokemonCSV)

	return r
}
