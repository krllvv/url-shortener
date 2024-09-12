package config

import "os"

type Config struct {
	DbUser     string
	DbPassword string
	DbHost     string
	DbPort     string
	DbName     string
	Port       string
}

func InitConfig() *Config {
	return &Config{
		DbUser:     os.Getenv("PG_USER"),
		DbPassword: os.Getenv("PG_PASSWORD"),
		DbHost:     os.Getenv("PG_HOST"),
		DbPort:     os.Getenv("PG_PORT"),
		DbName:     os.Getenv("PG_DATABASE"),
		Port:       os.Getenv("PORT"),
	}
}
