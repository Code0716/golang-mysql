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

func savePreImageInfo(imageInfo preupload) {
	var saveFunc createFileData
	saveFunc = &imageInfo
	saveFunc.CreateFileInfo()
}

// save  upload info
func (image *upload) CreateFileInfo() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Create(&image)
}

func saveImageInfo(imageInfo upload) {
	var saveFunc createFileData
	saveFunc = &imageInfo
	saveFunc.CreateFileInfo()
}

// GetPreUploadImg save preupload image and regist db
func GetPreUploadImg(ginContext *gin.Context) {
	form, _ := ginContext.MultipartForm()
	files := form.File["images"]

	// base64 string
	base64gify := make([]string, len(files))

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
		savePreImageInfo(newImage)

		base64gify[index] = encodePreupload(file.Filename)
	}

	// send to josn to front
	ginContext.JSON(http.StatusOK, base64gify)
}

// get save image and Encode to base64
func encodePreupload(fileNama string) string {
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
