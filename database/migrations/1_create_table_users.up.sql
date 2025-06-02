CREATE TABLE users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    username TEXT UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL  DEFAULT current_timestamp,
    updated_at TIMESTAMP NOT NULL  DEFAULT current_timestamp
);

CREATE TRIGGER update_users_updated_at
    AFTER UPDATE ON users
    WHEN old.updated_at <> current_timestamp
BEGIN
    UPDATE users
    SET updated_at = CURRENT_TIMESTAMP
    WHERE id = OLD.id;
END;
