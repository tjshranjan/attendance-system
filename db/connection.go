package db

import (
	"errors"
	"fmt"
	"os"

	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	"github.com/joho/godotenv"
)

var (
	// once sync.Once // Ensure initialization only happens once
	DB *pg.DB
)

//ConnectionManager represents the connection manager for teh databasae

type DatabaseImpl struct {
	DB *pg.DB
}

//NewConnectionManager creates a new connection manager instance

func NewDatabaseImpl() *DatabaseImpl {
	//Initialize the connection manager

	return &DatabaseImpl{}

}

// Establishing a connection to the PostgreSQL database
func (cm *DatabaseImpl) Init() error {
	// Loading the env variables
	if err := godotenv.Load(); err != nil {
		return fmt.Errorf("error loading the env file: %w", err)
	}

	// Setting up connection options
	opts := &pg.Options{
		Addr:     os.Getenv("ADDR"),
		User:     os.Getenv("USER"),
		Password: os.Getenv("PASSWORD"),
		Database: os.Getenv("DATABASE"),
	}

	// Establishing a connection to the database
	cm.DB = pg.Connect(opts)
	if cm.DB == nil {
		return errors.New("failed to connect to the database")
	}

	fmt.Println("Database connection established")

	// Create schema
	if err := cm.CreateSchema(); err != nil {
		// Close the connection if schema creation fails
		cm.CloseConnection()
		return fmt.Errorf("error creating schema: %v", err)
	}

	return nil
}

// createSchemaa creates the tables
func (cm *DatabaseImpl) CreateSchema() error {
	Students := Students{}

	Teachers := Teachers{}

	SAttendance := SAttendances{}

	TAttendance := TAttendances{}
	PunchTable := PunchTable{}
	// PunchOut := models.PunchOut{}

	models := []interface{}{&Students, &Teachers, &SAttendance, &TAttendance, &PunchTable}

	for _, model := range models {
		err := cm.DB.Model(model).CreateTable(&orm.CreateTableOptions{
			IfNotExists: true,
		})

		if err != nil {
			return err
		}

	}
	return nil
}

// closing the connection to the database
func (cm *DatabaseImpl) CloseConnection() {
	if cm.DB != nil {
		cm.DB.Close()
	}
	fmt.Println("Database connection closed")
}
