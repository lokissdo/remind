
package database

import (
	// "bufio"
	// "os"
	// "strings"
	"log"
	"os"
	"time"

	// "database/sql"
	"todo/common/configs"
	"todo/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	// "path/filepath"
	// "io/ioutil"
)

var Db *gorm.DB

func init() {
	initGorm()
}

func initGorm() {
	log.Println("Initializing database...")
	var err error
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      false,       // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
	Db, err = gorm.Open(mysql.Open(configs.Yaml.MySQL.Connection), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		panic(err)
	}

	// Ensure that the database connection is established
	sqlDB, err := Db.DB()
	if err != nil {
		panic(err)
	}

	// Ping the database to verify the connection
	if err := sqlDB.Ping(); err != nil {
		panic(err)
	}
	err = Db.Debug().AutoMigrate(&models.School{}, &models.Class{},
		&models.Parent{}, &models.Student{}, &models.Teacher{}, &models.TeacherTeachClass{},
		&models.Psychologist{}, &models.Calendar{}, &models.Appointment{},
		&models.Test{}, models.TestAssignToClass{}, &models.StudentResult{},
		&models.User{}, &models.UserRole{}, &models.Claims{}, &models.Journal{},
		&models.ChatRoom{}, &models.ChatMessage{}, &models.ConsultationHistory{}, &models.Mail{})
	if err != nil {
		log.Fatal("Error migrating database:", err)
		return
	}

	// err = Db.Debug().AutoMigrate()
	// if err != nil {
	// 	log.Fatal("Error migrating database:", err)
	// 	return
	// }
	// if err := models.InitializeResultStatuses(Db); err != nil {
	// 	panic("Error initializing result statuses: " + err.Error())
	// }

	log.Println("Database initialization complete.")
}
