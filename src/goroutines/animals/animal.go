package animals

type Animal interface {
    Sound() string
    Move() string
    Eat() string
    Sleep() string
    Name() string
    GetUniqueCharacteristics() map[string]interface{}
}

type Swimmer interface {
    Swim() string
}