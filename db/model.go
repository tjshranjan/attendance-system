package db

import (
	"time"
)

// Define constants for UserType
const (
	UserTypeTeacher   = "Teacher"
	UserTypeStudent   = "Student"
	UserTypePrincipal = "Principal"
)

// Student table
type Students struct {
	ID            int    `pg:"id" json:"id"`
	Name          string `pg:"name" json:"name"`
	Class         int    `pg:"class" json:"class"`
	Dob           string `pg:"dob" json:"dob"`
	Email         string `pg:"email" json:"email"`
	ContactNumber int64  `pg:"contactNumber" json:"contactNumber"`
}

// Teacher table
type Teachers struct {
	ID            int    `pg:"id,pk" json:"id"`
	Name          string `pg:"name" json:"name"`
	Dob           string `pg:"dob" json:"dob"`
	Email         string `pg:"email" json:"email"`
	ContactNumber int64  `pg:"contactNumber" json:"contactNumber"`
}

// Student Attendance Table

type SAttendances struct {
	ID         int       `pg:"id,pk" json:"id"`
	StudentID  int       `pg:"student_id" json:"studentId"`
	Status     bool      `pg:"status" json:"status"`
	RecordedAt time.Time `pg:"recorded_at" json:"recordedAt"`
	Student    Students  `pg:"rel:has-one"`
}

// Teacher Attendance Table
type TAttendances struct {
	ID         int       `pg:"id,pk" json:"id"`
	TeacherID  int       `pg:"teacher_id,fk:teachers" json:"teacherId"`
	Status     bool      `pg:"status" json:"status"`
	RecordedAt time.Time `pg:"recorded_at" json:"recordedAt"`
	Teacher    Teachers  `pg:"rel:has-one" json:"-"`
}

// PunchTable
type PunchTable struct {
	ID           int       `pg:"id,pk" json:"id"`
	UserID       int       `pg:"user_id,fk:teachers" json:"userId"`
	UserType     string    `pg:"user_type" json:"userType"`
	PunchInTime  time.Time `pg:"punch-in_time" json:"punchInTime"`
	PunchOutTime time.Time `pg:"punch_out_time" json:"punchOutTime"`
}

type ClassAttendance struct {
	Name   string `json:"name"`
	Date   string `json:"date"`
	Status bool   `json:"status"`
}
