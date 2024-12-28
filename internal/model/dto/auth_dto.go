package dto

type AuthDTO struct {
	Email    string `json:"email" validate:"required,email,min=3,max=255"`
	Password string `json:"password" validate:"required,min=6,max=255"`
}

type AuthResponseDTO struct {
	User  *UserResponseDTO `json:"user"`
	Token string           `json:"token"`
}

type AuthSendDTO struct {
	Email string `json:"email" validate:"required,email,min=3,max=255"`
}

type AuthVerifyDTO struct {
	Email string `json:"email" validate:"required,email,min=3,max=255"`
	OTP   string `json:"otp" validate:"required,min=6,max=6"`
}
