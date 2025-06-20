CREATE TABLE todolists (
   id INTEGER PRIMARY KEY AUTOINCREMENT,
   user_id INTEGER,
   text TEXT,
   created_at TIMESTAMP NOT NULL  DEFAULT current_timestamp,
   updated_at TIMESTAMP NOT NULL  DEFAULT current_timestamp,
   FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TRIGGER update_todolists_updated_at
    AFTER UPDATE ON todolists
    WHEN old.updated_at <> current_timestamp
BEGIN
    UPDATE todolists
    SET updated_at = CURRENT_TIMESTAMP
    WHERE id = OLD.id;
END;
