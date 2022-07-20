package migrate

import (
	"net/http"

	"github.com/Code0716/golang-mysql/app/api/constants"
	"github.com/Code0716/golang-mysql/app/api/pkg/db"
	"github.com/Code0716/golang-mysql/app/api/pkg/modules/images"
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
