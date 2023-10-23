package database

import (
	"eth-api/src/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB is a global variable that holds the connection to the database.
var DB *gorm.DB

// Connect establishes a connection to the MySQL database and initializes the DB variable.
func Connect() {
	var err error

	DB, err = gorm.Open(mysql.Open("root:root@tcp(db:3306)/eth?parseTime=true"), &gorm.Config{})

	if err != nil {
		panic("Could not connect with the database!")
	}
}

// AutoMigrate runs the auto-migration for the models, creating the necessary database tables.
func AutoMigrate() {
	err := DB.AutoMigrate(models.EthBlock{}, models.Uncles{}, models.Withdrawals{}, models.EthTransaction{}, models.AccessList{}, models.StorageKey{})
	if err != nil {
		panic("Could not migrate the database!")
	}
}

// GetDB returns the instance of the database connection.
func GetDB() *gorm.DB {
	return DB
}
