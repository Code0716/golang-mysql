package layout

import (
	"fmt"
	"net/http"

	"../pkg/modules"
	"../pkg/scraping"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Layout func
func Layout() {

	router := gin.Default()

	// CORS 対応
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	router.Use(cors.New(config))

	// api group
	api := router.Group("/api")
	{
		// 全レコード取得 cities
		api.GET("/city", getCities)
		//　１件取得 city
		api.GET("/city/:name", getCity)
		// 全レコード取得 countries
		api.GET("/country", getCountries)
		//　１件取得 country
		api.GET("/country/:name", getCountry)
		// 大陸一覧
		api.GET("/continent", getContinentsList)
		// スクレイピング
		api.GET("/scraping", getScrape)
	}

	router.Run(":8000")
}

// 街一覧
func getCities(ginContext *gin.Context) {
	// interface
	var world modules.WorldDB
	// struct
	var cities modules.Cities

	world = &cities
	world.GetAll()
	ginContext.JSON(http.StatusOK, cities)
}

//　まち
func getCity(ginContext *gin.Context) {
	// interface
	var world modules.WorldDB
	// struct
	var cities modules.Cities

	name := ginContext.Param("name")

	world = &cities
	world.GetSingle(name)
	ginContext.JSON(http.StatusOK, cities)
}

// 国一覧
func getCountries(ginContext *gin.Context) {
	// interface
	var world modules.WorldDB
	// struct
	var countries modules.Countries

	world = &countries

	world.GetAll()
	ginContext.JSON(http.StatusOK, countries)
}

// 国get
func getCountry(ginContext *gin.Context) {
	// interface
	var world modules.WorldDB
	// struct
	var countries modules.Countries

	name := ginContext.Param("name")

	world = &countries
	world.GetSingle(name)
	ginContext.JSON(http.StatusOK, countries)
}

//スクレイピング　TODO
func getScrape(ginContext *gin.Context) {
	data := scraping.Scrape()
	fmt.Println(data)
	//ginContext.JSON(http.StatusOK, data)

}

// 大陸一覧.
func getContinentsList(ginContext *gin.Context) {
	data := modules.GetContinentsDB("GetCountDB")
	ginContext.JSON(http.StatusOK, data)
}
