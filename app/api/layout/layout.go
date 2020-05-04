package layout

import (
	"../pkg/modules"
	"../pkg/modules/images"
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
		api.GET("/city", modules.GetCities)
		//　１件取得 city
		api.GET("/city/:name", modules.GetCity)
		// 全レコード取得 countries
		api.GET("/country", modules.GetCountries)
		//　１件取得 country
		api.GET("/country/:name", modules.GetCountry)
		// 大陸一覧
		api.GET("/continent", modules.GetContinentsList)
		// スクレイピング
		api.GET("/scraping", scraping.GetScrape)

		preImg := images.PreImageController{}

		// Preuploadされた一覧を取得
		api.GET("/image/pre_upload", preImg.GetAll)
		// Preuploadする
		api.POST("/image/pre_upload", preImg.Upload)
		// 画像一枚だけ取得して、Base64で返す。
		api.GET("/image/pre_upload/:id", preImg.GetFile)
		//　Preuploadのでデータを一枚削除
		api.DELETE("/image/pre_upload/delete", preImg.Delete)
		//　本アップロード
		api.GET("/image/upload", preImg.ComitUpload)
	}

	router.Run(":8000")
}
