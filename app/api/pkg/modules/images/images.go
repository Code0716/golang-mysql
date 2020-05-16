package images

import (
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	//"reflect"
	"encoding/json"
	"sync"
	"time"

	"../../../constants"
	"../../db"
	"../encoding"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	// "github.com/rwcarlsen/goexif/exif"
)

type (

	// Preupload model
	Preupload struct {
		gorm.Model
		Title    string     `json:"Title"`
		Path     string     `json:"Path"`
		ShotDate *time.Time `gorm:"column:ShotDate" json:"ShotDate"`
	}

	preuploads []Preupload

	// ↑↓共通化することを検討する。
	// db.Table("table name")でtableを指定できるようだ。 5/15

	//Upload model
	Upload struct {
		gorm.Model
		Title    string     `json:"Title"`
		Path     string     `json:"Path"`
		ShotDate *time.Time `gorm:"column:ShotDate" json:"ShotDate"`
	}

	uploads []Upload

	// PreImageController struct
	PreImageController struct{}
	// ImageController struct
	ImageController struct{}

	operateFileData interface {
		CreateFileInfo()
		GetFile(id string)
		DeleteFileInfo()
	}

	operateAllData interface {
		GetAllFileInfo()
		//	DeleteAllFileInfo()
	}
)

// 色々試したくて冗長になっている。

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

// GetFile 個々のファイルを取得
func (image *Preupload) GetFile(id string) {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Where("id = ?", id).Find(&image)
}

// GetFile 個々のファイルを取得
func (image *Upload) GetFile(id string) {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Where("id = ?", id).Find(&image)
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

// DeleteFileInfo func delete Preupload info
func (image *Preupload) DeleteFileInfo() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Unscoped().Delete(&image)
}

// DeleteFileInfo func delete upload info
func (image *Upload) DeleteFileInfo() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Unscoped().Delete(&image)
}

//------------------------------------------------------
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

	//lastMod := ginContext.Request.FormValue("date")
	lastMod := ginContext.PostForm("date")
	lastModBytes := []byte(lastMod)
	lastModArr := make([]*time.Time, len(lastMod))

	if err := json.Unmarshal(lastModBytes, &lastModArr); err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	// return json
	jsonData := make([]map[string]interface{}, len(files))

	wg := new(sync.WaitGroup)
	// save images
	for index, file := range files {
		wg.Add(2)

		go func(f *multipart.FileHeader) {
			defer wg.Done()

			if err := ginContext.SaveUploadedFile(f, constants.PreImagePath+f.Filename); err != nil {
				ginContext.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
				return
			}

		}(file)

		go func(i int, f *multipart.FileHeader) {
			defer wg.Done()
			// regist to db file info
			newImage := Preupload{Title: f.Filename, Path: constants.PreImagePath + f.Filename, ShotDate: lastModArr[i]}
			saveImageInfo(&newImage)
			jsonData[i] = map[string]interface{}{"info": newImage}
		}(index, file)

	}

	wg.Wait()
	// send josn to front
	ginContext.JSON(http.StatusOK, jsonData)
}

// GetAll Preupload image and registed db data
func (pre PreImageController) GetAll(ginContext *gin.Context) {

	var images preuploads

	getAllImageInfo(&images)

	jsonData := make([]map[string]interface{}, len(images))

	for index, file := range images {
		jsonData[index] = map[string]interface{}{"info": file}
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

	deleteImageInfo(&deleteData)

}

// ComitUpload func
func (pre PreImageController) ComitUpload(ginContext *gin.Context) {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	var images preuploads

	//現在pre_uploadデータベースにあるものを取得
	getAllImageInfo(&images)

	// goroutine で記述する意味がないかも。
	finPreUpload := make(chan bool)

	for _, image := range images {
		go func(img Preupload) {

			if err := preuploadToUpload(db, img); err != nil {
				ginContext.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			}
			finPreUpload <- true
		}(image)
	}

	<-finPreUpload

	ginContext.JSON(http.StatusOK, gin.H{"message": "Upload Complet"})

}

func preuploadToUpload(db *gorm.DB, image Preupload) error {

	return db.Transaction(func(tx *gorm.DB) error {
		// 削除する。
		if err := tx.Unscoped().Delete(&image).Error; err != nil {
			// エラーを返した場合はロールバックされます
			return err
		}
		// uploadに保存
		if err := tx.Create(&Upload{Title: image.Title, Path: constants.ImagePath + image.Title, ShotDate: image.ShotDate}).Error; err != nil {
			return err
		}
		// ファイル移動
		if err := os.Rename(image.Path, constants.ImagePath+image.Title); err != nil {
			return err
		}
		// nilを返すとコミットされる
		return nil
	})

}

// GetAll upload image and registed db data
func (upload ImageController) GetAll(ginContext *gin.Context) {

	var images uploads

	getAllImageInfo(&images)

	jsonData := make([]map[string]interface{}, len(images))

	for index, file := range images {
		jsonData[index] = map[string]interface{}{"info": file}
	}

	ginContext.JSON(http.StatusOK, jsonData)
}

// GetFile single
func (upload ImageController) GetFile(ginContext *gin.Context) {
	id := ginContext.Param("id")

	var image Upload

	getImage(&image, id)

	jsonData := make(map[string]string)

	base64gify := encoding.EncodeBase64(constants.ImagePath, image.Title)
	jsonData = map[string]string{"img": base64gify}

	ginContext.JSON(http.StatusOK, jsonData)
}

// DeleteAllImages images and info
func DeleteAllImages(ginContext *gin.Context) {
	flag := ginContext.Param("flag")
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()

	var images uploads

	db.Table(flag).Find(&images)

	for _, image := range images {
		if err := deleteImageAndInfo(db, image, flag); err != nil {
			ginContext.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "path": image})
		}
	}

	ginContext.JSON(http.StatusOK, gin.H{"message": "Deleted All " + flag + " images"})
}

func deleteImageAndInfo(db *gorm.DB, image Upload, flag string) error {

	return db.Transaction(func(tx *gorm.DB) error {
		// 削除する。
		if err := tx.Table(flag).Unscoped().Delete(&image).Error; err != nil {
			return err
		}

		// ファイル削除
		if err := os.Remove(image.Path); err != nil {
			return err
		}
		return nil
	})

}
