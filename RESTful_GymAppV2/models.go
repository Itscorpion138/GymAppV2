package main

import (
	"time"

	"github.com/google/uuid"
)

type Users struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	LastName   string    `json:"lastname"`
	Age        int       `json:"age"`
	Height     int       `json:"height"`
	Weight     float64   `json:"weight"`
	Gender     string    `json:"gender"`
	Plan       string    `json:"plan"`
	Created    time.Time `json:"created"`
	SkillLevel string    `json:"skilllevel"`
	Email      string    `json:"email"`
}

type Coach struct {
	ID         uuid.UUID `json:"id"`
	Name       string    `json:"name"`
	LastName   string    `json:"lastname"`
	Gender     string    `json:"gender"`
	Age        int       `json:"age"`
	Email      string    `json:"email"`
	Cost       float64   `json:"cost"`
	Verified   bool      `json:"verified"`
	ProfilePic string    `json:"profile_pic"`
	Bio        string    `json:"bio"`
}

type Program struct {
	ID uuid.UUID `json:"id"`
}
