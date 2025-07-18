package storage

import (
	"database/sql"
	"fmt"

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

func addUser(id int) error {
	query := `
		INSERT OR IGNORE INTO user(id) 
		VALUES(?)
	`
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil
}

func GetAllUsersFromDB() ([]int64, error) {
	query := `
		SELECT id FROM user
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

func GetStatistics(chatId int) (*Statistics, error) {
	query := `
		SELECT category, mood, count(id)
        FROM mood 
        WHERE chat_id = ? 
		AND timestamp >= date('now', 'weekday 1', '-7 days')
		AND timestamp < date('now', 'weekday 1')
        GROUP BY category, mood;
	`

	r, err := db.Query(query, chatId)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	stat := Statistics{make(map[string]Category)}
	for r.Next() {
		var categoryName string
		var emotion string
		var count int

		err = r.Scan(&categoryName, &emotion, &count)
		if err != nil {
			return nil, err
		}

		category, exist := stat.Categories[categoryName]
		fmt.Println(category, exist)
		if !exist {
			category = Category{
				Emotions: make(map[string]int),
			}
		}
		category.Emotions[emotion] = count

		stat.Categories[categoryName] = category
	}

	fmt.Println(stat)

	if err = r.Err(); err != nil {
		return nil, err
	}

	return &stat, nil
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
