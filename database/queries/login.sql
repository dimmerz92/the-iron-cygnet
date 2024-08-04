-- name: GetUserByEmail :one
SELECT
	u.*,
	r.name
FROM
	users AS u
	LEFT JOIN roles AS r ON u.role = roles.id
WHERE
	u.email = ?;
