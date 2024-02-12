package services

import (
	"fmt"
	"rajeevranjan/attendance-system/db"
	"rajeevranjan/attendance-system/repository"
	"time"

	"github.com/go-pg/pg/v10"
)

type TeacherService struct {
	Repo *repository.TeacherRepository
}

// NewTeacherService creates a new instance of TeacherService
func NewTeacherService(repo *repository.TeacherRepository) *TeacherService {
	return &TeacherService{
		Repo: repo,
	}
}

// PunchIn handles the logic for a teacher to punch in

func (ts *TeacherService) PunchInService(userId int) error {

	isOpenPunchIn, err := ts.Repo.IsOpenPunchIn(userId)
	if err == pg.ErrNoRows {
		fmt.Println("no any open punchin found")
	}

	if err != nil {
		fmt.Println("error in getting the latest punch-in status:", err)
		return err
	}

	if isOpenPunchIn {
		fmt.Println("already have an open punched in:", repository.ErrAlreadyPunchedIn)
		return repository.ErrAlreadyPunchedIn
	}

	isFirstPunchIn, err := ts.Repo.IsFirstPunchInToday(userId)
	if err != nil && err != pg.ErrNoRows {
		fmt.Println("error in checking if it's the first punch in today:", err)
		return err
	}
	if isFirstPunchIn {
		err := ts.Repo.MarkTeacherAttendance(userId)
		if err != nil {
			fmt.Println("error in marking teacher attendance:", err)
			return err
		}
		fmt.Println("marked teacher attendance for the first punch in today")
	}

	err = ts.Repo.PunchIn(userId)
	if err != nil {
		// Handle the error
		return fmt.Errorf("error in punchin: %v", err)
	}

	fmt.Println("Punch in successful")
	return nil
}

func (ts *TeacherService) PunchOutService(userId int) error {

	latestPunchIn, err := ts.Repo.LatestPunchIn(userId)
	if err == pg.ErrNoRows {
		fmt.Println("no any open punch in found!! Punch In first")
		return err
	}

	if err != nil {
		fmt.Println("error in getting the latest punch-in detail", err)
		return err
	}

	err = ts.Repo.PunchOut(userId, latestPunchIn)
	if err != nil {
		// Handle the error
		return fmt.Errorf("error in punchout: %v", err)
	}

	fmt.Println("Punch Out successful")
	return nil

}

func (ts *TeacherService) GetTeacherAttendanceService(teacherID int, month time.Month, year int) ([]db.TAttendances, error) {
	isValidTeacher, err := ts.Repo.IsValidTeacher(teacherID)
	if err != nil {
		return nil, fmt.Errorf("could not verify the teacher from db: %v", err)
	}

	if !isValidTeacher {
		return nil, fmt.Errorf("not a valid Teacher ID")
	}

	attendance, err := ts.Repo.GetTeacherAttendance(teacherID, month, year)
	if err != nil {
		return nil, fmt.Errorf("failed to get teacher attendance: %v", err)
	}

	return attendance, nil
}

func (ts *TeacherService) GetClassAttendanceService(class int, date time.Time) ([]db.ClassAttendance, error) {
	// Call the repository function to get class attendance
	classAttendance, err := ts.Repo.GetClassAttendance(class, date)
	if err != nil {
		return nil, fmt.Errorf("failed to get class attendance: %v", err)
	}
	return classAttendance, nil
}
