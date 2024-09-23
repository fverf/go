package main

import (
    "fmt"
    "project/animal"   // Импорт пакета animal
    "project/handler"  // Импорт обработчика взаимодействий
    "project/utils"    // Импорт вспомогательных функций
)

func main() {
    animals := []animal.Animal{
        animal.Monkey{},
        animal.Snake{},
        animal.Tiger{},
    }

    for _, a := range animals {
        utils.Describe(a)        // Здесь мы используем Describe
        handler.CheckSwim(a)     // И проверяем на способность плавать
        fmt.Println()
    }
}
