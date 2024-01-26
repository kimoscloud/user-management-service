package configuration

import (
	"os"
	"strconv"
	"time"
)

type DBConfig struct {
	databaseUser     string
	databasePassword string
	databaseHost     string
	databasePort     int
	databaseName     string
	maxIdleConns     int
	maxOpenConns     int
	connMaxLifetime  int
}

// Getters for DBConfig

func (dbConfig *DBConfig) GetDatabaseUser() string {
	return dbConfig.databaseUser
}

func (dbConfig *DBConfig) GetDatabasePassword() string {
	return dbConfig.databasePassword
}

func (dbConfig *DBConfig) GetDatabaseHost() string {
	return dbConfig.databaseHost
}

func (dbConfig *DBConfig) GetDatabasePort() int {
	return dbConfig.databasePort
}

func (dbConfig *DBConfig) GetDatabaseName() string {
	return dbConfig.databaseName
}
func (dbConfig *DBConfig) GetMaxIdleConnections() int {
	return dbConfig.maxIdleConns
}
func (dbConfig *DBConfig) GetMaxOpenConnections() int {
	return dbConfig.maxOpenConns
}
func (dbConfig *DBConfig) GetConnectionMaxLifetime() time.Duration {
	return time.Duration(dbConfig.connMaxLifetime) * time.Second
}

var dbConfiguration *DBConfig

func GetDBConfig() *DBConfig {
	if dbConfiguration == nil {
		initDBConfig()
	}
	return dbConfiguration
}

func initDBConfig() {
	var err error
	dbConfiguration = &DBConfig{}
	dbConfiguration.databaseUser = os.Getenv("DB_USER")
	dbConfiguration.databasePassword = os.Getenv("DB_PASS")
	dbConfiguration.databaseHost = os.Getenv("DB_HOST")
	dbConfiguration.databasePort, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic("Error parsing the database port")
	}
	dbConfiguration.databaseName = os.Getenv("DB_NAME")
	dbConfiguration.maxIdleConns, err = strconv.Atoi(os.Getenv("MAX_IDLE_CONNS"))
	dbConfiguration.maxOpenConns, err = strconv.Atoi(os.Getenv("MAX_OPEN_CONNS"))
	dbConfiguration.connMaxLifetime, err = strconv.Atoi(os.Getenv("CONN_MAX_LIFETIME"))

}
