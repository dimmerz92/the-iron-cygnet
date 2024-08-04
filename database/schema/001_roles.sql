-- +goose Up
CREATE TABLE IF NOT EXISTS roles (
	id INTEGER PRIMARY KEY,
	name TEXT UNIQUE NOT NULL
) STRICT;

INSERT INTO
	roles (name)
VALUES
	('entrant'),
	('admin'),
	('judge');

-- +goose Down
DROP TABLE roles;
