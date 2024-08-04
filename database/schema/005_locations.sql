-- +goose Up
CREATE TABLE IF NOT EXISTS locations (
	id INTEGER PRIMARY KEY,
	name TEXT UNIQUE,
	street TEXT UNIQUE NOT NULL,
	suburb TEXT NOT NULL,
	state TEXT NOT NULL,
	postcode TEXT NOT NULL CHECK (LENGTH(postcode) = 4 AND postcode BETWEEN '0000' AND '9999'),
	createdAt INTEGER NOT NULL DEFAULT (strftime('%s', 'now')),
	updatedAt INTEGER
) STRICT;

-- +goose StatementBegin
CREATE TRIGGER
	update_location
AFTER UPDATE ON
	locations
FOR EACH ROW
BEGIN
	UPDATE
		locations
	SET
		updatedAt = strftime('%s', 'now')
	WHERE
		id = NEW.id;
END;
-- +goose StatementEnd

-- +goose Down
DROP TABLE locations;
