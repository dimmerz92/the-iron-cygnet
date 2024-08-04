-- name: GetUserByEmail :one
SELECT
	u.*,
	r.name
FROM
	users AS u
	LEFT JOIN roles AS r ON u.role = r.id
WHERE
	u.email = ?;
