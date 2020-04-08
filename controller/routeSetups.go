package routes

import (
	// "EventShoop/Models"
	_ "log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// // api for states name
// func apiStates(c *gin.Context) {
// 	states, err := models.SQLStates()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	c.JSON(http.StatusOK, states)
// }

// // api for service name
// func apiServices(c *gin.Context) {
// 	services, err := models.SQLServices()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	c.JSON(http.StatusOK, services)
// }

// // api for location
// func apiLocation(c *gin.Context) {
// 	state := c.Query("state")
// 	locations, err := models.SQLLocations(state)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	c.JSON(http.StatusOK, locations)
// }

// serve home page
func home(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
