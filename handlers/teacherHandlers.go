package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"rajeevranjan/attendance-system/services"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

type TeacherHandler struct {
	Service *services.TeacherService
}

func NewTeacherHandler(service *services.TeacherService) *TeacherHandler {
	return &TeacherHandler{
		Service: service,
	}
}

func (th *TeacherHandler) PunchInHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type header to application/json
	w.Header().Set("Content-Type", "application/json")

	fmt.Println("Just entered in the handler")

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
		UserID string `json:"userId"`
	}
	fmt.Println(json.NewDecoder(r.Body))
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Extract the UserID from the decoded request body
	userID := requestBody.UserID
	fmt.Println(userID)
	// Perform your processing with the userID here
	// For example:
	userId, c := strconv.Atoi(userID)
	fmt.Println(userId, c)
	err := th.Service.PunchInService(userId)
	if err != nil {
		http.Error(w, "Failed to punch in", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (th *TeacherHandler) PunchOutHandler(w http.ResponseWriter, r *http.Request) {
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
		UserID string `json:"userId"`
	}
	if err := json.NewDecoder(r.Body).Decode(&requestBody); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	// Extract the UserID from the decoded request body
	userID := requestBody.UserID

	// Perform your processing with the userID here
	// For example:
	userId, _ := strconv.Atoi(userID)
	err := th.Service.PunchOutService(userId)
	if err != nil {
		http.Error(w, "Failed to punch out", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (th *TeacherHandler) GetTeacherAttendanceHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// // Handle CORS preflight request
	// if r.Method == http.MethodOptions {
	// 	w.Header().Set("Access-Control-Allow-Origin", "*")
	// 	w.Header().Set("Access-Control-Allow-Methods", "GET")
	// 	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	// 	w.WriteHeader(http.StatusOK)
	// 	return
	// }
	// Check if the request method is POST
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//parsing the URL Parameters
	vars := mux.Vars(r)
	teacherID, err := strconv.Atoi(vars["teacherID"])
	if err != nil {
		http.Error(w, "Invalid teacher ID", http.StatusBadRequest)
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

	// Perform your processing with the teacherID, month, and year here
	// For example:
	attendance, err := th.Service.GetTeacherAttendanceService(teacherID, time.Month(month), year)
	if err != nil {
		http.Error(w, "Failed to get teacher attendance", http.StatusInternalServerError)
		return
	}

	// Marshal the response data to JSON
	response, err := json.Marshal(attendance)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func (th *TeacherHandler) GetClassAttendanceHandler(w http.ResponseWriter, r *http.Request) {
	// Set content type header to application/json
	w.Header().Set("Content-Type", "application/json")

	// Handle CORS preflight request
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.WriteHeader(http.StatusOK)
		return
	}
	// Check if the request method is GET
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//parsing the URL Parameters
	vars := mux.Vars(r)

	class, err := strconv.Atoi(vars["class"])
	if err != nil || class > 10 || class < 1 {
		http.Error(w, "Invalid class", http.StatusBadRequest)
		return
	}
	day, err := strconv.Atoi(vars["day"])
	if err != nil || day > 31 || day < 1 {
		http.Error(w, "Invalid Day", http.StatusBadRequest)
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

	// Perform your processing with the teacherID, month, and year here
	// For example:
	classAttendance, err := th.Service.GetClassAttendanceService(class, time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))
	if err != nil {
		http.Error(w, "Failed to get teacher attendance", http.StatusInternalServerError)
		return
	}

	// Marshal the response data to JSON
	response, err := json.Marshal(classAttendance)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
