package routers

import (
	"rajeevranjan/attendance-system/handlers"

	"github.com/gorilla/mux"
)

type StudentRouter struct {
	studenthandler *handlers.StudentHandler
}

func NewStudentRouter(studenthandler *handlers.StudentHandler) *StudentRouter {
	return &StudentRouter{
		studenthandler: studenthandler,
	}
}

func (sr *StudentRouter) SetStudentRoutes(router *mux.Router) {
	// No need to specify handlers here, as StudentHandler.ServeHTTP will handle it
	router.HandleFunc("/punchin", sr.studenthandler.PunchInHandler).Methods("POST")
	router.HandleFunc("/punchout", sr.studenthandler.PunchOutHandler).Methods("POST")
	router.HandleFunc("/get-attendance/{studentId}/{month}/{year}", sr.studenthandler.GetStudentAttendanceHandler).Methods("GET")

}
