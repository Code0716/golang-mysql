package images

import (
	"encoding/base64"
	//"fmt"
	"net/http"
	"os"
	"time"
	//"mime/multipart"

	"../../../constants"
	"../../db"

	"github.com/gin-gonic/gin"
)

const preImagePath = "../images/preupload/"

type (
	imagesFuncs interface {
		// 全件取得
		//ReadAll()
		// 一件取得
		//ReadSingle(path string)

		CreateFileInfo(imageInfo preupload)

		//Update(ginContext *gin.Context, name string)

		//Delete(name string, path string)
	}

	preupload struct {
		ID     int       `json:"id"`
		Title  string    `json:"title"`
		Create time.Time `json:"create"`
		Path   string    `json:"path"`
	}

	preuploadArray []preupload
)

// save  upload info
func (image *preupload) CreateFileInfo(imageInfo preupload) {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	//　db.AutoMigrate(&cities)
	db.Create(&imageInfo)
}

// get  upload info
/*func (images *preuploadArray) ReadAll() *preuploadArray {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	//　db.AutoMigrate(&cities)
	db.Select("id,title,create,path").Find(&images)
	return images
}*/

func saveImageInfo(imageInfo preupload) {
	var saveFunc imagesFuncs
	var image preupload
	saveFunc = &image
	saveFunc.CreateFileInfo(imageInfo)
}

// get save image and Encode to base64
func encode(fileNama string) string {
	file, err := os.Open(preImagePath + fileNama)

	if err != nil {
		return err.Error()

	}

	defer file.Close()

	fi, _ := file.Stat() // interface
	size := fi.Size()    // file size

	data := make([]byte, size)
	file.Read(data)
	return base64.StdEncoding.EncodeToString(data)
}

// GetPreUploadImg 画像仮保存
func GetPreUploadImg(ginContext *gin.Context) {
	form, _ := ginContext.MultipartForm()
	files := form.File["images"]

	now := time.Now()

	// base64 string
	base64gify := make([]string, len(files))

	// save images
	for index, file := range files {
		err := ginContext.SaveUploadedFile(file, preImagePath+file.Filename)

		if err != nil {
			ginContext.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
		//
		newImage := preupload{Title: file.Filename, Create: now, Path: preImagePath + file.Filename}
		saveImageInfo(newImage)
		base64gify[index] = encode(file.Filename)
	}

	// send to josn to front
	ginContext.JSON(http.StatusOK, base64gify)
}
