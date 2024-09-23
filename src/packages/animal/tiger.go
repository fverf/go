// tiger.go
package animal

type Tiger struct{}

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
