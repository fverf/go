package animal

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
