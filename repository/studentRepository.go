package repository

import (
	"errors"
	"fmt"
	"rajeevranjan/attendance-system/db"
	"time"

	"github.com/go-pg/pg"
)

//Teacheer Repository represents the repository for teacher- related operations

type StudentRepository struct {
	db *db.DatabaseImpl
}

var (
	ErrAlreadyPunchedInStudent = errors.New("you have already Punched in for the day")
	ErrNoPunchInRecordStudent  = errors.New("no punch-in record found for the user")
)

// NewTeacherRepository creates a new instance of TeacherRepository
func NewStudentRepository(DB *db.DatabaseImpl) *StudentRepository {
	//Return the repository with the established database connection
	return &StudentRepository{
		db: DB,
	}
}

// PunchIn records a teacher's punch in for the day
func (sr *StudentRepository) PunchIn(userId int) error {
	fmt.Println("Punchin repo started")
	// Insert punch-in record
	punch := db.PunchTable{
		UserID:      userId,
		UserType:    db.UserTypeStudent,
		PunchInTime: time.Now(),
	}

	// Perform the insert operation
	_, err := sr.db.DB.Model(&punch).Insert()

	return err
}

func (sr *StudentRepository) IsOpenPunchIn(userid int) (bool, error) {
	var openPunchIn db.PunchTable
	currentDate := time.Now().Format("2006-01-02")

	count, err := sr.db.DB.Model(&openPunchIn).
		Where("user_id =?", userid).
		Where("user_type = ?", db.UserTypeStudent).
		Where("punch_in_time IS NOT NULL AND punch_out_time IS NULL").
		Where("date_trunc('day', punch_in_time) = ?", currentDate).
		Count()

	return count > 0, err

}

func (sr *StudentRepository) IsFirstPunchInToday(userId int) (bool, error) {
	currentDate := time.Now().Format("2006-01-02")
	var sattendence db.SAttendances

	// var count int
	count, err := sr.db.DB.Model(&sattendence).
		Where("student_id = ?", userId).
		Where("date_trunc('day', recorded_at) = ?", currentDate).
		Count()

	return count == 0, err
}

func (sr *StudentRepository) MarkStudentAttendance(userID int) error {
	// currentDate := time.Now().Format("2006-01-02")
	newAttendance := db.SAttendances{
		StudentID:  userID,
		Status:     true,
		RecordedAt: time.Now(),
	}

	_, err := sr.db.DB.Model(&newAttendance). // Ensure no conflict with existing record for the same teacher on the same day
							Insert()

	return err
}

func (sr *StudentRepository) PunchOut(userID int, punchintime time.Time) error {

	fmt.Println("Punchout repo started")

	_, err := sr.db.DB.Model((*db.PunchTable)(nil)).
		Set("punch_out_time = ?", time.Now()).
		Where("user_id = ?", userID).
		Where("user_type =?", db.UserTypeStudent).
		Where("punch_in_time =?", punchintime).
		Update()

	return err

}

func (sr *StudentRepository) LatestPunchIn(userid int) (time.Time, error) {
	var punchInTime time.Time
	currentDate := time.Now().Format("2006-01-02")
	err := sr.db.DB.Model((*db.PunchTable)(nil)).Column("punch_in_time").
		Where("user_id =?", userid).
		Where("user_type =?", db.UserTypeStudent).
		Where("punch_in_time IS NOT NULL AND punch_out_time IS NULL").
		Where("date_trunc('day', punch_in_time) = ?", currentDate).
		Order("punch_in_time DESC").
		Limit(1).Select(&punchInTime)

	if err == pg.ErrNoRows {
		return time.Time{}, err
	}

	return punchInTime, err

}

func (sr *StudentRepository) IsValidStudent(studentID int) (bool, error) {
	// Initialize a teacher object to check if the ID exists
	student := db.Students{ID: studentID}

	// Check if the teacher exists in the database
	exists, err := sr.db.DB.Model(&student).WherePK().Exists()
	if err != nil {
		return false, fmt.Errorf("error checking student existence: %v", err)
	}

	// Return the result
	return exists, nil
}

func (sr *StudentRepository) GetStudentAttendance(studentID int, month time.Month, year int) ([]db.SAttendances, error) {
	// Initialize a slice to store the attendance records
	var attendance []db.SAttendances

	// Query the database for attendance records based on teacherID, month, and year
	err := sr.db.DB.Model(&attendance).
		Where("student_id = ?", studentID).
		Where("EXTRACT(MONTH FROM recorded_at) = ?", month).
		Where("EXTRACT(YEAR FROM recorded_at) = ?", year).
		Select()

	if err != nil {
		return nil, fmt.Errorf("error getting student attendance: %v", err)
	}

	return attendance, nil
}
