package dto

import (
	"time"

	"github.com/google/uuid"
)

type UserCreateDTO struct {
	Name     string `json:"name" validate:"min=3,max=255"`
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"min=6,max=255"`
	Phone    string `json:"phone,omitempty" validate:"omitempty,min=3,max=20"`
}

type UserUpdateDTO struct {
	Name     string `json:"name" validate:"omitempty,min=3,max=255"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password" validate:"omitempty,min=6,max=255"`
	Phone    string `json:"phone" validate:"omitempty,min=3,max=20"`
	RoleID   int64  `json:"role_id" validate:"omitempty,min=1,max=2"`
}

type UserResponseDTO struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	Phone     *string    `json:"phone,omitempty"`
	Password  string     `json:"password,omitempty"`
	RoleID    int64      `json:"role_id,omitempty"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty"`
}
