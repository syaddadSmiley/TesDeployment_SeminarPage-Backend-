package database

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func Connect() (*sql.DB, error) {
	test := "postgres://mikckilscjohfy:8cb05ddab9540a77d076d2de582adc3ac70bc560af811846fce942dcc8d9544e@ec2-34-199-68-114.compute-1.amazonaws.com:5432/d6b18mu2qcbgu7"
	// db, err := sql.Open("sqlite3", "./database/final_project.db")
	db, err := sql.Open("mysql", test)
	if err != nil {
		return nil, err
	}
	DB = db
	return db, nil
}
