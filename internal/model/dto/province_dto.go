package dto

type ProvinceCreateDTO struct {
	Name string `json:"name" validate:"required"`
}

type ProvinceUpdateDTO struct {
	Name string `json:"name" validate:"required"`
}
