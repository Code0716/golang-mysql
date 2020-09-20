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
		api.GET("/city", func(c *gin.Context) { cityController.GetAll(c) }) // 全レコード取得 cities

		preImagesController := controllers.NewPreImagesController(ConnectMySQL())
		api.GET("/image/pre_upload", func(c *gin.Context) { preImagesController.GetAll(c) }) // Preuploadされた一覧を取得

		imagesController := controllers.NewImagesController(ConnectMySQL())
		api.GET("/image/upload", func(c *gin.Context) { imagesController.GetAll(c) }) // uploadされた一覧を取得

	}

	//router.GET("/migrate", migrate.Migrate)

	router.Run(":8000")
}
