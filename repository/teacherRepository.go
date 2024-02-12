package repository

import (
	"errors"
	"fmt"
	"rajeevranjan/attendance-system/db"
	"time"

	"github.com/go-pg/pg"
)

//Teacheer Repository represents the repository for teacher- related operations

type TeacherRepository struct {
	db *db.DatabaseImpl
}

var (
	ErrAlreadyPunchedIn = errors.New("you have already Punched in for the day")
	ErrNoPunchInRecord  = errors.New("no punch-in record found for the user")
)

// NewTeacherRepository creates a new instance of TeacherRepository
func NewTeacherRepository(DB *db.DatabaseImpl) *TeacherRepository {
	//Return the repository with the established database connection
	return &TeacherRepository{
		db: DB,
	}
}

// PunchIn records a teacher's punch in for the day
func (tr *TeacherRepository) PunchIn(userId int) error {
	fmt.Println("Punchin repo started")
	// Insert punch-in record
	punch := db.PunchTable{
		UserID:      userId,
		UserType:    db.UserTypeTeacher,
		PunchInTime: time.Now(),
	}

	// Perform the insert operation
	_, err := tr.db.DB.Model(&punch).Insert()

	return err
}

func (tr *TeacherRepository) IsOpenPunchIn(userid int) (bool, error) {
	var openPunchIn db.PunchTable
	currentDate := time.Now().Format("2006-01-02")

	count, err := tr.db.DB.Model(&openPunchIn).
		Where("user_id =?", userid).
		Where("punch_in_time IS NOT NULL AND punch_out_time IS NULL").
		Where("date_trunc('day', punch_in_time) = ?", currentDate).
		Count()

	return count > 0, err

}

func (tr *TeacherRepository) IsFirstPunchInToday(userId int) (bool, error) {
	currentDate := time.Now().Format("2006-01-02")
	var tattendence db.TAttendances

	// var count int
	count, err := tr.db.DB.Model(&tattendence).
		Where("teacher_id = ?", userId).
		Where("date_trunc('day', recorded_at) = ?", currentDate).
		Count()

	return count == 0, err
}

// func (tr *TeacherRepository) MarkTeacherAttendance(userID int) error {
// 	currentDate := time.Now().Format("2006-01-02")
// 	var tattendence db.TAttendances

// 	_, err := tr.db.DB.Model(&tattendence).
// 		Set("status = ?", true).
// 		Where("teacher_id = ?", userID).
// 		Where("DATE_TRUNC('day', recorded_at) = ?", currentDate).
// 		Update()

// 	return err
// }

func (tr *TeacherRepository) MarkTeacherAttendance(userID int) error {
	// currentDate := time.Now().Format("2006-01-02")
	newAttendance := db.TAttendances{
		TeacherID:  userID,
		Status:     true,
		RecordedAt: time.Now(),
	}

	_, err := tr.db.DB.Model(&newAttendance). // Ensure no conflict with existing record for the same teacher on the same day
							Insert()

	return err
}

// PunchOut records a teacher's punch out for the day
// func (repo *TeacherRepository) PunchOut(userId int) error {
// 	punch := new(db.PunchTable)
// 	_, err := repo.db.DB.Model(punch).
// 		Where("user_id=?", userId).
// 		Where("punch_out_time IS NULL AND punch_in_time IS NOT NULL").
// 		Set("punch_out_time = ?", time.Now()).
// 		Update()
// 	return err
// }

func (tr *TeacherRepository) PunchOut(userID int, punchintime time.Time) error {

	fmt.Println("Punchout repo started")
	// Updating the punch-out record
	// punch := db.PunchTable{
	// 	UserID:      userID,
	// 	UserType:    db.UserTypeTeacher,
	// 	PunchOutTime: time.Now(),
	// }

	_, err := tr.db.DB.Model((*db.PunchTable)(nil)).
		Set("punch_out_time = ?", time.Now()).
		Where("user_id = ?", userID).
		Where("punch_in_time =?", punchintime).
		Update()

	return err

}

func (tr *TeacherRepository) LatestPunchIn(userid int) (time.Time, error) {
	var punchInTime time.Time
	currentDate := time.Now().Format("2006-01-02")
	err := tr.db.DB.Model((*db.PunchTable)(nil)).Column("punch_in_time").
		Where("user_id =?", userid).
		Where("user_type =?", db.UserTypeTeacher).
		Where("punch_in_time IS NOT NULL AND punch_out_time IS NULL").
		Where("date_trunc('day', punch_in_time) = ?", currentDate).
		Order("punch_in_time DESC").
		Limit(1).Select(&punchInTime)

	if err == pg.ErrNoRows {
		return time.Time{}, err
	}

	return punchInTime, err

}

func (tr *TeacherRepository) IsValidTeacher(teacherID int) (bool, error) {
	// Initialize a teacher object to check if the ID exists
	teacher := db.Teachers{ID: teacherID}

	// Check if the teacher exists in the database
	exists, err := tr.db.DB.Model(&teacher).WherePK().Exists()
	if err != nil {
		return false, fmt.Errorf("error checking teacher existence: %v", err)
	}

	// Return the result
	return exists, nil
}

func (tr *TeacherRepository) GetTeacherAttendance(teacherID int, month time.Month, year int) ([]db.TAttendances, error) {
	// Initialize a slice to store the attendance records
	var attendance []db.TAttendances

	// Query the database for attendance records based on teacherID, month, and year
	err := tr.db.DB.Model(&attendance).
		Where("teacher_id = ?", teacherID).
		Where("EXTRACT(MONTH FROM recorded_at) = ?", month).
		Where("EXTRACT(YEAR FROM recorded_at) = ?", year).
		Select()

	if err != nil {
		return nil, fmt.Errorf("error getting teacher attendance: %v", err)
	}

	return attendance, nil
}

func (tr *TeacherRepository) GetClassAttendance(class int, date time.Time) ([]db.ClassAttendance, error) {
	// Retrieve the class attendance records from the database

	var attendances []db.SAttendances
	err := tr.db.DB.Model(&attendances).Relation("Students").Where("students.class =?", class).
		Where("Date(recorded_at)=?", date.Format("2006-01-02")).Select()

	if err != nil {
		return nil, fmt.Errorf("failed to get class attendance: %v", err)
	}

	// Map the database records to db.ClassAttendance struct instances
	var classAttendances []db.ClassAttendance
	for _, attendance := range attendances {
		classAttendance := db.ClassAttendance{
			Name:   attendance.Student.Name,
			Date:   attendance.RecordedAt.Format("2006-01-02"), // Format date as "YYYY-MM-DD"
			Status: attendance.Status,
		}
		classAttendances = append(classAttendances, classAttendance)
	}

	return classAttendances, nil
}
