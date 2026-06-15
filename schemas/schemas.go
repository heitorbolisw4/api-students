package schemas

import (
	"gorm.io/gorm"
)

type Student struct {
	gorm.Model
	Name     string `json:"name"`
	Cpf      string `json:"cpf"`
	Mail     string `json:"mail"`
	Age      int    `json:"age"`
	IsActive bool   `json:"isActive"`
}
