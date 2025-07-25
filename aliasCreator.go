package main

//Testing

import (
	"fmt"
	"strings"
	"database/sql"
	"log"
	_ "github.com/mattn/go-sqlite3"
)
var emails []string = []string{
	"user1@domain1.com",
	"user2@domain1.com",
	"user3@domain1.com",
	"user4@domain1.com"}

func createAlias(email string) string{
	username := strings.SplitAfter(email,"@")
	newAlias := fmt.Sprintf("%sdomain2.com", username[0])
	return newAlias
}

func tableExists() {
	//Make sure the db and table to store users exists.

	db, _ := sql.Open("sqlite3", "./users.db")
	statement, _ := db.Prepare("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, email TEXT, alias TEXT)")

	statement.Exec()
}
func addNewRecord(email string, alias string){
	db, _ := sql.Open("sqlite3", "./users.db")
	statement,_ := db.Prepare("INSERT INTO users(email,alias) VALUES(?,?)")
	statement.Exec(email,alias)
}


func main() {
	//make sure db and table exists
	tableExists()

	db, _ := sql.Open("sqlite3", "./users.db")

	for _, email := range emails{
		var exists bool
		err := db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE email= ?);",email).Scan(&exists)
		if err != nil {
			log.Fatal(err)
			return
		}

		if exists {
			fmt.Printf("REcord with email %s exists\n", email)
		} else {
			alias := createAlias(email)
			fmt.Printf("Creating Alias: %s for user %s\n", alias, email)
			addNewRecord(email,alias)

		}
	}
}
