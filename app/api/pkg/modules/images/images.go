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

const preImagePath = "../../images/preupload/"

//const imagePath = "../images/upload/"

type (
	operateFileData interface {
		CreateFileInfo()
		GetFile(id string)
		DeleteFileInfo()
	}

	operateAllData interface {
		GetAllFileInfo()
	}

	// Preupload model
	Preupload struct {
		gorm.Model
		Title string `json:"Title"`
		Path  string `json:"Path"`
	}

	//Upload model
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

// CreateFileInfo func Preupload info save
func (image *Preupload) CreateFileInfo() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Create(&image)
}

// CreateFileInfo save  upload info
func (image *Upload) CreateFileInfo() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Create(&image)
}

func (image *Preupload) GetFile(id string) {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Where("id = ?", id).Find(&image)
}

func (image *Upload) GetFile(id string) {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Where("id = ?", id).Find(&image)
}

// DeleteFileInfo func delete Preupload info
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

// DeleteFileInfo func delete upload info
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

func getImage(opreate operateFileData, id string) {
	operateFileData.GetFile(opreate, id)
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

		jsonData[index] = map[string]interface{}{"info": newImage}

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
		//画像は一つづつ読み込むようにする。
		jsonData[index] = map[string]interface{}{"info": file}

	}

	ginContext.JSON(http.StatusOK, jsonData)
}

// GetFile single
func (pre PreImageController) GetFile(ginContext *gin.Context) {
	id := ginContext.Param("id")

	var image Preupload

	getImage(&image, id)

	jsonData := make(map[string]string)

	base64gify := encodeBase64(preImagePath, image.Title)
	jsonData = map[string]string{"img": base64gify}

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

// ComitUpload func
/*func ComitUpload(ginContext *gin.Context) {
	// TODO
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()

	preuploadToUpload(db)
}

func preuploadToUpload(db *gorm.DB, iamges *preuploads) error {
	return db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Delete().Error; err != nil {
			// エラーを返した場合はロールバックされます
			return err
		}

		if err := tx.Create(&Animal{Name: "Lion"}).Error; err != nil {
			return err
		}

		// nilを返すとコミットされる
		return nil
	})
}
*/
