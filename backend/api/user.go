package main

import uuid "github.com/google/uuid"

type User struct {
	Name     string    `json:"name"`
	Lastname string    `json:"lastname"`
	Password string    `json:"password"`
	UserId   uuid.UUID `json:"id"`
}
