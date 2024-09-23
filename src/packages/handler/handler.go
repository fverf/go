// handler.go
package handler

import (
    "fmt"
    "project/animal" // Замените "project" на ваше фактическое имя модуля
)

// CheckSwim проверяет, может ли животное плавать
func CheckSwim(anyAnimal animal.Animal) {
    if swimmer, ok := anyAnimal.(animal.Swimmer); ok {
        fmt.Println(anyAnimal.Name(), "плавает:", swimmer.Swim())
    } else {
        fmt.Println(anyAnimal.Name(), "не плавает")
    }
}
