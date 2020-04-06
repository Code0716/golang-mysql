package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	//"github.com/jinzhu/gorm"

	"./constants"
	"./handler/connectDB"
)

// City struct
// gorm.Modelと記述するとcreated_at、updated_at、deleted_atが定義される。
type City struct {
	//gorm.Model
	ID          int    //`json:"id"`
	Name        string //`json:"name"`
	CountryCode string //`json:"country_code"`
	District    string //`json:"district"`
	Population  int    //`json:"population"`
}

func main() {

	r := gin.Default()
	r.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello world")
	})

	r.GET("/", func(c *gin.Context) {
		db := connectDB.ConnectMySQL(constants.DBTableWorld)
		defer db.Close()
		db.LogMode(true)
		// 勝手に複数形になるのを抑制
		db.SingularTable(true)

		city := []City{}
		//db.AutoMigrate(&city)
		db.Select("id,name,district,population").Find(&city)

		/*for _, value := range city {
			fmt.Println(value)
		}*/
		c.JSON(http.StatusOK, city)
		fmt.Println("DONE")
	})

	r.Run(":8080")
}
