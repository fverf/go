package utils

import (
	"fmt"
	"project/animal"
)

func Describe(a animal.Animal) {
    fmt.Println("животное:", a.Name())
    fmt.Println("звук:", a.Sound())
    fmt.Println("движение:", a.Move())
    fmt.Println("еда:", a.Eat())
    fmt.Println("спит:", a.Sleep())
}
