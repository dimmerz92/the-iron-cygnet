-- +goose Up
CREATE TABLE IF NOT EXISTS event_categories (
	eventId TEXT NOT NULL REFERENCES events(id) ON DELETE CASCADE,
	categoryId INTEGER NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
	PRIMARY KEY (eventId, categoryId)
) STRICT;

-- +goose Down
DROP TABLE event_categories;
