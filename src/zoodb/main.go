package main

import (
	"fmt"
	"hello-go/db"
	"log"
)

func main() {
	database := "bde.db"

	conn, err := db.CreateConnection(database) //бд подключается
	if err != nil {
		log.Fatalf("ошибка подключения: %v", err)
	}
	defer db.CloseConnection(conn)

	err = db.CreateTable(conn) //создается таблица если ее нет
	if err != nil {
		log.Fatalf("ошибка создания таблицы: %v", err)
	}

	err = db.AddAnimal(conn, "тигрица", "рычание", "бегает", "мясо", "не спит") // доб данных в таблицу энималс
	if err != nil {
		log.Fatalf("ошибка добавления животного: %v", err)
	}

	err = db.AddAnimal(conn, "змея", "пспспс", "ползает", "мышки", "лежа")
	if err != nil {
		log.Fatalf("ошибка добавления животного: %v", err)
	}

	err = db.AddAnimal(conn, "панда", "живот урчит", "катится", "пельмени", "храпит")
	if err != nil {
		log.Fatalf("ошибка добавления животного: %v", err)
	}

	animals, err := db.GetAllAnimals(conn) //тут получаем всех из энималс
	if err != nil {
		log.Fatalf("ошибка выборки: %v", err)
	}

	// вывод всех животных
	fmt.Println("список животных в бд:")
	for _, animal := range animals {
		fmt.Printf("айди: %d,\nживотное: %s,\nзвук: %s,\nдвигается: %s,\nеда: %s,\nспит: %s\n\n",
			animal.ID, animal.Name, animal.Sound, animal.Move, animal.Food, animal.Sleep)
	}

	// вывод отсортированного списка животных
	sortedAnimals, err := db.GetAnimalsSortedByName(conn) // вывод отсортированных
	if err != nil {
		log.Fatalf("ошибка сортировки данных: %v", err)
	}

	fmt.Println("отсортированный список животных по имени:")
	for _, animal := range sortedAnimals {
		fmt.Printf("айди: %d,\nживотное: %s,\nзвук: %s,\nдвигается: %s,\nеда: %s,\nспит: %s\n\n",
			animal.ID, animal.Name, animal.Sound, animal.Move, animal.Food, animal.Sleep)
	}
}

// вывод в терминале:
/// список животных в бд:
// айди: 1,
// животное: тигрица,
// звук: рычание,
// двигается: бегает,
// еда: мясо,
// спит: не спит

// айди: 2,
// животное: змея,
// звук: пспспс,
// двигается: ползает,
// еда: мышки,
// спит: лежа

// айди: 3,
// животное: панда,
// звук: живот урчит,
// двигается: катится,
// еда: пельмени,
// спит: храпит

// отсортированный список животных по имени:
// айди: 2,
// животное: змея,
// звук: пспспс,
// двигается: ползает,
// еда: мышки,
// спит: лежа

// айди: 3,
// животное: панда,
// звук: живот урчит,
// двигается: катится,
// еда: пельмени,
// спит: храпит

// айди: 1,
// животное: тигрица,
// звук: рычание,
// двигается: бегает,
// еда: мясо,
// спит: не спит
