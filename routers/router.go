package routers

import "github.com/gorilla/mux"

type RouterImpl struct {
	principalRouter *PrincipalRouter
	teacherRouter   *TeacherRouter
	studentRouter   *StudentRouter

	router *mux.Router
}

func NewRouterImpl(principalRouter *PrincipalRouter, teacherRouter *TeacherRouter, studentRouter *StudentRouter, router *mux.Router) *RouterImpl {
	return &RouterImpl{
		principalRouter: principalRouter,
		teacherRouter:   teacherRouter,
		studentRouter:   studentRouter,
		router:          router,
	}
}

func (r *RouterImpl) Init() {
	// studentSubRouter := r.router.PathPrefix("/student").Subrouter()
	// r.studentRouter.SetStudentRoutes(studentSubRouter)

	principalSubRouter := r.router.PathPrefix("/principal").Subrouter()
	r.principalRouter.SetPrincipalRoutes(principalSubRouter)

	teacherSubRouter := r.router.PathPrefix("/teacher").Subrouter()
	r.teacherRouter.SetTeacherRoutes(teacherSubRouter)

	studentSubRouter := r.router.PathPrefix("/student").Subrouter()
	r.studentRouter.SetStudentRoutes(studentSubRouter)

}
