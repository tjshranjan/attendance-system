package services

import (
	"fmt"
	"rajeevranjan/attendance-system/db"
	"rajeevranjan/attendance-system/repository"
	"time"

	"github.com/go-pg/pg/v10"
)

type StudentService struct {
	Repo *repository.StudentRepository
}

// NewTeacherService creates a new instance of TeacherService
func NewStudentService(repo *repository.StudentRepository) *StudentService {
	return &StudentService{
		Repo: repo,
	}
}

// PunchIn handles the logic for a teacher to punch in

func (ss *StudentService) PunchInService(userId int) error {

	isOpenPunchIn, err := ss.Repo.IsOpenPunchIn(userId)
	if err == pg.ErrNoRows {
		fmt.Println("no any open punchin found")
	}

	if err != nil {
		fmt.Println("error in getting the latest punch-in status:", err)
		return err
	}

	if isOpenPunchIn {
		fmt.Println("already have an open punched in:", repository.ErrAlreadyPunchedInStudent)
		return repository.ErrAlreadyPunchedInStudent
	}

	isFirstPunchIn, err := ss.Repo.IsFirstPunchInToday(userId)
	if err != nil && err != pg.ErrNoRows {
		fmt.Println("error in checking if it's the first punch in today:", err)
		return err
	}
	if isFirstPunchIn {
		err := ss.Repo.MarkStudentAttendance(userId)
		if err != nil {
			fmt.Println("error in marking student attendance:", err)
			return err
		}
		fmt.Println("marked student attendance for the first punch in today")
	}

	err = ss.Repo.PunchIn(userId)
	if err != nil {
		// Handle the error
		return fmt.Errorf("error in punchin: %v", err)
	}

	fmt.Println("Punch in successful")
	return nil
}

func (ss *StudentService) PunchOutService(userId int) error {

	latestPunchIn, err := ss.Repo.LatestPunchIn(userId)
	if err == pg.ErrNoRows {
		fmt.Println("no any open punch in found!! Punch In first")
		return err
	}

	if err != nil {
		fmt.Println("error in getting the latest punch-in detail", err)
		return err
	}

	err = ss.Repo.PunchOut(userId, latestPunchIn)
	if err != nil {
		// Handle the error
		return fmt.Errorf("error in punchout: %v", err)
	}

	fmt.Println("Punch Out successful")
	return nil

}

func (ss *StudentService) GetStudentAttendanceService(studentID int, month time.Month, year int) ([]db.SAttendances, error) {
	IsValidStudent, err := ss.Repo.IsValidStudent(studentID)
	if err != nil {
		return nil, fmt.Errorf("could not verify the Student from db: %v", err)
	}

	if !IsValidStudent {
		return nil, fmt.Errorf("not a valid Student ID")
	}

	attendance, err := ss.Repo.GetStudentAttendance(studentID, month, year)
	if err != nil {
		return nil, fmt.Errorf("failed to get student attendance: %v", err)
	}

	return attendance, nil
}
