package storage

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func addMoodToDb(chat_id int, mood string, category string) error {
	query := `
		INSERT INTO mood(mood, chat_id, category)
		values(?, ?, ?)	
`
	_, err := db.Exec(query, mood, chat_id, category)
	if err != nil {
		return err
	}

	return nil
}

func GetAllUsersFromDB() ([]int64, error) {
	query := `
		SELECT DISTINCT(chat_id) FROM mood
	`

	r, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	var chatIDs []int64
	for r.Next() {
		var chatID int64
		if err := r.Scan(&chatID); err != nil {
			return nil, err
		}
		chatIDs = append(chatIDs, chatID)
	}

	if err := r.Err(); err != nil {
		return nil, err
	}

	return chatIDs, nil
}

func InitDb() error {
	if db != nil {
		return nil
	}

	var err error
	db, err = sql.Open("sqlite3", "/app/data/mood.db")
	if err != nil {
		return err
	}

	if err = db.Ping(); err != nil {
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
