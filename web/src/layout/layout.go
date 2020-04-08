package layout

import (
	"net/http"

	"../pkg/modules"
	"github.com/gin-gonic/gin"
)

// Routing func
func Layout() {
	// struct
	router := gin.Default()
	// interface
	var world modules.WorldDB

	// struct
	var cities modules.Cities
	var countries modules.Countries

	// 全レコード取得 cities
	router.GET("/cities", func(c *gin.Context) {
		world = &cities
		world.GetAll()
		c.JSON(http.StatusOK, cities)

	})

	//　１件取得 city
	router.GET("/city/:name", func(c *gin.Context) {
		name := c.Param("name")
		world = &cities
		world.GetSingle(name)
		c.JSON(http.StatusOK, cities)

	})

	// 全レコード取得 countries
	router.GET("/countries", func(c *gin.Context) {
		world = &countries
		world.GetAll()
		c.JSON(http.StatusOK, countries)
	})

	//　１件取得 country
	router.GET("/country/:name", func(c *gin.Context) {
		name := c.Param("name")
		world = &countries
		world.GetSingle(name)
		c.JSON(http.StatusOK, countries)
	})

	router.Run(":8080")
}

//curl http://localhost:8080/
