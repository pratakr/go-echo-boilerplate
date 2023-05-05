package database

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/dbresolver"
)

func ConnectMysql() *gorm.DB {

	hostMaster := os.Getenv("DB_HOST_WRITER")
	hostReader := os.Getenv("DB_HOST_READER")
	if hostReader == "" {
		hostReader = hostMaster
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dsnMaster := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		hostMaster,
		port,
		dbName)

	dsnReader := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		hostReader,
		port,
		dbName)

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,          // Disable color
		},
	)

	db, err := gorm.Open(mysql.Open(dsnMaster), &gorm.Config{
		Logger:                 newLogger,
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	if err != nil {
		log.Fatal("Error:", err)
	}

	maxIdle, err := strconv.Atoi(os.Getenv("DB_MAX_IDLE"))
	if err != nil {
		maxIdle = 5
	}

	maxOpen, err := strconv.Atoi(os.Getenv("DB_MAX_OPEN"))
	if err != nil {
		maxOpen = 20
	}

	maxLife, err := strconv.Atoi(os.Getenv("DB_MAX_LIFE"))
	if err != nil {
		maxLife = 60
	}

	dbResolverCfg := dbresolver.Config{
		Sources:  []gorm.Dialector{mysql.Open(dsnMaster)},
		Replicas: []gorm.Dialector{mysql.Open(dsnReader)},
		Policy:   dbresolver.RandomPolicy{}}

	db.Use(
		dbresolver.Register(dbResolverCfg).
			SetConnMaxIdleTime(5 * time.Minute).
			SetConnMaxLifetime(time.Duration(maxLife) * time.Minute).
			SetMaxIdleConns(maxIdle).
			SetMaxOpenConns(maxOpen),
	)

	db.Session(&gorm.Session{PrepareStmt: true})
	return db
}
