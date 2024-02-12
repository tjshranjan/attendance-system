package repository

import (
	"errors"
	"fmt"
	"rajeevranjan/attendance-system/db"
	"time"

	"github.com/go-pg/pg"
)

var ErrStudentAlreadyExist = errors.New("student already exists")
var ErrTeacherAlreadyExist = errors.New("teacher already exists")

// PrincipalRepository represents sthe repository for Principal-related operations
type PrincipalRepository struct {
	db *db.DatabaseImpl
}

// NewPrincipalRepository creates a new instance of PrincipalRepository
func NewPrincipalRepository(DB *db.DatabaseImpl) *PrincipalRepository {

	return &PrincipalRepository{
		db: DB,
	}
}

func (pr *PrincipalRepository) AddStudentRepo(student *db.Students) error {
	// Add logging statement
	fmt.Println("Adding student to the database:", student)

	// Attempt to insert the student into the database
	_, err := pr.db.DB.Model(student).Insert()
	if err != nil {
		fmt.Println("Error adding student to the database:", err)

		// Check if the error is a PostgreSQL error
		if pgErr, ok := err.(pg.Error); ok {
			// Log specific details from pg.Error
			fmt.Println("PostgreSQL has Error:", pgErr.IntegrityViolation())
			// Add more fields as needed
		}
	}

	return err
}

func (pr *PrincipalRepository) IsStudentExists(studentID int) (bool, error) {
	student := &db.Students{ID: studentID}

	err := pr.db.DB.Model(student).Where("id = ?", studentID).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			fmt.Printf("No student found with ID: %d\n", studentID)
			return false, nil
		}
		fmt.Printf("Error retrieving student with ID %d: %v\n", studentID, err)
		return false, err
	}
	fmt.Printf("Student found with ID: %d\n", studentID)
	return true, nil
}

func (pr *PrincipalRepository) AddTeacherRepo(teacher *db.Teachers) error {
	// Add logging statement
	fmt.Println("Adding teacher to the database:", teacher)

	// Attempt to insert the student into the database
	_, err := pr.db.DB.Model(teacher).Insert()
	if err != nil {
		fmt.Println("Error adding teacher to the database:", err)

		// Check if the error is a PostgreSQL error
		if pgErr, ok := err.(pg.Error); ok {
			// Log specific details from pg.Error
			fmt.Println("PostgreSQL has Error:", pgErr.IntegrityViolation())
			// Add more fields as needed
		}
	}

	return err
}

func (pr *PrincipalRepository) IsTeacherExists(teacherID int) (bool, error) {
	teacher := &db.Teachers{ID: teacherID}

	err := pr.db.DB.Model(teacher).Where("id = ?", teacherID).Select()
	if err != nil {
		if err == pg.ErrNoRows {
			fmt.Printf("No teacher found with ID: %d\n", teacherID)
			return false, nil
		}
		fmt.Printf("Error retrieving teacher with ID %d: %v\n", teacherID, err)
		return false, err
	}
	fmt.Printf("Teacher found with ID: %d\n", teacherID)
	return true, nil
}

func (pr *PrincipalRepository) GetTeacherAttendance(teacherID int, month time.Month, year int) ([]db.TAttendances, error) {

	var teacherAttendance []db.TAttendances
	fmt.Println("We are in repo to execute quuery")

	err := pr.db.DB.Model(&teacherAttendance).
		Where("teacher_id = ?", teacherID).
		Where("EXTRACT (MONTH FROM recorded_at) =?", month).
		Where("EXTRACT (YEAR FROM recorded_at) =?", year).
		Select()

	fmt.Println("Query Performed", err)

	return teacherAttendance, err
}
