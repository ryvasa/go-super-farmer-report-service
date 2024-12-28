package dto

type RegionCreateDto struct {
	CityID     int64 `json:"city_id" validate:"min=1"`
	ProvinceID int64 `json:"province_id"  validate:"min=1,max=255"`
}
