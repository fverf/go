package main

import (
    "fmt"
    "project/animal"
    "project/handler"
    "project/utils"
)

func main() {
    animals := []animal.Animal{
        animal.Monkey{},
        animal.Snake{},
        animal.Tiger{},
    }

    for _, a := range animals {
        utils.Describe(a)
        handler.CheckSwim(a)
        fmt.Println()
    }
}
