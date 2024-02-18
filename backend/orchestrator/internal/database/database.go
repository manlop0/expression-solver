package database

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func ConnectDb() *sql.DB {
	connStr := os.Getenv("DATABASE_URL")
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Couldn't connect to DB: %s", err)
	}

	SetupDatabase(db)

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db
}

func SetupDatabase(db *sql.DB) {
	_, err := db.Exec("CREATE TABLE IF NOT EXISTS expressions (id SERIAL PRIMARY KEY, value TEXT, date TIMESTAMP, status INT2, result TEXT)")
	if err != nil {
		log.Fatalf("Error while creating 'expressions' table: %s", err)
	}

	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS operations (
        name TEXT,
        duration INTEGER
    );

    INSERT INTO operations (name, duration)
    SELECT * FROM (VALUES
        ('+', 10),
        ('-', 10),
        ('*', 10),
        ('/', 10)
    ) AS temp(name, duration)
    WHERE (SELECT COUNT(*) FROM operations) = 0;
`)

	if err != nil {
		log.Fatalf("Error while creating 'operations' table: %s", err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS workers (id INTEGER, working BOOLEAN, workingon TEXT)")
	if err != nil {
		log.Fatalf("Error while creating 'workers' table: %s", err)
	}
}
