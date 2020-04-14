package modules

import (
	"../../constants"
	"../db"
)

// WorldDB interface
type WorldDB interface {
	// 全件取得
	GetAll()
	// 一件取得
	GetSingle(name string)
}

// City struct
// gorm.Modelと記述するとcreated_at、updated_at、deleted_atが定義される。
type City struct {
	//gorm.Model
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Code       string `json:"code"`
	Population string `json:"population"`
}

// Cities []City
type Cities []City

/*GetAll()を共通化できないか？ーーーーーーーーーーーーーーーーーーーーーー*/

// GetAll []City
func (cities *Cities) GetAll() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.LogMode(true)
	// 勝手に複数形になるのを抑制
	db.SingularTable(true)

	//　db.AutoMigrate(&city)
	db.Select("id,name,code,population").Find(&cities)
}

// GetSingle City
func (cities *Cities) GetSingle(name string) {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.LogMode(true)
	// 勝手に複数形になるのを抑制
	db.SingularTable(true)
	db.Select("id,name,code,population").Find(&cities, "name = ?", name)
}

// Country struct
type Country struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	Continent string `json:"continent"`
}

// Countries []Country
type Countries []Country

// GetAll []Country
func (countries *Countries) GetAll() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.LogMode(true)
	// 勝手に複数形になるのを抑制
	db.SingularTable(true)
	db.Select("code,name,continent").Find(&countries)
}

// GetSingle Country
func (countries *Countries) GetSingle(name string) {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.LogMode(true)
	// 勝手に複数形になるのを抑制
	db.SingularTable(true)
	db.Select("code,name,continent").Find(&countries, "name = ?", name)
}

// Continent types
type Continent struct {
	Continent string `json:"continent"`
}

// CountryContinents array
type CountryContinents []Continent

// GetContinentsDB func
// count
func GetContinentsDB(path string) CountryContinents {
	var countryContinents CountryContinents
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.LogMode(true)
	// 勝手に複数形になるのを抑制
	db.SingularTable(true)
	db.Raw("SELECT DISTINCT continent FROM country").Scan(&countryContinents)
	return countryContinents
}