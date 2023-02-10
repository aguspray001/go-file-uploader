package postgres

import (
	"fmt"

	"github.com/aguspray001/upload-service/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConfig struct {
	DBHost     string
	DBUser     string
	DBPassword string
	DBName     string
	DBPort     string
}

func (conf *PostgresConfig) ConnectToDB() (*gorm.DB, error) {
	conf.DBHost = "10.0.2.15"
	conf.DBUser = "belajardbuser"
	conf.DBPassword = "belajardbpass"
	conf.DBName = "belajardbname"
	conf.DBPort = "5432"

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", conf.DBHost, conf.DBUser, conf.DBPassword, conf.DBName, conf.DBPort)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect the DB server")
	}
	fmt.Println("Success connect to DB!")
	// db.Migrator().DropTable(&models.UploadFile{})
	db.AutoMigrate(&models.UploadFile{})
	return db, nil
}
