CREATE TABLE checkboxes (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    todolist_id INTEGER,
    checked BOOLEAN NOT NULL DEFAULT false,
    text TEXT,
    created_at TIMESTAMP NOT NULL  DEFAULT current_timestamp,
    updated_at TIMESTAMP NOT NULL  DEFAULT current_timestamp,
    FOREIGN KEY (todolist_id) REFERENCES todolists(id)
);

CREATE TRIGGER update_checkboxes_updated_at
    AFTER UPDATE ON checkboxes
    WHEN old.updated_at <> current_timestamp
BEGIN
    UPDATE checkboxes
    SET updated_at = CURRENT_TIMESTAMP
    WHERE id = OLD.id;
END;
