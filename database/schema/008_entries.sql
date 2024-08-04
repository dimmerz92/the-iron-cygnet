-- +goose Up
CREATE TABLE IF NOT EXISTS entries (
	id TEXT PRIMARY KEY,
	name TEXT NOT NULL,
	score INTEGER CHECK (score BETWEEN 0 AND 10),
	userId TEXT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	eventId TEXT NOT NULL REFERENCES events(id) ON DELETE CASCADE,
	categoryId INTEGER NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
	createdAt INTEGER NOT NULL DEFAULT (strftime('%s', 'now')),
	shortCode TEXT UNIQUE NOT NULL,
	updatedAt INTEGER
) STRICT;

-- +goose StatementBegin
CREATE TRIGGER
	update_entry
AFTER UPDATE ON
	entries
FOR EACH ROW
BEGIN
	UPDATE
		entries
	SET
		updateAt = strftime('%s', 'now')
	WHERE
		id = NEW.id;
END;
-- +goose StatementEnd

-- +goose Down
DROP TABLE entries;
