package handler

import (
    "fmt"
    "project/animal" 
)

func CheckSwim(anyAnimal animal.Animal) {
    if swimmer, ok := anyAnimal.(animal.Swimmer); ok {
        fmt.Println(anyAnimal.Name(), "плавает:", swimmer.Swim())
    } else {
        fmt.Println(anyAnimal.Name(), "не плавает")
    }
}
