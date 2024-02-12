package routers

import (
	"rajeevranjan/attendance-system/handlers"

	"github.com/gorilla/mux"
)

type TeacherRouter struct {
	teacherhandler *handlers.TeacherHandler
}

func NewTeacherRouter(teacherhandler *handlers.TeacherHandler) *TeacherRouter {
	return &TeacherRouter{
		teacherhandler: teacherhandler,
	}
}

func (tr *TeacherRouter) SetTeacherRoutes(router *mux.Router) {

	router.HandleFunc("/punchin", tr.teacherhandler.PunchInHandler).Methods("POST")
	router.HandleFunc("/punchout", tr.teacherhandler.PunchOutHandler).Methods("POST")
	router.HandleFunc("/get-attendance/{teacherID}/{month}/{year}", tr.teacherhandler.GetTeacherAttendanceHandler).Methods("GET")
	router.HandleFunc("/get-class-attendance/{class}/{day}/{month}/{year}", tr.teacherhandler.GetClassAttendanceHandler).Methods("GET")
}
