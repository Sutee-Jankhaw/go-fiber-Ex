package models

import (
	"time"

	"gorm.io/gorm"
)

type UserPayLoad struct {
	EmployeeID uint   `json:"employee_id"`
	Name       string `json:"name"`
	LastName   string `json:"lastname"`
	Birthday   string `json:"birthday"`
	Age        int    `json:"age"`
	Email      string `json:"email"`
	Tel        string `json:"tel"`
}

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

type UserRes struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	EmployeeID uint   `json:"employee_id"`
	Generation string `json:generation`
}
type ResultData struct {
	Data         []UserRes `json:"data"`
	Count        int       `json:"count"`
	GenZ         int       `json:"gen_z"`
	GenY         int       `json:"gen_y"`
	GenX         int       `json:"gen_x"`
	BabyBloomer  int       `json:"baby_bloomer"`
	GIgeneration int       `json:"gi_generation"`
}
