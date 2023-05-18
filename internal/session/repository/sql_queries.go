package repository

const (
	createSession = `INSERT INTO sessions(refresh_token, expires_at, user_id) VALUES($1, $2, $3) RETURNING *`
	updateSession = `UPDATE sessions SET refresh_token = $1 expires_at = $2`
)
