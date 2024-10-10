package main

import "fmt"

type Animal interface {
	Sound() string
	Move() string
	Eat() string
	Sleep() string
	Name() string
}

type Swimmer interface {
	Swim() string
}

type Monkey struct{}
type Snake struct{}
type Tiger struct{}
type Panda struct{}
type Bogomol struct{}

func (Monkey) Sound() string {
	return "oo oo aa aa!"
}

func (Monkey) Move() string {
	return "прыгает"
}

func (Monkey) Eat() string {
	return "бананы"
}

func (Monkey) Sleep() string {
	return "спит на дереве"
}

func (Monkey) Name() string {
	return "обезьяна"
}

func (Snake) Sound() string {
	return "pspspspsp"
}

func (Snake) Move() string {
	return "ползет"
}

func (Snake) Eat() string {
	return "на интервальном голодании"
}

func (Snake) Sleep() string {
	return "лежа"
}

func (Snake) Name() string {
	return "змея"
}

func (Tiger) Sound() string {
	return "рррррр!"
}

func (Tiger) Move() string {
	return "бежит"
}

func (Tiger) Eat() string {
	return "мясо"
}

func (Tiger) Sleep() string {
	return "тренирует кунгфу"
}

func (Tiger) Name() string {
	return "тигр"
}

func (Tiger) Swim() string {
	return "кошка плывет"
}

func (Panda) Sound() string {
	return "мягкое урчание"
}

func (Panda) Move() string {
	return "катится"
}

func (Panda) Eat() string {
	return "бамбук"
}

func (Panda) Sleep() string {
	return "месяцами"
}

func (Panda) Name() string {
	return "панда"
}

func (Panda) Swim() string {
	return "как бочка"
}

func (Bogomol) Sound() string {
	return "шуршит лапками"
}

func (Bogomol) Move() string {
	return "в припрыжку"
}

func (Bogomol) Eat() string {
	return "насекомые"
}

func (Bogomol) Sleep() string {
	return "неподвижно"
}

func (Bogomol) Name() string {
	return "богомол"
}

func checkswim(anyanimal Animal) {
	if swimmer, good := anyanimal.(Swimmer); good { /////тут проверяется умеет ли животное а плавать
		fmt.Println(anyanimal.Name(), "плавает:", swimmer.Swim())
	} else {
		fmt.Println(anyanimal.Name(), "не плавает")
	}
}

func main() {
	animals := []Animal{Monkey{}, Snake{}, Tiger{}, Panda{}, Bogomol{}}

	for _, a := range animals {
		fmt.Println("животное:", a.Name())
		fmt.Println("звук:", a.Sound())
		fmt.Println("движение:", a.Move()) ///
		fmt.Println("еда:", a.Eat())
		fmt.Println("спит:", a.Sleep())

		checkswim(a)
		fmt.Println()
	}
}
