package images

import (
	"encoding/base64"
	"fmt"
	//"mime/multipart"
	"net/http"
	"os"
	//"reflect"
	"time"

	"../../../constants"
	"../../db"

	"github.com/gin-gonic/gin"
)

const preImagePath = "../images/preupload/"

type (
	operateFileData interface {
		CreateFileInfo()
		DeleteFileInfo()
	}

	operateAllData interface {
		GetAllFileInfo()
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

	preuploads []preupload
	uploads    []upload
	// PreImageController struct
	PreImageController struct{}
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

// delete preupload info
func (image *preupload) DeleteFileInfo() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Delete(&image)
}

// read all preupload info
func (images *preuploads) GetAllFileInfo() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Select("id,title,path").Find(&images)
}

// read all upload info
func (images *uploads) GetAllFileInfo() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Find(&images)
}

// delete upload info
func (image *upload) DeleteFileInfo() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Delete(&image)
}

func saveImageInfo(operate operateFileData) {
	operate.CreateFileInfo()
}

func deleteImageInfo(operate operateFileData) {
	operate.DeleteFileInfo()
}

func getAllImageInfo(operate operateAllData) {
	operateAllData.GetAllFileInfo(operate)
}

// Upload preupload image and regist db
func (pre PreImageController) Upload(ginContext *gin.Context) {
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

// ↑共通化できるか？↓

// GetAll preupload image and registed db data
func (pre PreImageController) GetAll(ginContext *gin.Context) {

	var images preuploads

	getAllImageInfo(&images)

	jsonData := make([]map[string]interface{}, len(images))

	for index, file := range images {
		base64gify := encodeBase64(preImagePath, file.Title)
		jsonData[index] = map[string]interface{}{"img": base64gify, "info": file}

	}

	ginContext.JSON(http.StatusOK, jsonData)
}

// Delete preupload image and registed db data
func (pre PreImageController) Delete(ginContext *gin.Context) {

	var deleteData preupload

	if err := ginContext.BindJSON(&deleteData); err != nil {
		ginContext.JSON(http.StatusBadRequest, err.Error())
	}

	if err := os.Remove(deleteData.Path); err != nil {
		fmt.Println(err)
	}

	fmt.Println(deleteData)

	deleteImageInfo(&deleteData)

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
