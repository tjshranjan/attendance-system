package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rajeevranjan/attendance-system/db"
	"rajeevranjan/attendance-system/services"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type PrincipalHandler struct {
	services *services.PrincipalService
}

func NewPrincipalHandler(services *services.PrincipalService) *PrincipalHandler {
	return &PrincipalHandler{
		services: services,
	}
}

func (ph *PrincipalHandler) AddStudentHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// Check request method for CORS preflight
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	// Check content type
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Invalid content type. Expected application/json", http.StatusUnsupportedMediaType)
		return
	}

	var student db.Students

	err := json.NewDecoder(r.Body).Decode(&student)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		fmt.Println("Error decoding request body:", err) // Log error for debugging
		return
	}

	err = ph.services.AddStudentService(&student)
	if err != nil {
		fmt.Println("Error adding student:", err) // Log error for debugging
		if err == services.ErrStudentAlreadyExist {
			http.Error(w, "Student already exists", http.StatusConflict)
			return
		}
		http.Error(w, "Failed to add student", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(student)
}

func (ph *PrincipalHandler) AddTeacherHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// Check request method for CORS preflight
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	// Check content type
	contentType := r.Header.Get("Content-Type")
	if contentType != "application/json" {
		http.Error(w, "Invalid content type. Expected application/json", http.StatusUnsupportedMediaType)
		return
	}

	var teacher db.Teachers

	err := json.NewDecoder(r.Body).Decode(&teacher)

	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		fmt.Println("Error decoding request body:", err) // Log error for debugging
		return
	}

	err = ph.services.AddTeacherService(&teacher)
	if err != nil {
		fmt.Println("Error adding teacher:", err) // Log error for debugging
		if err == services.ErrStudentAlreadyExist {
			http.Error(w, "Teacher already exists", http.StatusConflict)
			return
		}
		http.Error(w, "Failed to add teacher", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(teacher)
}

func (ph *PrincipalHandler) GetTeacherAttendanceHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	teacherID, err := strconv.Atoi(params["teacherID"])

	if err != nil {
		http.Error(w, "Invalaid Teacher Id", http.StatusBadRequest)
		return
	}

	month, err := strconv.Atoi(params["month"])
	if err != nil || month < 1 || month > 12 {
		http.Error(w, "Invalid month", http.StatusBadRequest)
		return

	}

	year, err := strconv.Atoi(params["year"])
	if err != nil || year < 0 {
		http.Error(w, "Invalid year", http.StatusBadRequest)
		return
	}

	teacherAttendance, err := ph.services.GetTeacherAttendanceService(teacherID, time.Month(month), year)
	if err != nil {
		http.Error(w, "Error fetching teacher attendance", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(teacherAttendance)

}
