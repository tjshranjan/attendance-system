package services

import (
	"errors"
	"fmt"
	"rajeevranjan/attendance-system/db"
	"rajeevranjan/attendance-system/repository"
	"time"

	"github.com/go-pg/pg/v10"
)

var (
	ErrStudentAlreadyExist = errors.New("student already exists in the table with the same ID")

	ErrTeacherAlreadyExist = errors.New("teacher already exists in the table with the same ID")
)

// PrincipalService represents the service layer for Principal-related operations
type PrincipalService struct {
	repo *repository.PrincipalRepository
}

// NewPrincipalService creates a new instance of PrincipalService
func NewPrincipalService(Repo *repository.PrincipalRepository) *PrincipalService {
	return &PrincipalService{
		repo: Repo,
	}
}

func (ps *PrincipalService) AddStudentService(student *db.Students) error {
	// Log the start of the AddStudentservice  method
	fmt.Println("AddStudentService started:", student)

	// Check if the student already exists

	exists, err := ps.repo.IsStudentExists(student.ID)

	if err != nil {
		return err
	}

	if exists {
		fmt.Println("Student already exists:", student.ID)
		return repository.ErrStudentAlreadyExist
	}

	// If student does not exist, add the student
	err = ps.repo.AddStudentRepo(student)
	if err != nil {
		fmt.Println("Error adding student to database:", err)
		return err
	}

	// Log the successful addition of the student
	fmt.Println("Student added successfully:", student.ID)
	return nil
}

func (ps *PrincipalService) AddTeacherService(teacher *db.Teachers) error {
	// Log the start of the AddStudentservice  method
	fmt.Println("AddTeacherService started:", teacher)

	// Check if the student already exists

	exists, err := ps.repo.IsTeacherExists(teacher.ID)

	if err != nil {
		return err
	}

	if exists {
		fmt.Println("teacher already exists:", teacher.ID)
		return repository.ErrStudentAlreadyExist
	}

	// If student does not exist, add the student
	err = ps.repo.AddTeacherRepo(teacher)
	if err != nil {
		fmt.Println("Error adding student to database:", err)
		return err
	}

	// Log the successful addition of the student
	fmt.Println("Teacher added successfully:", teacher.ID)
	return nil
}

func (ps *PrincipalService) GetTeacherAttendanceService(teacherID int, month time.Month, year int) ([]db.TAttendances, error) {
	attendance, err := ps.repo.GetTeacherAttendance(teacherID, month, year)
	if err != nil {
		// Handle the error
		if errors.Is(err, pg.ErrNoRows) {
			// No rows found
			return nil, fmt.Errorf("no attendance records found for teacher ID %d in %s %d", teacherID, month.String(), year)
		}
		// Other error occurred
		return nil, fmt.Errorf("error retrieving teacher attendance: %v", err)
	}
	return attendance, nil
}
