package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	db, err := sql.Open("sqlite3", "./storage_data/mood.db")
	if err != nil {
		log.Fatalf("ошибка подключения к базе: %v", err)
	}
	defer db.Close()

	driver, err := sqlite3.WithInstance(db, &sqlite3.Config{})
	if err != nil {
		log.Fatalf("ошибка инициализации драйвера: %v", err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"sqlite3",
		driver,
	)
	if err != nil {
		log.Fatalf("ошибка создания мигратора: %v", err)
	}

	// m.Force(3)

	// err = m.Up()
	// err = m.Down()
	err = m.Steps(1)

	if err != nil && err != migrate.ErrNoChange {
		log.Fatalf("ошибка применения миграций: %v", err)
	}

	log.Println("✅ Миграции применены успешно")
}
