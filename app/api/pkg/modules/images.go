package modules

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetPreUploadImg 画像仮保存
func GetPreUploadImg(ginContext *gin.Context) {

	form, _ := ginContext.MultipartForm()
	files := form.File["images"]

	fmt.Println(files)
	for _, file := range files {
		err := ginContext.SaveUploadedFile(file, "../images/preupload/"+file.Filename)
		if err != nil {
			ginContext.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}

	ginContext.JSON(http.StatusOK, gin.H{"message": "success!!"})

}
