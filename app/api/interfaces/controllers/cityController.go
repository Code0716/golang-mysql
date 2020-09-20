package controllers

import (
	"../../usecase"
	"../database"
)

type CityController struct {
	Interactor usecase.CityInteractor
}

func NewCityController(sqlHandler database.SqlHandler) *CityController {
	return &CityController{
		Interactor: usecase.CityInteractor{
			CityRepository: &database.CityRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *CityController) GetAll(c Context) {
	cities, err := controller.Interactor.GetAllCity()
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, cities)
}
