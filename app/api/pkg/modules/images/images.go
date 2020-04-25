package images

import (
	"encoding/base64"
	"fmt"
	//"mime/multipart"
	"net/http"
	"os"
	//"reflect"

	"../../../constants"
	"../../db"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

const preImagePath = "../images/preupload/"

//const imagePath = "../images/upload/"

type (
	operateFileData interface {
		CreateFileInfo()
		DeleteFileInfo()
	}

	operateAllData interface {
		GetAllFileInfo()
	}

	Preupload struct {
		gorm.Model
		Title string `json:"Title"`
		Path  string `json:"Path"`
	}

	Upload struct {
		gorm.Model
		Title string `json:"Title"`
		Path  string `json:"Path"`
	}

	preuploads []Preupload
	uploads    []Upload
	// PreImageController struct
	PreImageController struct{}
)

// save  Preupload info
func (image *Preupload) CreateFileInfo() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Create(&image)
}

// save  upload info
func (image *Upload) CreateFileInfo() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Create(&image)
}

// delete Preupload info
func (image *Preupload) DeleteFileInfo() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Delete(&image)
}

// read all Preupload info
func (images *preuploads) GetAllFileInfo() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Find(&images)
}

// read all upload info
func (images *uploads) GetAllFileInfo() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Find(&images)
}

// delete upload info
func (image *Upload) DeleteFileInfo() {
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

// Upload Preupload image and regist db
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
		// regist to db file info
		newImage := Preupload{Title: file.Filename, Path: preImagePath + file.Filename}
		saveImageInfo(&newImage)

		base64gify := encodeBase64(preImagePath, file.Filename)

		jsonData[index] = map[string]interface{}{"img": base64gify, "info": newImage}

	}

	// send to josn to front
	ginContext.JSON(http.StatusOK, jsonData)
}

// ↑ 共通化できるか？ -----------------------------------------------------------------↓

// GetAll Preupload image and registed db data
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

// Delete Preupload image and registed db data
func (pre PreImageController) Delete(ginContext *gin.Context) {

	var deleteData Preupload

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
