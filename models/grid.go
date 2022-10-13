package model

import (
	"bytes"
	"fmt"
)

type Grid struct {
    lenX int
    lenY int
    Points map[Point]Light
}

// constructor
func NewGrid(lenX, lenY int) (grid Grid) {
    grid = Grid{lenX: lenX, lenY: lenY, Points: make(map[Point]Light)}

    for gridX := 0; gridX <= grid.lenX; gridX++ {
        for gridY := 0; gridY <= grid.lenY; gridY++ {
            grid.Points[Point{gridX, gridY}] = Light{TurnOff}
        }
    }

    return grid
} 

type GridOutputter struct {
    Grid Grid
}

func (grid *Grid) String() string {
    b := new(bytes.Buffer)
    fmt.Fprint(b, "[")
    newLine := 0
    for _, value := range grid.Points {
        if newLine != 0 && newLine % (grid.lenX + 1) == 0 {
            fmt.Fprint(b, "\n")
        }
        fmt.Fprintf(b, "{%v}", value.State)
        newLine++
    }
    fmt.Fprint(b, "]")
    return b.String()
}