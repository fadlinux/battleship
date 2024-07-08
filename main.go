package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type CoordinatPosition interface {
	GetPosition(positionsXY [][]int)
	HitTarget(targets [][]int) int
	Result()
}

type BattleShipPlayer struct {
	ships    [][]int
	position [][]int
	size     int
}

func (bg *BattleShipPlayer) HitTarget(targets [][]int) int {
	hits := 0
	for _, target := range targets {
		x, y := target[0], target[1]
		if bg.position[x][y] == 'B' {
			bg.position[x][y] = 'X'
			hits++
		} else if bg.position[x][y] == '_' {
			bg.position[x][y] = 'O'
		}
	}
	ß
	return hits
}

func (bg *BattleShipPlayer) GetPosition(positions [][]int) {
	for _, pos := range positions {
		x, y := pos[0], pos[1]
		bg.position[x][y] = 'B'
	}
}

func (bg *BattleShipPlayer) Result() {
	for _, row := range bg.position {
		for _, cell := range row {
			fmt.Print(string(cell), " ")
		}
		fmt.Println()
	}
}

func NewBattleShipPlayer(ships [][]int, size int) *BattleShipPlayer {
	bs := &BattleShipPlayer{
		size:     size,
		position: make([][]int, size),
		ships:    ships,
	}

	for i := range bs.position {
		bs.position[i] = make([]int, size)
		for j := range bs.position[i] {
			bs.position[i][j] = '_'
		}
	}

	bs.GetPosition(ships)

	return bs

}

func toInt(s string) int {
	var i int
	fmt.Sscanf(s, "%d", &i)
	return i
}

func openReadFile(filename string) ([]string, error) {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	lines := strings.Split(string(data), "\n")
	return lines, nil
}

func parsePositions(s string) [][]int {
	pos := strings.Split(s, ":")
	positions := make([][]int, len(pos))
	for i, p := range pos {
		var x, y int
		fmt.Sscanf(p, "%d,%d", &x, &y)
		positions[i] = []int{x, y}
	}
	return positions
}

func parseInput(lines []string) (int, int, [][]int, [][]int, int, [][]int, [][]int) {
	matrixM := toInt(lines[0])
	matrixN := toInt(lines[1])

	shipsP1 := parsePositions(lines[2])
	shipsP2 := parsePositions(lines[3])

	T := toInt(lines[4])

	p1Targets := parsePositions(lines[5])
	p2Targets := parsePositions(lines[6])

	return matrixM, matrixN, shipsP1, shipsP2, T, p1Targets, p2Targets
}

func main() {
	lines, err := openReadFile("sample.txt")
	if err != nil {
		log.Println("Error read input file:", err)
		return
	}

	matrixM, _, shipsP1, shipsP2, _, p1Targets, p2Targets := parseInput(lines)

	positionX := NewBattleShipPlayer(shipsP1, matrixM)
	positionY := NewBattleShipPlayer(shipsP2, matrixM)

}
