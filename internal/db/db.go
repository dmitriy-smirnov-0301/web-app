package db

import (
	"database/sql"
	"fmt"
	"ice-creams-app/internal/configs"
	"ice-creams-app/internal/pkg/logger"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type DataBase interface {
	Connect(confFileName string) error
}

type DB struct {
	Connection *sql.DB
}

func New() *DB {
	return &DB{}
}

func (db *DB) Connect(dbName string) error {

	log := logger.GetLogger()

	config := &configs.DBConfig{}

	confFileName := "DB_NAME_" + strings.ToUpper(dbName)

	err := config.LoadConfig(confFileName)
	if err != nil {
		log.Fatalf("Failed to load database configuration: %v", err)
		return err
	}

	connectionURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s&connect_timeout=%s",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
		config.DBSSLMode,
		config.DBConnTimeout,
	)

	/* if config.AppEnv != "local" {
		migrationPath := "file://db/migrations/" + dbName
		if err := migrations.RunMigrations(migrationPath, connectionURL); err != nil {
			log.Fatalf("Failed to run migrations for database \"%s\": %v", dbName, err)
			return err
		}
	} */

	const attemptLimit = 5

	for attempt := 0; attempt < attemptLimit; attempt++ {
		db.Connection, err = sql.Open("postgres", connectionURL)
		if err != nil {
			log.Warnf("Failed to open connection: %v", err)
			log.Warnf("Database is not ready yet, retrying in 2 seconds... (attempt %d/%d)", attempt+1, attemptLimit)
			time.Sleep(2 * time.Second)
			continue
		}
		if pingErr := db.Connection.Ping(); pingErr != nil {
			log.Errorf("Ping failed: %v", pingErr)
			log.Warnf("Database is not ready yet, retrying in 2 seconds... (attempt %d/%d)", attempt+1, attemptLimit)
			time.Sleep(2 * time.Second)
			continue
		}

		log.Infof("Connected to the \"%s\" database successfully. Data source: %s", config.DBName, connectionURL)
		return nil
	}

	log.Fatalf("Failed to connect to the \"%s\" database after %d attempts: %v", config.DBName, attemptLimit, err)
	return err

}
