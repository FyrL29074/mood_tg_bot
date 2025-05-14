package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func AddMoodToDb(chat_id int, mood string) error {
	InitDb()

	query := `
		INSERT INTO mood(mood, chat_id)
		values(?, ?)	
`
	_, err := db.Exec(query, mood, chat_id)
	if err != nil {
		return err
	}

	return nil
}

func InitDb() error {
	if db != nil {
		return nil
	}

	var err error
	db, err = sql.Open("sqlite3", "mood.db")
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
		return err
	}

	query := `
		CREATE TABLE IF NOT EXISTS mood (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
			mood TEXT,
			chat_id INTEGER
		);
	`

	_, err = db.Exec(query)
	if err != nil {
		return err
	}

	return nil
}

func CloseDB() error {
	if db != nil {
		return db.Close()
	}
	return nil
}
