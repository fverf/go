// animal.go
package animal

// Animal интерфейс
type Animal interface {
    Sound() string
    Move() string
    Eat() string
    Sleep() string
    Name() string
}

// Swimmer интерфейс
type Swimmer interface {
    Swim() string
}
