package database

var (
	NAW_USER = "INSERT INTO users (logins, password, phone_number, locale, activated, registration_date, code) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id"

	CHECK_USER = "SELECT id FROM users WHERE logins = $1"
)
