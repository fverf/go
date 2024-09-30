package db

import (
    "database/sql"
    _ "github.com/mattn/go-sqlite3"
)

func ConnectDB() *sql.DB {
    db, _ := sql.Open("sqlite3", "./db/gobd.db")

    createTableSQL := `CREATE TABLE IF NOT EXISTS animals (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        name TEXT,
        sound TEXT,
        move TEXT,
        food TEXT,
        sleep TEXT
    );`
    db.Exec(createTableSQL)
    return db
}

func AddAnimal(db *sql.DB, name, sound, move, food, sleep string) { //доб животное в таблицу animals
    query := `INSERT INTO animals (name, sound, move, food, sleep) VALUES (?, ?, ?, ?, ?)`
    db.Exec(query, name, sound, move, food, sleep)
}

func GetAnimals(db *sql.DB) []map[string]string {
    query := `SELECT name, sound, move, food, sleep FROM animals`
    rows, _ := db.Query(query)
    defer rows.Close()

    var animals []map[string]string
    for rows.Next() {
        var name, sound, move, food, sleep string
        rows.Scan(&name, &sound, &move, &food, &sleep)
        animal := map[string]string{
            "name":  name,
            "sound": sound,
            "move":  move,
            "food":  food,
            "sleep": sleep,
        }
        animals = append(animals, animal)
    }
    return animals
}

func GetSortedAnimals(db *sql.DB, sortedBy string) []map[string]string {
    query := "SELECT name, sound, move, food, sleep FROM animals ORDER BY " + sortedBy
    rows, _ := db.Query(query)
    defer rows.Close()

    var animals []map[string]string
    for rows.Next() {
        var name, sound, move, food, sleep string
        rows.Scan(&name, &sound, &move, &food, &sleep)
        animal := map[string]string{
            "name":  name,
            "sound": sound,
            "move":  move,
            "food":  food,
            "sleep": sleep,
        }
        animals = append(animals, animal)
    }
    return animals
}
