package db

import (
	"database/sql"
	"fmt"
	"log"
	_ "github.com/mattn/go-sqlite3"
)

type Animal struct { //хранение структурой информации об этом дурацком зоопарке
	ID    int
	Name  string
	Sound string
	Move  string
	Food  string
	Sleep string
}

func CreateConnection(dbFile string) (*sql.DB, error) { //бд подключается
	conn, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Printf("ошибка подключения к базе данных: %v", err)
		return nil, fmt.Errorf("ошибка подключения к базе данных: %w", err)
	}
	return conn, nil
}

func CreateTable(conn *sql.DB) error { //создается таблица энималс если ее нет
	query := `
	CREATE TABLE IF NOT EXISTS animals (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		sound TEXT,
		move TEXT,
		food TEXT,
		sleep TEXT
	);`
	_, err := conn.Exec(query)
	if err != nil {
		return fmt.Errorf("ошибка создания таблицы: %w", err)
	}
	return nil
}

func AddAnimal(conn *sql.DB, name, sound, move, food, sleep string) error { //добавляет новую запись в энималс
	if name == "" || sound == "" || move == "" || food == "" || sleep == "" {
		return fmt.Errorf("не все поля заполнены")
	}

	query := `INSERT INTO animals (name, sound, move, food, sleep) VALUES (?, ?, ?, ?, ?)`
	_, err := conn.Exec(query, name, sound, move, food, sleep)
	if err != nil {
		return fmt.Errorf("ошибка добавления животного (%s): %w", name, err)
	}
	return nil
}

func GetAllAnimals(conn *sql.DB) ([]Animal, error) { //возвр список животных из энималс
	query := `SELECT * FROM animals`
	rows, err := conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("ошибка выборки: %w", err)
	}
	defer rows.Close()

	var animals []Animal
	for rows.Next() {
		var animal Animal
		err := rows.Scan(&animal.ID, &animal.Name, &animal.Sound, &animal.Move, &animal.Food, &animal.Sleep)
		if err != nil {
			return nil, err
		}
		animals = append(animals, animal)
	}
	return animals, nil
}

func GetAnimalsSortedByName(conn *sql.DB) ([]Animal, error) { //возвр отсортированный список по алфавиту
	query := `SELECT * FROM animals ORDER BY name ASC`
	rows, err := conn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("ошибка выборки данных: %w", err)
	}
	defer rows.Close()

	var animals []Animal
	for rows.Next() {
		var animal Animal
		err := rows.Scan(&animal.ID, &animal.Name, &animal.Sound, &animal.Move, &animal.Food, &animal.Sleep)
		if err != nil {
			return nil, err
		}
		animals = append(animals, animal)
	}
	return animals, nil
}

func CloseConnection(conn *sql.DB) { //закрывашка соединения с бд
	if conn != nil {
		conn.Close()
	}
}
