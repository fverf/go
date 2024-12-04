package processor

import (
    "fmt"
    "sync"
    "time"
    "github.com/gen2brain/beeep"
    "animal-project/animals"
    )

type AnimalProcessor struct{}

func (p *AnimalProcessor) ProcessAnimal(a animals.Animal, wg *sync.WaitGroup, soundWG *sync.WaitGroup) {
    defer wg.Done()
    time.Sleep(time.Second)
    
    go func() { 
        defer soundWG.Done() //горутина для звуков 
        fmt.Printf("звук %s: %s\n", a.Name(), a.Sound())
    }()

    fmt.Printf("животное: %s\n", a.Name())
    fmt.Printf("движение: %s\n", a.Move())
    fmt.Printf("еда: %s\n", a.Eat())
    fmt.Printf("сон: %s\n", a.Sleep())

    if swimmer, ok := a.(animals.Swimmer); ok { //плавает нет
        fmt.Printf("%s плавает: %s\n", a.Name(), swimmer.Swim())
    } else {
        fmt.Printf("%s не плавает\n", a.Name())
    }

    fmt.Println("характеристики:")
    for char, value := range a.GetUniqueCharacteristics() {
        fmt.Printf("- %s: %v\n", char, value)
    }
    fmt.Println()
}

func (p *AnimalProcessor) SendNotification(animal string, wg *sync.WaitGroup) {
    defer wg.Done()
    
    go func() {
        err := beeep.Notify("животное обработано", fmt.Sprintf("%s завершил обработку", animal), "") //тут асинхронное уведомление
        if err != nil {
            fmt.Printf("ошибка уведомления для %s: %v\n", animal, err)
        }
    }()
}