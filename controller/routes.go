package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/static"
)

// Allroutes : used for creating and managing
// all the routes
func Allroutes() *gin.Engine {

	r := gin.Default()

	r.Use(static.Serve("/", static.LocalFile("frontend_logic/dist/", true)))

	r.NoRoute(func(c *gin.Context) {
		c.File("frontend_logic/dist/index.html")
	})

	// api := r.Group("/json/api")
	// {
	// 	api.GET("/states", apiStates)
	// 	api.GET("/services", apiServices)
	// 	api.GET("/locations", apiLocation)
	// }
	
	return r
}
