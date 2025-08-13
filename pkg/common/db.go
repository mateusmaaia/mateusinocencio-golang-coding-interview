package common

import (
	"errors"
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDBWithConfig() (*gorm.DB, error) {

	connectionError := errors.New("Could not connect to database: invalid config file")

	dbHost := viper.GetString("db.dbhost")
	if len(dbHost) == 0 {
		return nil, connectionError
	}

	dbDatabaseName := viper.GetString("db.dbname")
	if len(dbDatabaseName) == 0 {
		return nil, connectionError
	}

	dbUsername := viper.GetString("db.username")
	if len(dbUsername) == 0 {
		return nil, connectionError
	}

	dbPassword := viper.GetString("db.password")
	if len(dbPassword) == 0 {
		return nil, connectionError
	}

	return ConnectDB(dbUsername, dbPassword, dbHost, dbDatabaseName)
}

func ConnectDB(username, password, host, dbName string) (*gorm.DB, error) {

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", username, password, host, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	tx := db.Exec("CREATE TABLE connection_test (test_col INT)")
	if tx.Error != nil {
		return nil, tx.Error
	}
	db.Exec("DROP TABLE connection_test")

	return db, nil
}
