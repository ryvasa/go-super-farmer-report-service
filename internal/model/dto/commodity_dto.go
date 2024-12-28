package dto

type CommodityCreateDTO struct {
	Name        string `json:"name" validate:"required,min=3,max=255"`
	Code        string `json:"code" validate:"required,min=3,max=255"`
	Description string `json:"description" validate:"required,min=3,max=255"`
	Duration    string `json:"duration" validate:"required,min=1"`
}

type CommodityUpdateDTO struct {
	Name        *string `json:"name,omitempty" validate:"omitempty,min=3,max=255"`
	Code        *string `json:"code,omitempty" validate:"omitempty,min=3,max=255"`
	Description *string `json:"description,omitempty" validate:"omitempty,min=3,max=255"`
	Duration    *string `json:"duration,omitempty" validate:"omitempty,min=1"`
}
