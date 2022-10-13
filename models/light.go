package model

type Brightness int

const (
    TurnOff Brightness = iota
    TurnOn
    Toggle
)

type Light struct {
    State Brightness
}

func (light *Light) TurnOn() Light {
    light.State = TurnOn
    return *light
}

func (light *Light) TurnOff() Light{
    light.State = TurnOff
    return *light
}

func (light *Light) Toggle() Light{
    light.State = Toggle
    return *light
}
