-- +goose Up
CREATE TABLE IF NOT EXISTS event_categories (
	eventId TEXT NOT NULL REFERENCES events(id),
	categoryId INTEGER NOT NULL REFERENCES categories(id),
	PRIMARY KEY (eventId, categoryId)
) STRICT;

-- +goose Down
DROP TABLE event_categories;
