package bot

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func usersTable() {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS users (username VARCHAR(30) PRIMARY KEY, history VARCHAR(50));")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	stmt.Exec()
}

func InitDB() {
	db, _ = sql.Open("sqlite3", "./dcTrackerW.db")
	usersTable()

}
