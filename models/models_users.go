package models

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	EmployeeID uint      `json:"employee_id"`
	Name       string    `json:"name"`
	LastName   string    `json:"lastname"`
	BirthDay   time.Time `json:"birthday" gorm:"type:date"`
	Age        int       `json:"age"`
	Email      string    `json:"email"`
	Tel        string    `json:"tel"`
}
