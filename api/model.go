package api

type User struct {
	Name    string `json:"name,omitempty" validate:"required"`
	Phone   string `json:"phone,omitempty" validate:"required"`
	Email   string `json:"email,omitempty" validate:"required"`
	Address string `json:"address,omitempty" validate:"required"`
}

type ParsePhone struct {
	Phone string `json:"phone,omitempty" validate:"required"`
}

type JsonResponse struct {
	Status  int    `json:"status,omitempty" validate:"required"`
	Message string `json:"message,omitempty" validate:"required"`
	Data    any    `json:"data,omitempty" validate:"required"`
}

type OTP struct {
	Phone string `json:"phone,omitempty" validate:"required"`
	Code  string `json:"code,omitempty" validate:"required"`
}
