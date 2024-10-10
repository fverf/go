package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "./gobd")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	_, err = db.Exec("DROP TABLE IF EXISTS users;")
	if err != nil {
		log.Fatal(err)
	}

	createUsersTable := `CREATE TABLE IF NOT EXISTS users (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        username TEXT NOT NULL,
        email TEXT NOT NULL UNIQUE
    );`
	_, err = db.Exec(createUsersTable)
	if err != nil {
		log.Fatal(err)
	}

	insertUserSQL := `INSERT INTO users (username, email) VALUES (?, ?)`

	_, err = db.Exec(insertUserSQL, "Катька", "katya@gmail.com")
	if err != nil {
		log.Fatal(err) ///если ошибка
	}

	_, err = db.Exec(insertUserSQL, "Мишка", "mishka@gmail.com")
	if err != nil {
		log.Fatal(err) //ошибка
	}

	//читает данные
	rows, err := db.Query(`SELECT username, email FROM users`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	//тут вывод
	fmt.Println("Пользователи:")
	for rows.Next() {
		var username, email string
		err = rows.Scan(&username, &email)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Имя: %s, Email: %s\n", username, email)
	}
}

//в терминале выводится:
// Пользователи:
// Имя: Катька, Email: katya@gmail.com
// Имя: Мишка, Email: mishka@gmail.com
