package utility

import (
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	Db  *gorm.DB
	err error
)

func configGormLogger() logger.Interface {
	// GORM logger configuration
	return logger.New(
		log.New(logFile, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,  // Slow SQL threshold
			LogLevel:                  logger.Error, // Log level
			IgnoreRecordNotFoundError: false,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,         // Disable color
		},
	)
}

// Connect to Database
func DbConnect() {

	//Call getDsn() for Data Source Name
	dsn := getDsn()

	//Config gorm logger
	newLogger := configGormLogger()

	// Db connection
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger, FullSaveAssociations: true})

	if err != nil {
		Log.Panic(err.Error())
		panic(err.Error())
	}

	sql, err := Db.DB()

	if err != nil {
		Log.Panic(err.Error())
		panic(err.Error())
	}

	if err := sql.Ping(); err != nil {
		Log.Panic(err.Error())
		panic(err.Error())
	}
}
