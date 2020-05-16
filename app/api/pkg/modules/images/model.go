package images

import (
	"../../../constants"
	"../../db"
	"time"

	"github.com/jinzhu/gorm"
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
	// db.Table("table name")でtableを指定でき、共通化できる.

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
		createFileInfo()
		getFile(id string)
		deleteFileInfo()
	}

	operateAllData interface {
		getAllFileInfo()
		//	DeleteAllFileInfo()
	}
)

// 色々試したくて冗長になっている。

// createFileInfo func Preupload info save
func (image *Preupload) createFileInfo() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Create(&image)
}

// createFileInfo save  upload info
func (image *Upload) createFileInfo() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Create(&image)
}

// getFile 個々のファイルを取得
func (image *Preupload) getFile(id string) {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Where("id = ?", id).Find(&image)
}

// getFile 個々のファイルを取得
func (image *Upload) getFile(id string) {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Where("id = ?", id).Find(&image)
}

// read all Preupload info
func (images *preuploads) getAllFileInfo() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Find(&images)
}

// read all upload info
func (images *uploads) getAllFileInfo() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Find(&images)
}

// deleteFileInfo func delete Preupload info
func (image *Preupload) deleteFileInfo() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Unscoped().Delete(&image)
}

// deleteFileInfo func delete upload info
func (image *Upload) deleteFileInfo() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Unscoped().Delete(&image)
}

//------------------------------------------------------
func saveImageInfo(operate operateFileData) {
	operate.createFileInfo()
}

func deleteImageInfo(operate operateFileData) {
	operate.deleteFileInfo()
}

func getImage(opreate operateFileData, id string) {
	opreate.getFile(id)
}

func getAllImageInfo(operate operateAllData) {
	operate.getAllFileInfo()
}
