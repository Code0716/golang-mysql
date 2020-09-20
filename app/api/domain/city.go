package domain

type (
	City struct {
		ID         int
		Name       string
		Code       string
		District   string
		Population int
	}
	Cities []City
)
