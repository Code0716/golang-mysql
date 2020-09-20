package usecase

import (
	"../domain"
)

type CityInteractor struct {
	CityRepository CityRepository
}

// CreateFileInfo save  upload info
func (interactor *CityInteractor) GetAllCity() (cities domain.Cities, err error) {
	cities, err = interactor.CityRepository.GetAll()
	return
}
