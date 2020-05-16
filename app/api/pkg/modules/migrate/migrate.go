package migrate

import (
	"net/http"

	"../../../constants"
	"../../db"
	"../images"
	"github.com/gin-gonic/gin"
)

// Migrate db
func Migrate(ginContext *gin.Context) {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()

	if !db.HasTable("preupload") && !db.HasTable("upload") {
		db.AutoMigrate(&images.Preupload{}, &images.Upload{})
		ginContext.JSON(http.StatusOK, gin.H{"message": "Migrate compliet"})
		return
	}
	ginContext.JSON(http.StatusOK, gin.H{"message": "Noting change"})

}
