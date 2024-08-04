-- name: CreateUser :exec
INSERT INTO
	users (id, firstName, lastName, email, password)
VALUES
	(?, ?, ?, ?, ?);

-- name: CheckEmailExistence :one
SELECT
	1
FROM
	users
WHERE
	email = ?;
