package dto

import "github.com/google/uuid"

type RoleCreateDTO struct {
	Name string `json:"name" validate:"required,min=3,max=255"`
}

type RoleResponseDTO struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
