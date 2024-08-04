-- +goose Up
CREATE TABLE IF NOT EXISTS users (
	id TEXT PRIMARY KEY,
	firstName TEXT NOT NULL,
	lastName TEXT NOT NULL,
	email TEXT UNIQUE NOT NULL CHECK (email LIKE '_%@_%.__%'),
	role INTEGER NOT NULL DEFAULT (1) REFERENCES roles(id) ON DELETE SET DEFAULT,
	password BLOB NOT NULL,
	createdAt INTEGER NOT NULL DEFAULT (strftime('%s', 'now')),
	updatedAt INTEGER
) STRICT;

-- +goose StatementBegin
CREATE TRIGGER
	update_user
AFTER UPDATE ON
	users
FOR EACH ROW
BEGIN
	UPDATE
		users
	SET
		updatedAt = strftime('%s', 'now')
	WHERE
		id = NEW.id;
END;
-- +goose StatementEnd

-- +goose Down
DROP TABLE users;
