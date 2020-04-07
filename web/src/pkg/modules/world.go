package modules

import (
	"../../constants"
	"../db"
)

// City struct
// gorm.Modelと記述するとcreated_at、updated_at、deleted_atが定義される。
type City struct {
	//gorm.Model
	ID         int    `json:"id"`
	Name       string `json:"name"`
	District   string `json:"district"`
	Population int    `json:"population"`
}

// Cities []City
type Cities []City

// Country struct
type Country struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	Continent string `json:"continent"`
}

// Countries []Country
type Countries []Country

// WorldDB interface
type WorldDB interface {
	// 全件取得
	GetAll()
	// 一件取得
	//GetSingle()
}

/*GetAll()を共通化できないか？ーーーーーーーーーーーーーーーーーーーーーー*/

// GetAll []City
func (cities *Cities) GetAll() {
	db := db.ConnectMySQL(constants.DBTableWorld)
	defer db.Close()
	db.LogMode(true)
	// 勝手に複数形になるのを抑制
	db.SingularTable(true)

	//　db.AutoMigrate(&city)
	//　初期データとして投入したものがdb.Find(&city)で取得出来ないのでこの形にした。
	db.Select("id,name,district,population").Find(&cities)
}

// GetAll []Country
func (countrys *Countries) GetAll() {
	db := db.ConnectMySQL(constants.DBTableWorld)
	defer db.Close()
	db.LogMode(true)
	// 勝手に複数形になるのを抑制
	db.SingularTable(true)

	db.Select("code,name,continent").Find(&countrys)
}

// GetAll func
func GetAll(world WorldDB) {
	world.GetAll()
}
