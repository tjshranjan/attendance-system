package main

import (
	"fmt"
	"log"
	"net/http"
	"rajeevranjan/attendance-system/db"
	"rajeevranjan/attendance-system/handlers"
	"rajeevranjan/attendance-system/repository"
	"rajeevranjan/attendance-system/routers"
	"rajeevranjan/attendance-system/services"

	"github.com/gorilla/mux"
	// "github.com/yourusername/yourprojectname/db" // Import your db package
)

func main() {
	// Create a new connection manager
	db := db.NewDatabaseImpl()

	// Establish a connection to the database
	if err := db.Init(); err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}
	defer db.CloseConnection() // Defer closing the database connection

	principalRepo := repository.NewPrincipalRepository(db)
	teacherRepo := repository.NewTeacherRepository(db)
	studentRepo := repository.NewStudentRepository(db)

	principalServices := services.NewPrincipalService(principalRepo)
	teacherServices := services.NewTeacherService(teacherRepo)
	studentServices := services.NewStudentService(studentRepo)

	principalHandler := handlers.NewPrincipalHandler(principalServices)
	teacherHandler := handlers.NewTeacherHandler(teacherServices)
	studentHandler := handlers.NewStudentHandler(studentServices)

	principalRouter := routers.NewPrincipalRouter(principalHandler)
	teacherRouter := routers.NewTeacherRouter(teacherHandler)
	studentRouter := routers.NewStudentRouter(studentHandler)

	mainRouter := mux.NewRouter()

	routers := routers.NewRouterImpl(principalRouter, teacherRouter, studentRouter, mainRouter)
	routers.Init()
	// Start the HTTP server
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", mainRouter)
}
