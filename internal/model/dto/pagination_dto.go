package dto

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type PaginationDTO struct {
	Limit  int            `json:"limit" form:"limit"`
	Page   int            `json:"page" form:"page"`
	Sort   string         `json:"sort" form:"sort"`
	Filter ParamFilterDTO `json:"filter" form:"filter"`
}
type ParamFilterDTO struct {
	UserName      string     `json:"user_name" form:"user_name"`
	CommodityName string     `json:"commodity_name" form:"commodity_name"`
	CityName      string     `json:"city_name" form:"city_name"`
	CityID        *int64     `json:"city_id" form:"city_id"`
	CommodityID   *uuid.UUID `json:"commodity_id" form:"commodity_id"`
	StartDate     time.Time  `json:"start_date" form:"start_date"`
	EndDate       time.Time  `json:"end_date" form:"end_date"`
}

type PaginationResponseDTO struct {
	TotalRows  int64       `json:"total_rows"`
	TotalPages int         `json:"total_pages"`
	Page       int         `json:"page"`
	Limit      int         `json:"limit"`
	Data       interface{} `json:"data"`
}

func (p *PaginationDTO) Validate() error {
	if p.Page < 1 {
		return errors.New("page must be greater than 0")
	}
	if p.Limit < 1 {
		return errors.New("limit must be greater than 0")
	}
	if p.Limit > 100 {
		return errors.New("limit must not exceed 100")
	}
	return nil
}
