package handlers

import (
	"encoding/json"
	"net/http"
	"rajeevranjan/attendance-system/services"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type StudentHandler struct {
	Service *services.StudentService
}

func NewStudentHandler(service *services.StudentService) *StudentHandler {
	return &StudentHandler{
		Service: service,
	}
}

func (sh *StudentHandler) PunchInHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Handle CORS preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Decode the request body into a struct
	var requestBody struct {
		UserID int `json:"userId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Extract the UserID from the decoded request body
	userID := requestBody.UserID

	// Perform your processing with the userID here
	// For example:
	err := sh.Service.PunchInService(userID)
	if err != nil {
		http.Error(w, "Failed to punch in", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (sh *StudentHandler) PunchOutHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Handle CORS preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}
	// Check if the request method is POST
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	// Decode the request body into a struct
	var requestBody struct {
		UserID int `json:"userId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Extract the UserID from the decoded request body
	userID := requestBody.UserID

	// Perform your processing with the userID here
	// For example:
	err := sh.Service.PunchOutService(userID)
	if err != nil {
		http.Error(w, "Failed to punch out", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (th *StudentHandler) GetStudentAttendanceHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Handle CORS preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}
	// Check if the request method is POST
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//parsing the URL Parameters
	vars := mux.Vars(r)
	StudentID, err := strconv.Atoi(vars["studentId"])
	if err != nil {
		http.Error(w, "Invalid Student ID", http.StatusBadRequest)
		return
	}
	month, err := strconv.Atoi(vars["month"])
	if err != nil || month < 1 || month > 12 {
		http.Error(w, "Invalid month", http.StatusBadRequest)
		return
	}
	year, err := strconv.Atoi(vars["year"])
	if err != nil || year < 1900 || year > time.Now().Year() {
		http.Error(w, "Invalid year", http.StatusBadRequest)
		return
	}

	// Perform your processing with the StudentID, month, and year here
	// For example:
	attendance, err := th.Service.GetStudentAttendanceService(StudentID, time.Month(month), year)
	if err != nil {
		http.Error(w, "Failed to get Student attendance", http.StatusInternalServerError)
		return
	}

	// Marshal she response data to JSON
	response, err := json.Marshal(attendance)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)

}
