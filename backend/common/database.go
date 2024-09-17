package common

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

// dsn := "root:19842100@tcp(127.0.0.1:3307)/gin?charset=utf8mb4&parseTime=True&loc=Local"
func InitDB() *gorm.DB {
	// driverName := "mysql"
	// username := "root"
	// password := "19842100"
	// host := "127.0.0.1"
	// port := "3307"
	// databaseName := "gin"
	// charset := "utf8mb4"
	// parseTime := "True"
	// driverName := viper.GetString("datasourse.driverName")
	username := viper.GetString("datasourse.username")
	password := viper.GetString("datasourse.password")
	host := viper.GetString("datasourse.host")
	port := viper.GetString("datasourse.port")
	databaseName := viper.GetString("datasourse.databaseName")
	charset := viper.GetString("datasourse.charset")
	parseTime := viper.GetString("datasourse.parseTime")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=Local", username, password, host, port, databaseName, charset, parseTime)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}

func GetDB() *gorm.DB {
	return db
}
