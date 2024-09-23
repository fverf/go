// snake.go
package animal

type Snake struct{}

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
