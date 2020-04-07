package main

import (
	//"fmt"
	"net/http"

	"./pkg/modules"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	// 全レコード取得
	r.GET("/cities", func(c *gin.Context) {
		// struct
		var cities modules.Cities
		// interface
		var world modules.WorldDB
		world = &cities
		world.GetAll()
		c.JSON(http.StatusOK, cities)

	})

	// 全レコード取得
	r.GET("/countries", func(c *gin.Context) {
		// struct
		var countries modules.Countries
		// interface
		var world modules.WorldDB
		world = &countries
		world.GetAll()
		c.JSON(http.StatusOK, countries)

	})

	/*
		TODO
		//一件取得
			r.GET("/country/:name", func(c *gin.Context) {
				// struct
				var country modules.Country
				// interface
				var world modules.WorldDB
				name := c.Param("name")
				world(&country, name)
				world.GetSingle()
				c.JSON(http.StatusOK, country)

			})*/

	r.Run(":8080")
}

//curl http://localhost:8080/
