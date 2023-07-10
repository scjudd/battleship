package main

import (
	"errors"
)

var (
	ShipCarrier    = &Ship{name: "Carrier", length: 5}
	ShipBattleship = &Ship{name: "Battleship", length: 4}
	ShipCruiser    = &Ship{name: "Cruiser", length: 3}
	ShipSubmarine  = &Ship{name: "Submarine", length: 3}
	ShipDestroyer  = &Ship{name: "Destroyer", length: 2}
)

var (
	ErrAlreadyPlaced   = errors.New("ship has already been placed")
	ErrAlreadyShot     = errors.New("a shot has already been fired at the given position")
	ErrInvalidPosition = errors.New("position is out of bounds")
	ErrOverlapping     = errors.New("ship cannot be placed on top of an existing ship")
)

const (
	gridWidth  = 10
	gridHeight = 10
)

type Grid [gridHeight][gridWidth]cell

type cell struct {
	placedShip *placedShip
	shot       bool
}

type placedShip struct {
	ship   *Ship
	health int
}

type Ship struct {
	name   string
	length int
}

type Position struct{ X, Y int }

func (p Position) Offset(n int, o Orientation) Position {
	if o == Horizontal {
		return Position{p.X + n, p.Y}
	}
	return Position{p.X, p.Y + n}
}

type Orientation int

const (
	Horizontal Orientation = iota
	Vertical
)

type ShotResult int

const (
	Missed ShotResult = iota
	Hit
	Sunk
)

func (g *Grid) CheckPlacement(s *Ship, p Position, o Orientation) error {
	if p.X < 0 || p.Y < 0 ||
		o == Horizontal && p.X+s.length >= gridWidth ||
		o == Vertical && p.Y+s.length >= gridHeight {
		return ErrInvalidPosition
	}

	for i := 0; i < s.length; i++ {
		p2 := p.Offset(i, o)
		c := g[p2.Y][p2.X]
		if c.shot {
			return ErrAlreadyShot
		}
		if c.placedShip != nil {
			return ErrOverlapping
		}
	}

	for y := 0; y < gridHeight; y++ {
		for x := 0; x < gridWidth; x++ {
			ps := g[y][x].placedShip
			if ps != nil && ps.ship == s {
				return ErrAlreadyPlaced
			}
		}
	}

	return nil
}

func (g *Grid) PlaceShip(s *Ship, p Position, o Orientation) error {
	if err := g.CheckPlacement(s, p, o); err != nil {
		return err
	}

	ps := &placedShip{
		ship:   s,
		health: s.length,
	}

	for i := 0; i < s.length; i++ {
		p2 := p.Offset(i, o)
		g[p2.Y][p2.X].placedShip = ps
	}

	return nil
}

func (g *Grid) Fire(p Position) (ShotResult, error) {
	c := &g[p.Y][p.X]

	var err error
	if c.shot {
		err = ErrAlreadyShot
	}

	c.shot = true

	if c.placedShip == nil {
		return Missed, err
	}

	if err == nil && c.placedShip.health > 0 {
		c.placedShip.health -= 1
	}

	if c.placedShip.health == 0 {
		return Sunk, err
	}

	return Hit, err
}

func (g *Grid) Defeated() bool {
	for y := 0; y < gridHeight; y++ {
		for x := 0; x < gridWidth; x++ {
			ps := g[y][x].placedShip
			if ps != nil {
				if ps.health > 0 {
					return false
				}
			}
		}
	}
	return true
}

func main() {
	cannedShots(cannedGrids())
}
