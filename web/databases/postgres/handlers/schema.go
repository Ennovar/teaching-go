package handlers

var schema = []string{
	`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		email text NOT NULL,
		password varchar(60) NOT NULL,
		permissions int8 NOT NULL,
		session_id varchar(16)
	)`,
	`CREATE TABLE IF NOT EXISTS thread (
		id SERIAL PRIMARY KEY,
		author_id int NOT NULL,
		title text NOT NULL,
		body text NOT NULL,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW()
	)`,
	`CREATE TABLE IF NOT EXISTS reply (
		id SERIAL PRIMARY KEY,
		thread_id int NOT NULL,
		replier_id int NOT NULL,
		body text NOT NULL,
		created_at TIMESTAMP DEFAULT NOW()
	)`,
}