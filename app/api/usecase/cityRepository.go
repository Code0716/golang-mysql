package usecase

import "../domain"

type CityRepository interface {
	GetAll() (domain.Cities, error)
}
