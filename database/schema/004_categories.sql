-- +goose Up
CREATE TABLE IF NOT EXISTS categories (
	id INTEGER PRIMARY KEY,
	name TEXT UNIQUE NOT NULL,
	description TEXT NOT NULL,
	shortCode TEXT UNIQUE NOT NULL,
	createdAt INTEGER NOT NULL DEFAULT (strftime('%s', 'now')),
	updatedAt INTEGER
) STRICT;

-- +goose StatementBegin
CREATE TRIGGER
	update_category
AFTER UPDATE ON
	categories
FOR EACH ROW
BEGIN
	UPDATE
		categories
	SET
		updatedAt = strftime('%s', 'now')
	WHERE
		id = NEW.id;
END;
-- +goose StatementEnd

-- +goose Down
DROP TABLE categories;
