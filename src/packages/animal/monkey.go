// monkey.go
package animal

type Monkey struct{}

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
