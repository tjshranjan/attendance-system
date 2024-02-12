package routers

import (
	"rajeevranjan/attendance-system/handlers"

	"github.com/gorilla/mux"
)

type PrincipalRouter struct {
	principalhandler *handlers.PrincipalHandler
}

func NewPrincipalRouter(principalhandler *handlers.PrincipalHandler) *PrincipalRouter {
	return &PrincipalRouter{
		principalhandler: principalhandler,
	}
}

func (pr *PrincipalRouter) SetPrincipalRoutes(router *mux.Router) {
	router.HandleFunc("/add-student", pr.principalhandler.AddStudentHandler).Methods("POST")
	router.HandleFunc("/add-teacher", pr.principalhandler.AddTeacherHandler).Methods("POST")
	router.HandleFunc("/get-teacher-attendance/{teacherID}/{month}/{year}", pr.principalhandler.GetTeacherAttendanceHandler).Methods("GET")
}
