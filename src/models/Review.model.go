package models

import "github.com/google/uuid"

type Review struct {
	Base
	Content string
	MovieID uuid.UUID
	UserID  uuid.UUID `gorm:"index"`
	User    User
	Rating  float64
}
