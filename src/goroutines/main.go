package main

import (
"sync"
"animal-project/animals"
"animal-project/processor"
)

func main() {
    animalList := []animals.Animal{
        animals.Snake{
            Poisonous: true,
            Size: "мини",
            CamouflageSkill:  true,
        },
        animals.Tiger{
            HighSpeed: true,
            Size: "большой",
            Predator: true,
        },
    }

    proc := &processor.AnimalProcessor{}  //создание процессора
    
    var processWG sync.WaitGroup //для обработки животных
    var soundWG sync.WaitGroup //для звуков
    var notifyWG sync.WaitGroup //уведомления

    for _, animal := range animalList { //их обработка
        processWG.Add(1)
        soundWG.Add(1)
        go proc.ProcessAnimal(animal, &processWG, &soundWG) 

        notifyWG.Add(1) //отпр уведомление для каждого
        go proc.SendNotification(animal.Name(), &notifyWG)
    }

    processWG.Wait() //ожидание завершения обработок
    soundWG.Wait()
    notifyWG.Wait()

    println("господи они все обработались")
}