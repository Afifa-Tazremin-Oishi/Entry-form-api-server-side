package util

import (
	"fmt"
	"log"
	"os"

	"strconv"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

// Create Connection To Accounts Schema Using Gorm
func CreateConnection() *gorm.DB {
	fmt.Println("Connecting....")
	dbHost := GetConfigString("db_host")
	dbPort := GetConfigInt("db_port")
	dbName := GetConfigString("db_name")
	dbUsername := GetConfigString("db_username")
	dbPassword := GetConfigString("db_password")

	dataSourceName := "host=" + dbHost + " user=" + dbUsername + " password=" + dbPassword + " dbname=" + dbName + " port=" + strconv.Itoa(dbPort) + " sslmode=disable"

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			LogLevel: logger.Info, // Log level
		},
	)
	db, err := gorm.Open(postgres.Open(dataSourceName), &gorm.Config{
		Logger: newLogger,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   GetConfigString("my_schema_name") + ".",
			SingularTable: true,
		},
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		// panic("failed to connect database")
		panic(err)
	} else {
		return db
	}
}

// GetConfigString returns values of string variable from local-config.json
func GetConfigString(key string) string {
	viper.AddConfigPath("./config")
	viper.SetConfigName("local-config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatalf("Invalid type assertion: " + key)
		return ""
	}
	return value
}

// GetConfigInt returns values of int variable from local-config.json
func GetConfigInt(key string) int {
	viper.AddConfigPath("./config")
	viper.SetConfigName("local-config")
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	value := viper.GetInt(key)
	return value
}
