package modules

import (
	"net/http"

	"../../constants"
	"../db"

	"github.com/gin-gonic/gin"
)

type (
	// WorldDB interface
	WorldDB interface {
		// 全件取得
		GetAll()
		// 一件取得
		GetSingle(name string)
	}
	// City struct
	City struct {
		//gorm.Model
		ID         int    `json:"id"`
		Name       string `json:"name"`
		Code       string `json:"code"`
		Population string `json:"population"`
	}
	// Cities type
	Cities []City

	// Country struct
	Country struct {
		Code      string `json:"code"`
		Name      string `json:"name"`
		Continent string `json:"continent"`
	}

	// Countries []Country
	Countries []Country

	// Continent types
	Continent struct {
		Continent string `json:"continent"`
	}

	// CountryContinents array
	CountryContinents []Continent
)

// GetAll()を共通化できないか？ ーーーーーーーーーーーーーーーーーーーーーー
// 引数をinterface{}型にするか？

// GetAll []City
func (cities *Cities) GetAll() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()

	//　db.AutoMigrate(&cities)
	db.Select("id,name,code,population").Find(&cities)
}

// GetSingle City
func (cities *Cities) GetSingle(name string) {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()

	db.Select("id,name,code,population").Find(&cities, "name = ?", name)
}

// GetAll []Country
func (countries *Countries) GetAll() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()

	db.Select("code,name,continent").Find(&countries)
}

// GetSingle Country
func (countries *Countries) GetSingle(name string) {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()

	db.Select("code,name,continent").Find(&countries, "name = ?", name)
}

// getContinentsDB func
// count
func getContinentsDB(path string) CountryContinents {
	var countryContinents CountryContinents
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()

	db.Raw("SELECT DISTINCT continent FROM country").Scan(&countryContinents)
	return countryContinents
}

// GetCities 街一覧
func GetCities(ginContext *gin.Context) {
	// interface
	var world WorldDB
	// struct
	var cities Cities

	world = &cities
	world.GetAll()
	ginContext.JSON(http.StatusOK, cities)
}

// GetCity 街
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

// GetCountries 国一覧
func GetCountries(ginContext *gin.Context) {
	// interface
	var world WorldDB
	// struct
	var countries Countries

	world = &countries

	world.GetAll()
	ginContext.JSON(http.StatusOK, countries)
}

// GetCountry 国Get
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

// GetContinentsList 大陸一覧.
func GetContinentsList(ginContext *gin.Context) {
	data := getContinentsDB("GetCountDB")
	ginContext.JSON(http.StatusOK, data)
}
