package config

import "os"

// DBConfig struct
type DBConfig struct {
	Addr     string
	Host     string
	Port     string
	DBName   string
	UserName string
	Password string
}

// NewDBConfig set DBConfig config and returns it
func NewDBConfig() *DBConfig {

	return &DBConfig{
		Addr:     os.Getenv("MONGO_ADDR"),
		Host:     os.Getenv("MONGO_HOST"),
		Port:     os.Getenv("MONGO_PORT"),
		DBName:   os.Getenv("MONGO_DATABASE"),
		UserName: os.Getenv("MONGO_USER"),
		Password: os.Getenv("MONGO_PASS"),
	}
}
