-- name: CreateSession :exec
INSERT INTO
	sessions (id, userId, expiry)
VALUES
	(?, ?, ?);

-- name: GetSession :one
SELECT
	s.id AS sessionId,
	s.expiry,
	u.id AS userId,
	r.name AS role
FROM
	sessions AS s
	LEFT JOIN users AS u ON s.userId = u.id
	LEFT JOIN roles AS r ON u.role = r.id
WHERE
	s.id = ?;

-- name: DeleteSession :exec
DELETE FROM
	sessions
WHERE
	id = ?;

-- name: DeleteExpiredSessions :exec
DELETE FROM
	sessions
WHERE
	expiry <= strftime('%s', 'now');
