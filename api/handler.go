package api

import (
	"errors"
	"fmt"
	"strconv"

	"net/http"

	"github.com/heitorbolisw4/api-students/db"
	"github.com/labstack/echo/v5"
	"gorm.io/gorm"
)

// Handler
func (api *API) getStudents(c *echo.Context) error {
	students, err := api.DB.GetStudents()

	if err != nil {
		return c.String(http.StatusNotFound, "Failed to get students")
	}
	return c.JSON(http.StatusOK, students)
}

func (api *API) createStudent(c *echo.Context) error {
	student := db.Student{}
	if err := c.Bind(&student); err != nil {
		return err
	}
	if err := api.DB.AddStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Error to create student")
	}

	return c.String(http.StatusOK, "Create student")
}

func (api *API) getStudent(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to Get Student ID")
	}
	student, err := api.DB.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student not found!")
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get Student")
	}

	return c.JSON(http.StatusOK, student)
}

func (api *API) updateStudent(c *echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to Get Student ID")
	}

	receivedStudent := db.Student{}
	if err := c.Bind(&receivedStudent); err != nil {
		return err
	}
	updateStudent, err := api.DB.GetStudent(id)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return c.String(http.StatusNotFound, "Student not found!")
	}

	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get Student")
	}
	student := updateStudentInfo(receivedStudent, updateStudent)
	if err := api.DB.UpdateStudent(student); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to save Student")
	}
	return c.JSON(http.StatusOK, student)
}

func (api *API) deleteStudent(c *echo.Context) error {
	id := c.Param("id")
	deleteStud := fmt.Sprintf("Delete %s student", id)
	return c.String(http.StatusOK, deleteStud)
}

func updateStudentInfo(receivedStudent db.Student, student db.Student) db.Student {
	if receivedStudent.Name != "" {

		student.Name = receivedStudent.Name
	}

	if receivedStudent.Cpf != "" {

		student.Cpf = receivedStudent.Cpf
	}

	if receivedStudent.Mail != "" {

		student.Mail = receivedStudent.Mail
	}

	if receivedStudent.Age > 0 {

		student.Age = receivedStudent.Age
	}

	if receivedStudent.IsActive != student.IsActive {

		student.IsActive = receivedStudent.IsActive
	}

	return student
}
