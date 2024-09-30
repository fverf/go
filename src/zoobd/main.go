package main

import (
    "fmt"
    "hello-go/db"
)

func main() {
    database := db.ConnectDB()
    defer database.Close()

    db.AddAnimal(database, "Обезьяна", "oo oo aa aa!", "прыгает", "бананы", "спит на дереве") //доб новое животное
    db.AddAnimal(database, "Змея", "pspspspsp", "ползет", "на интервальном голодании", "лежа")
    db.AddAnimal(database, "Тигр", "рррррр!", "бежит", "мясо", "тренирует кунгфу")
    db.AddAnimal(database, "Панда", "мягкое урчание", "катится", "бамбук", "месяцами")

    animals := db.GetAnimals(database)//получаем и выводим
    for _, animal := range animals {
        fmt.Printf("животное: %s\nзвук: %s\nдвижение: %s\nеда: %s\nспит: %s\n\n", 
            animal["name"], animal["sound"], animal["move"], animal["food"], animal["sleep"])
    }

    fmt.Println("отсортированные животные:")
    sortedAnimals := db.GetSortedAnimals(database, "name")
    for _, animal := range sortedAnimals {
        fmt.Println(animal["name"])
    }
}
