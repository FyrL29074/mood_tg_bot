CREATE TABLE IF NOT EXISTS user (
    id INTEGER PRIMARY KEY
);

INSERT INTO user(id) 
SELECT DISTINCT(chat_id) FROM mood;