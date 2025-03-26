package main

import uuid "github.com/google/uuid"

type User struct {
	Name     string    `json:"name" validate:"required,alpha"`
	Lastname string    `json:"lastname" validate:"required,alpha"`
	Email    string    `json:"email" validate:"required,email"`
	Password string    `json:"password" validate:"required,min=8"`
	UserId   uuid.UUID `json:"id"`
}
