package db

import (
	"github.com/rs/zerolog/log"

	"gorm.io/driver/sqlite"
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

type StudentHandler struct {
	DB *gorm.DB
}

func Init() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("student.db"), &gorm.Config{})

	if err != nil {
		log.Error().Msg("Failed to connect to the database.")
	}
	db.AutoMigrate(&Student{})

	return db

}

func NewStudentHandler(db *gorm.DB) *StudentHandler {
	return &StudentHandler{DB: db}
}

func (s *StudentHandler) AddStudent(student Student) error {
	db := Init()

	if result := db.Create(&student); result.Error != nil {
		log.Error().Msg("Failed to create the student!")
		return result.Error
	}

	log.Info().Msg("Create Student Ok!")
	return nil
}

func (s *StudentHandler) GetStudents() ([]Student, error) {

	students := []Student{}

	err := s.DB.Find(&students).Error

	return students, err
}

func (s *StudentHandler) GetStudent(id int) (Student, error) {
	var student Student

	err := s.DB.First(&student, id)
	return student, err.Error
}

func (s *StudentHandler) UpdateStudent(updateStudent Student) error {
	return s.DB.Save(&updateStudent).Error
}
