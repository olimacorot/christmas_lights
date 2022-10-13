package main_with_interface

import (
	model "practicing/christmas_lights/models"
)

type LightAction interface {
    TakeAction(light model.Light) model.Light
}

type TurnOnBetween struct {
}

func (t TurnOnBetween) TakeAction(light model.Light) model.Light {
    return light.TurnOn()
}

type TurnOffBetween struct {
}

func (t TurnOffBetween) TakeAction(light model.Light) model.Light {
    return light.TurnOff()
}

type ToggleBetween struct {
}

func (t ToggleBetween) TakeAction(light model.Light) model.Light {
    return light.Toggle()
}

func DoActionBetween(grid model.Grid, pointOne, pointTwo model.Point, instruction LightAction) map[model.Point]model.Light {
    for point, light := range grid.Points {
        if point.X >= pointOne.X && point.X <= pointTwo.X  &&
            point.Y >= pointOne.Y && point.Y <= pointTwo.Y {
                grid.Points[point] = instruction.TakeAction(light)
        }
    }

    return grid.Points
}

