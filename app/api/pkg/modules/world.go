package modules

import (
	"net/http"

	"../../constants"
	"../db"

	"github.com/gin-gonic/gin"
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

// getContinentsDB func
// count
func getContinentsDB(path string) CountryContinents {
	var countryContinents CountryContinents
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.LogMode(true)
	// 勝手に複数形になるのを抑制
	db.SingularTable(true)
	db.Raw("SELECT DISTINCT continent FROM country").Scan(&countryContinents)
	return countryContinents
}

// 街一覧
func GetCities(ginContext *gin.Context) {
	// interface
	var world WorldDB
	// struct
	var cities Cities

	world = &cities
	world.GetAll()
	ginContext.JSON(http.StatusOK, cities)
}

//　まち
func GetCity(ginContext *gin.Context) {
	// interface
	var world WorldDB
	// struct
	var cities Cities

	name := ginContext.Param("name")

	world = &cities
	world.GetSingle(name)
	ginContext.JSON(http.StatusOK, cities)
}

// 国一覧
func GetCountries(ginContext *gin.Context) {
	// interface
	var world WorldDB
	// struct
	var countries Countries

	world = &countries

	world.GetAll()
	ginContext.JSON(http.StatusOK, countries)
}

// 国Get
func GetCountry(ginContext *gin.Context) {
	// interface
	var world WorldDB
	// struct
	var countries Countries

	name := ginContext.Param("name")

	world = &countries
	world.GetSingle(name)
	ginContext.JSON(http.StatusOK, countries)
}

// 大陸一覧.
func GetContinentsList(ginContext *gin.Context) {
	data := getContinentsDB("GetCountDB")
	ginContext.JSON(http.StatusOK, data)
}
