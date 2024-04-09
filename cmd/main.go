package main

import (
	carСatalog "carDirectory"
	"carDirectory/handler"
	"carDirectory/logger"
	"carDirectory/repository"
	"carDirectory/service"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	_ "github.com/lib/pq"
)

var l = logger.Get()

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	l.Info().Msg("Logger is initialized")

	migrateDatabase()

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Username: os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	})

	if err != nil {
		l.Error().Msg(fmt.Sprintf("failed to initialize db: %s", err.Error()))
	}

	repos := repository.NewCarRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(carСatalog.Server)
	if err := srv.Run("8080", handlers.InitRoutes()); err != nil {
		l.Error().Msg(fmt.Sprintf("Car catalog service failed to run: %s", err.Error()))
	}

	l.Info().Msg(fmt.Sprintf("Car catalog service started on port %s", "8080"))
}

// migrateDatabase migrates the database schema using the specified migrations directory and connection string.
//
// No parameters.
// No return values.
func migrateDatabase() {
	l.Info().Msg("Migrating database...")

	migrationsDir := "file://./scheme"
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	m, err := migrate.New(migrationsDir, connectionString)
	if err != nil {
		l.Error().Msg(fmt.Sprintf("Could not initialize migrations: %v", err))
	}

	// Применяем все доступные миграции
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		l.Error().Msg(fmt.Sprintf("Could not apply migrations: %v", err))
	}

	l.Info().Msg("Migrations applied")
}
