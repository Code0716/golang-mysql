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

// GetAll []City
func (cities *Cities) GetAll() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()

	//　db.AutoMigrate(&cities)
	db.Select("id,name,code,population").Find(&cities)
}

// GetAll []Country
func (countries *Countries) GetAll() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()

	db.Select("code,name,continent").Find(&countries)
}

// GetAll []Continent
func (countryContinents *CountryContinents) GetAll() {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Raw("SELECT DISTINCT continent FROM country").Scan(&countryContinents)
}

// GetSingle City
func (cities *Cities) GetSingle(name string) {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()

	db.Select("id,name,code,population").Find(&cities, "name = ?", name)
}

// GetSingle Country
func (countries *Countries) GetSingle(name string) {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()

	db.Select("code,name,continent").Find(&countries, "name = ?", name)
}

// GetSingle Continents
func (countryContinents *CountryContinents) GetSingle(name string) {
	db := db.ConnectMySQL(constants.DBWorld)
	defer db.Close()
	db.Select("continent").Find(&countryContinents, "name = ?", name)
}

// GetList from db
func GetList(worlddb WorldDB) {
	WorldDB.GetAll(worlddb)
}

// GetRow from db
func GetRow(worlddb WorldDB, name string) {
	WorldDB.GetSingle(worlddb, name)
}

// GetCities 街一覧
func GetCities(ginContext *gin.Context) {
	// struct
	var cities Cities
	GetList(&cities)
	ginContext.JSON(http.StatusOK, cities)
}

// GetCountries 国一覧
func GetCountries(ginContext *gin.Context) {
	// struct
	var countries Countries
	GetList(&countries)
	ginContext.JSON(http.StatusOK, countries)
}

// GetContinentsList 大陸一覧.
func GetContinentsList(ginContext *gin.Context) {

	var continents CountryContinents
	GetList(&continents)
	ginContext.JSON(http.StatusOK, continents)
}

// GetCity 街
func GetCity(ginContext *gin.Context) {
	// struct
	var cities Cities

	name := ginContext.Param("name")

	GetRow(&cities, name)
	ginContext.JSON(http.StatusOK, cities)
}

// GetCountry 国Get
func GetCountry(ginContext *gin.Context) {

	// struct
	var countries Countries

	name := ginContext.Param("name")

	GetRow(&countries, name)
	ginContext.JSON(http.StatusOK, countries)
}

// GetContinent func
func GetContinent(ginContext *gin.Context) {
	// struct
	var continents CountryContinents

	name := ginContext.Param("name")

	GetRow(&continents, name)
	ginContext.JSON(http.StatusOK, continents)
}
