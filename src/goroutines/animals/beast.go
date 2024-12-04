package animals

type Snake struct {
    Poisonous bool
    Size string
    CamouflageSkill bool
}

func (s Snake) Sound() string {
    return "pspspspsp"
}

func (s Snake) Move() string {
    return "ползет"
}

func (s Snake) Eat() string {
    return "на интервальном голодании"
}

func (s Snake) Sleep() string {
    return "лежа"
}

func (s Snake) Name() string {
    return "змея"
}

func (s Snake) GetUniqueCharacteristics() map[string]interface{} {
    return map[string]interface{}{
        "ядовитая": s.Poisonous,
        "размер": s.Size,
        "умеет прятаться": s.CamouflageSkill,
    }
}

type Tiger struct {
    HighSpeed bool
    Size string
    Predator bool
}

func (t Tiger) Sound() string {
    return "ррр"
}

func (t Tiger) Move() string {
    return "бежит"
}

func (t Tiger) Eat() string {
    return "мясо"
}

func (t Tiger) Sleep() string {
    return "тренирует кунгфу"
}

func (t Tiger) Name() string {
    return "тигрица"
}

func (t Tiger) Swim() string {
    return "плавает как кошка"
}

func (t Tiger) GetUniqueCharacteristics() map[string]interface{} {
    return map[string]interface{}{
        "очень быстрая": t.HighSpeed,
        "размер": t.Size,
        "хищница": t.Predator,
    }
}