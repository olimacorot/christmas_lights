package main

import (
	model "practicing/christmas_lights/models"
)

func doActionBetween(grid model.Grid, pointOne, pointTwo model.Point, instruction string) map[model.Point]model.Light {
    for point, light := range grid.Points {
        if pointsExistinRangeOfGrid(point, pointOne, pointTwo) {
            grid.Points[point] = returnNewValueOfPoint(light, instruction)
        }
    }

    return grid.Points
}

func pointsExistinRangeOfGrid(point, pointOne, pointTwo model.Point) bool {
    return (point.X >= pointOne.X && point.X <= pointTwo.X  && 
        point.Y >= pointOne.Y && point.Y <= pointTwo.Y)
}

func returnNewValueOfPoint(light model.Light, instruction string) model.Light{
    if (instruction == "turnOff") {
        return light.TurnOff()
    }
    
    if (instruction == "toggle") {
        return light.Toggle()
    }

    return light.TurnOn()
}