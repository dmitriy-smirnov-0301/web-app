package configs

import (
	"errors"
	"ice-creams-app/internal/pkg/finder"
	"ice-creams-app/internal/pkg/logger"
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	LoadConfig(confFileName string) error
}

type DBConfig struct {
	AppEnv        string `json:"app_env"`
	DBUser        string `json:"db_user"`
	DBPassword    string `json:"db_password"`
	DBHost        string `json:"db_host"`
	DBPort        string `json:"db_port"`
	DBName        string `json:"db_name"`
	DBSSLMode     string `json:"db_sslmode"`
	DBConnTimeout string `json:"db_conntimeout"`
}

type AuthConfig struct {
	Username string `json:"auth_username"`
	Password string `json:"auth_password"`
}

func (db *DBConfig) LoadConfig(confFileName string) error {

	log := logger.GetLogger()

	confPath, err := finder.FindConfigsDir("", ".env")
	if err != nil {
		log.Warnf("Failed to find .env file: %s. Using environment variables", err)
	} else {
		log.Infof("Loading configuration from: %s", confPath)
		if err := godotenv.Load(confPath); err != nil {
			log.Errorf("Error loading configuration, %s", err)
			return err
		}
	}

	db.AppEnv = os.Getenv("APP_ENV")
	db.DBUser = os.Getenv("DB_USER")
	db.DBPassword = os.Getenv("DB_PASSWORD")
	db.DBHost = os.Getenv("DB_HOST")
	db.DBPort = os.Getenv("DB_PORT")
	db.DBName = os.Getenv(confFileName)
	db.DBSSLMode = os.Getenv("DB_SSLMODE")
	db.DBConnTimeout = os.Getenv("DB_CONNTIMEOUT")

	missingVars := make([]string, 0, 6)
	if db.AppEnv == "" {
		missingVars = append(missingVars, "APP_ENV")
	}
	if db.DBUser == "" {
		missingVars = append(missingVars, "DB_USER")
	}
	if db.DBPassword == "" {
		missingVars = append(missingVars, "DB_PASSWORD")
	}
	if db.DBHost == "" {
		missingVars = append(missingVars, "DB_HOST")
	}
	if db.DBPort == "" {
		missingVars = append(missingVars, "DB_PORT")
	}
	if db.DBName == "" {
		missingVars = append(missingVars, confFileName)
	}

	if len(missingVars) > 0 {
		log.Errorf("Missing required environment variables: %v", missingVars)
		return errors.New("missing required environment variables")
	}

	log.Infof("AppEnv: %s", db.AppEnv)
	log.Infof("DBUser: %s", db.DBUser)
	log.Infof("DBPassword: %s", db.DBPassword)
	log.Infof("DBHost: %s", db.DBHost)
	log.Infof("DBPort: %s", db.DBPort)
	log.Infof("DBName: %s", db.DBName)

	log.Infof("Database configuration %s loaded successfully", confFileName)
	return nil

}

func (auth *AuthConfig) LoadConfig(confFileName string) error {

	log := logger.GetLogger()

	confPath, err := finder.FindConfigsDir("", ".env")
	if err != nil {
		log.Warnf("Failed to find .env file: %s. Using environment variables", err)
	} else {
		log.Infof("Loading configuration from: %s", confPath)
		if err := godotenv.Load(confPath); err != nil {
			log.Errorf("Error loading configuration, %s", err)
			return err
		}
	}

	auth.Username = os.Getenv("AUTH_USERNAME")
	auth.Password = os.Getenv("AUTH_PASSWORD")

	missingVars := make([]string, 0, 2)
	if auth.Username == "" {
		missingVars = append(missingVars, "AUTH_USERNAME")
	}
	if auth.Password == "" {
		missingVars = append(missingVars, "AUTH_PASSWORD")
	}

	if len(missingVars) > 0 {
		log.Errorf("Missing required environment variables: %v", missingVars)
		return errors.New("missing required environment variables")
	}

	log.Infof("Username: %s", auth.Username)
	log.Infof("Password: %s", auth.Password)

	log.Infof("Authorization configuration %s loaded successfully", confFileName)
	return nil

}
