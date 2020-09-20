package infrastructure

import (
	"../interfaces/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Router() {

	router := gin.Default()

	// CORS 対応
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:8080"}
	router.Use(cors.New(config))

	// api group
	api := router.Group("/api")
	{
		cityController := controllers.NewCityController(ConnectMySQL())

		// 全レコード取得 cities
		api.GET("/city", func(c *gin.Context) { cityController.GetAll(c) })

	}

	//router.GET("/migrate", migrate.Migrate)

	router.Run(":8000")
}
