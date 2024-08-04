-- +goose Up
CREATE TABLE IF NOT EXISTS events (
	id TEXT PRIMARY KEY,
	name TEXT UNIQUE NOT NULL,
	description TEXT NOT NULL,
	locationId INTEGER NOT NULL REFERENCES locations(id) ON DELETE CASCADE,
	visible INTEGER NOT NULL CHECK (visible IN (0, 1)),
	eventDateTime INTEGER NOT NULL,
	shortCode TEXT UNIQUE NOT NULL,
	createdAt INTEGER NOT NULL DEFAULT (strftime('%s', 'now')),
	updatedAt INTEGER
) STRICT;

-- +goose StatementBegin
CREATE TRIGGER
	update_event
AFTER UPDATE ON
	events
FOR EACH ROW
BEGIN
	UPDATE
		events
	SET
		updatedAt = strftime('%s', 'now')
	WHERE
		id = NEW.id;
END;
-- +goose StatementEnd

-- +goose Down
DROP TABLE events;
