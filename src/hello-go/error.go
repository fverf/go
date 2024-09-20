package main
import (
    "fmt"
)

func describe(name string, sound string, move string, eat string, sleep string) {
    fmt.Println("животное:", name)
    fmt.Println("звук:", sound)
    fmt.Println("движение:", move)
    fmt.Println("еда:", eat)
    fmt.Println("спит:", sleep)
}

func main() {
    animals := map[string][]string{
        "змея":   {"pspspsps", "ползет", "голодает", "под камушком"},
        "тигр":      {"рррррр!", "бежит", "мясо", "тренирует кунгфу"},
    }
    var choice string
    fmt.Println("змея или тигр")
    fmt.Scanln(&choice)

    if info, exists := animals[choice]; exists {
        describe(choice, info[0], info[1], info[2], info[3])
    } else {
        fmt.Println("ошибка: неизвестное животное")
    }
}