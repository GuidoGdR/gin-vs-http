package database

import (
	"database/sql"

	_ "github.com/lib/pq" // Driver de Postgres
)

func Postgre(dataSourceName string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	if err := createPostgreTable(db); err != nil {
		return nil, err
	}

	return db, nil
}

func createPostgreTable(db *sql.DB) error {

	//CREATE EXTENSION IF NOT EXISTS "pgcrypto";

	const users = `
	CREATE TABLE IF NOT EXISTS users (
        id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
        username VARCHAR(100) NOT NULL UNIQUE CHECK (char_length(username) >= 4),
        password VARCHAR(255) NOT NULL,
        email VARCHAR(255) NOT NULL UNIQUE,
        first_name VARCHAR(100),
        last_name VARCHAR(100),
        date_joined TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
        is_active BOOLEAN NOT NULL DEFAULT TRUE
    );`

	const products = `
	CREATE TABLE IF NOT EXISTS products (
        id INTEGER SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        price REAL NOT NULL,
        stock INTEGER NOT NULL
    );`

	const customers = `
	CREATE TABLE IF NOT EXISTS customers (
		id INTEGER SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		email TEXT UNIQUE NOT NULL
	);`

	const shop_carts = `
	CREATE TABLE IF NOT EXISTS shop_carts (
		id INTEGER SERIAL PRIMARY KEY,
		customer_id INTEGER NOT NULL,
		product_id INTEGER NOT NULL,
		FOREIGN KEY (customer_id) REFERENCES customers(id),
		FOREIGN KEY (product_id) REFERENCES products(id)
	);`

	//query := fmt.Sprintf("%s\n%s\n%s", products, customers, shop_carts)
	const query = users //+ products + customers + shop_carts

	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}
