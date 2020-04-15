package images

import (
	"encoding/base64"
	"fmt"
	"net/http"
	//"mime/multipart"
	"os"

	"github.com/gin-gonic/gin"
)

type imagesFunc interface {
	// 全件取得
	GetAll()
	// 一件取得
	GetSingle(path string)

	Update(ginContext *gin.Context, name string)

	Create(ginContext *gin.Context)

	Delete(name string, path string)
}

const preImagePath = "../images/preupload/"

type images struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Create string `json:"create"`
	Path   string `json:"path"`
}

// GetPreUploadImg 画像仮保存
func GetPreUploadImg(ginContext *gin.Context) {

	form, _ := ginContext.MultipartForm()
	files := form.File["images"]

	base64gify := make([]string, len(files))

	// save images
	for index, file := range files {
		err := ginContext.SaveUploadedFile(file, preImagePath+file.Filename)

		if err != nil {
			ginContext.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
		base64gify[index] = encode(file.Filename)
	}

	ginContext.JSON(http.StatusOK, base64gify)
}

// get save image and Encode to base64
func encode(fileNama string) string {
	fmt.Println(fileNama)
	file, _ := os.Open(preImagePath + fileNama)
	defer file.Close()

	fi, _ := file.Stat() // interface
	size := fi.Size()    // file size

	data := make([]byte, size)
	file.Read(data)
	return base64.StdEncoding.EncodeToString(data)
}
