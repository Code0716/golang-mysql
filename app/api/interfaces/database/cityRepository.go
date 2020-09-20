package database

import "../../domain"

type CityRepository struct {
	SqlHandler
}

func (repo *CityRepository) GetAll() (cities domain.Cities, err error) {
	rows, err := repo.Query("SELECT * FROM city")
	defer rows.Close()
	if err != nil {
		return
	}
	var result domain.Cities

	for rows.Next() {
		var id int
		var name string
		var code string
		var district string
		var population int

		if err := rows.Scan(&id, &name, &code, &district, &population); err != nil {
			continue
		}
		city := domain.City{
			ID:         id,
			Name:       name,
			Code:       code,
			District:   district,
			Population: population,
		}
		result = append(result, city)
	}

	return result, err
}
