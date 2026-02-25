//it holds DB credentials and other configuration settings for the application. It provides a structured way to manage and access these settings throughout the application, ensuring that sensitive information is kept secure and easily configurable.
package config

import "os"

type Config struct {
	DBHost string  `mapstructure:"db_host"` // DBHost is the database host address
	DBPort string   `mapstructure:"db_port"` //port number
	DBUser string   `mapstructure:"db_user"`// DBUser is the database username for authenication
	DBPassword string `mapstructure:"db_password"`//password for authentication
	DBName string `mapstructure:"db_name"`//name of the database connect to
}

func LoadConfig() *Config {
	return &Config{           /*In these Config{}-> creates a new instace of config with default zero values*/
		DBHost: os.Getenv("DB_HOST"),
		DBPort: os.Getenv("DB_PORT"),
		DBUser: os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName: os.Getenv("DB_NAME"),
	}
}