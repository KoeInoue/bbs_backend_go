package domain

import (
	"time"
)

type Users []User

type User struct {
	Id        int    `gorm:"primary_key"`
	Name      string `gorm:"size:255"`
	Email     string `gorm:"size:255"`
	Password  string `gorm:"size:255"`
	CreatedAt time.Time
}
