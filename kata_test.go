package main

import (
	model "github.com/olimacorot/christmas_lights/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_turnOnAllLights(t *testing.T) {
    
    turnOnFirstLineX := map[model.Point]model.Light {
        model.Point{0,0}: model.Light{1}, model.Point{0,1}: model.Light{0}, model.Point{0,2}: model.Light{0},
        model.Point{1,0}: model.Light{1}, model.Point{1,1}: model.Light{0}, model.Point{1,2}: model.Light{0},
        model.Point{2,0}: model.Light{1}, model.Point{2,1}: model.Light{0}, model.Point{2,2}: model.Light{0},
    }

    turnOnFirstLineY := map[model.Point]model.Light {
        model.Point{0,0}: model.Light{1}, model.Point{0,1}: model.Light{1}, model.Point{0,2}: model.Light{1},
        model.Point{1,0}: model.Light{0}, model.Point{1,1}: model.Light{0}, model.Point{1,2}: model.Light{0},
        model.Point{2,0}: model.Light{0}, model.Point{2,1}: model.Light{0}, model.Point{2,2}: model.Light{0},
    }

    toggleFirstXY := map[model.Point]model.Light {
        model.Point{0,0}: model.Light{2}, model.Point{0,1}: model.Light{2}, model.Point{0,2}: model.Light{2},
        model.Point{1,0}: model.Light{2}, model.Point{1,1}: model.Light{2}, model.Point{1,2}: model.Light{2},
        model.Point{2,0}: model.Light{2}, model.Point{2,1}: model.Light{2}, model.Point{2,2}: model.Light{2},
    }

    turnOnAll := map[model.Point]model.Light {
        model.Point{0,0}: model.Light{1}, model.Point{0,1}: model.Light{1}, model.Point{0,2}: model.Light{1},
        model.Point{1,0}: model.Light{1}, model.Point{1,1}: model.Light{1}, model.Point{1,2}: model.Light{1},
        model.Point{2,0}: model.Light{1}, model.Point{2,1}: model.Light{1}, model.Point{2,2}: model.Light{1},
    }

    turnOfAll := map[model.Point]model.Light {
        model.Point{0,0}: model.Light{0}, model.Point{0,1}: model.Light{0}, model.Point{0,2}: model.Light{0},
        model.Point{1,0}: model.Light{0}, model.Point{1,1}: model.Light{0}, model.Point{1,2}: model.Light{0},
        model.Point{2,0}: model.Light{0}, model.Point{2,1}: model.Light{0}, model.Point{2,2}: model.Light{0},
    }

    tests := map[string]struct {
        instruction string
        pointOne model.Point
        pointTwo model.Point
        want map[model.Point]model.Light
    }{
        "turn on first line x" : {
            instruction: "turnOn",
            pointOne: model.Point{0, 0}, 
            pointTwo: model.Point{2, 0}, 
            want: turnOnFirstLineX,
        },
        "turn on first line y" : {
            instruction: "turnOn",
            pointOne: model.Point{0, 0}, 
            pointTwo: model.Point{0, 2},
            want: turnOnFirstLineY,
        },
        "toggle first line x and y" : {
            instruction: "toggle",
            pointOne: model.Point{0, 0}, 
            pointTwo: model.Point{2, 2}, 
            want: toggleFirstXY,
        },
        "turn on all" : {
            instruction: "turnOn",
            pointOne: model.Point{0, 0}, 
            pointTwo: model.Point{2, 2}, 
            want: turnOnAll,
        },
        "turn off all" : {
            instruction: "turnOff",
            pointOne: model.Point{0, 0}, 
            pointTwo: model.Point{2, 2}, 
            want: turnOfAll,
        },
    }

    for name, tt := range tests {
        grid := model.NewGrid(2, 2)
        t.Run(name, func(t *testing.T) {
            assert.Equal(t, tt.want, doActionBetween(grid, tt.pointOne, tt.pointTwo, tt.instruction))
        })
    }
}
