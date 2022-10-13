package main

import (
	model "practicing/christmas_lights/models"
)

func DoActionBetween(grid model.Grid, pointOne, pointTwo model.Point, instruction string) map[model.Point]model.Light {
    for point, light := range grid.Points {
        if point.X >= pointOne.X && point.X <= pointTwo.X  &&
            point.Y >= pointOne.Y && point.Y <= pointTwo.Y {

                // hacer refactor
                if (instruction == "turnOn") {
                    grid.Points[point] = light.TurnOn()
                } else if (instruction == "turnOff") {
                    grid.Points[point] = light.TurnOff()
                } else if (instruction == "toggle") {
                    grid.Points[point] = light.Toggle()
                }
        }
    }

    return grid.Points
}

