ALTER TABLE mood RENAME TO mood_old

CREATE TABLE mood (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    timestamp DATETIME DEFAULT CURRENT_TIMESTAMP,
    mood TEXT,
    chat_id INTEGER
);

INSERT INTO mood(id, timestamp, mood, chat_id)
SELECT id, timestamp, mood, chat_id FROM mood_old;

DROP TABLE mood_old;