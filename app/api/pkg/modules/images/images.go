package images

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"time"

	"../../../constants"
	"../../db"

	"github.com/gin-gonic/gin"
)

const preImagePath = "../images/preupload/"

type (
	createFileData interface {
		CreateFileInfo()
	}

	preupload struct {
		ID     int       `json:"id"`
		Title  string    `json:"title"`
		Create time.Time `json:"create"`
		Path   string    `json:"path"`
	}

	upload struct {
		ID     int       `json:"id"`
		Title  string    `json:"title"`
		Create time.Time `json:"create"`
		Path   string    `json:"path"`
	}
)

// save  preupload info
func (image *preupload) CreateFileInfo() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Create(&image)
}

// save  upload info
func (image *upload) CreateFileInfo() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Create(&image)
}

func saveImageInfo(saveFunc createFileData) {
	saveFunc.CreateFileInfo()
}

// GetPreUploadImg save preupload image and regist db
func GetPreUploadImg(ginContext *gin.Context) {
	form, _ := ginContext.MultipartForm()
	files := form.File["images"]

	jsonData := make([]map[string]interface{}, len(files))

	// save images
	for index, file := range files {
		err := ginContext.SaveUploadedFile(file, preImagePath+file.Filename)

		if err != nil {
			ginContext.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		now := time.Now()
		// regist to db file info
		newImage := preupload{Title: file.Filename, Create: now, Path: preImagePath + file.Filename}
		saveImageInfo(&newImage)

		base64gify := encodeBase64(preImagePath, file.Filename)

		jsonData[index] = map[string]interface{}{"img": base64gify, "info": newImage}

	}

	// send to josn to front
	ginContext.JSON(http.StatusOK, jsonData)
}

// DeletePreUploadImage delete preupload image and registed db data
func DeletePreUploadImage(ginContext *gin.Context) {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()

	var deleteData preupload

	if err := ginContext.BindJSON(&deleteData); err != nil {
		ginContext.String(http.StatusBadRequest, "Request is failed: "+err.Error())
	}

	if err := os.Remove(deleteData.Path); err != nil {
		fmt.Println(err)
	}

	fmt.Println(deleteData)

	db.Delete(&deleteData)

}

// get save image and Encode to base64
func encodeBase64(savePath string, fileNama string) string {
	file, err := os.Open(savePath + fileNama)

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
