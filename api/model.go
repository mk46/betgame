package api

import "time"

type User struct {
	Name    string `json:"name,omitempty" validate:"required" bson:"name"`
	Phone   string `json:"phone,omitempty" validate:"required" bson:"phone"`
	Email   string `json:"email,omitempty" validate:"required" bson:"email"`
	Balance int    `json:"balance,omitempty" validate:"required" bson:"balance"`
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

type Game struct {
	Name   string    `json:"name,omitempty" validate:"required"`
	Start  time.Time `json:"start,omitempty" validate:"required"`
	End    time.Time `json:"end,omitempty" validate:"required"`
	Result time.Time `json:"result,omitempty" validate:"required"`
}
