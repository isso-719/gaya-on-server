package config

import (
	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
	"os"
)

type appConfig struct {
	HTTPInfo *HTTPInfo
	SQLInfo  *SQLInfo
}

type HTTPInfo struct {
	EndPoint string
	Port     string
}

type SQLInfo struct {
	DBType      string
	SQLUser     string
	SQLPassword string
	SQLAddress  string
	SQLDBName   string
}

func loadDotEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Warn(".env file not loaded")
	}
}

func LoadConfig() *appConfig {
	loadDotEnv()

	endPoint, ok := os.LookupEnv("END_POINT")
	if !ok {
		endPoint = ""
	}
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}
	httpInfo := &HTTPInfo{
		EndPoint: endPoint,
		Port:     port,
	}

	dbType, ok := os.LookupEnv("DB_TYPE")
	if !ok {
		dbType = "mysql"
	}
	sqlUser, ok := os.LookupEnv("MYSQL_USER")
	if !ok {
		sqlUser = "root"
	}
	sqlPassword, ok := os.LookupEnv("MYSQL_PASSWORD")
	if !ok {
		sqlPassword = "password"
	}
	sqlAddr, ok := os.LookupEnv("MYSQL_ADDRESS")
	if !ok {
		sqlAddr = "db:3306"
	}
	sqlDBName, ok := os.LookupEnv("MYSQL_DATABASE")
	if !ok {
		sqlDBName = "g_gayaon"
	}

	dbInfo := &SQLInfo{
		DBType:      dbType,
		SQLUser:     sqlUser,
		SQLPassword: sqlPassword,
		SQLAddress:  sqlAddr,
		SQLDBName:   sqlDBName,
	}

	conf := appConfig{
		SQLInfo:  dbInfo,
		HTTPInfo: httpInfo,
	}

	return &conf
}
