package dto

type CityCreateDTO struct {
	Name       string `json:"name" validate:"required"`
	ProvinceID int64  `json:"province_id" validate:"required"`
}

type CityUpdateDTO struct {
	Name       string `json:"name,omitempty" validate:"min=1,max=255"`
	ProvinceID int64  `json:"province_id,omitempty"`
}
