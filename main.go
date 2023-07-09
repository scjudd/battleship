package main

import (
	"errors"
	"fmt"
)

const (
	boardWidth  = 10
	boardHeight = 10
)

type gameState struct {
	ships [2][]*ship
	shots [2]map[position]shot
}

type position struct{ x, y int }

type ship struct {
	name     string
	occupies []position
	hits     []position
}

type shot int

const (
	unknown shot = iota
	missed
	hit
)

type player int

const (
	player1 player = iota
	player2
)

type rotation int

const (
	horizontal rotation = iota
	vertical
)

var (
	PlacedOutOfBoundsError = errors.New("ship placed outside of board")
	PlacementOverlapError  = errors.New("ship placed on top of an existing ship")
	AlreadyPlacedError     = errors.New("a ship of the same type has already been placed")
)

// Place a new ship on player p's grid. The given position xy represents the
// bow of the ship, and the rotation r determines whether the ship will extend
// down or to the right.
func placeShip(s *gameState, p player, xy position, r rotation, name string, length int) error {
	newShip := &ship{
		name:     name,
		occupies: make([]position, 0, length),
	}

	for i := 0; i < length; i++ {
		var o position
		switch r {
		case horizontal:
			o = position{xy.x + i, xy.y}
		case vertical:
			o = position{xy.x, xy.y + i}
		}
		newShip.occupies = append(newShip.occupies, o)
	}

	for _, o := range newShip.occupies {
		if o.x < 0 || o.x >= boardWidth || o.y < 0 || o.y >= boardHeight {
			return PlacedOutOfBoundsError
		}
	}

	for _, otherShip := range s.ships[int(p)] {
		if newShip.name == otherShip.name {
			return AlreadyPlacedError
		}

		if overlaps(newShip, otherShip) {
			return PlacementOverlapError
		}
	}

	s.ships[int(p)] = append(s.ships[int(p)], newShip)

	return nil
}

func placeCarrier(s *gameState, p player, xy position, r rotation) error {
	return placeShip(s, p, xy, r, "Carrier", 5)
}

func placeBattleship(s *gameState, p player, xy position, r rotation) error {
	return placeShip(s, p, xy, r, "Battleship", 4)
}

func placeCruiser(s *gameState, p player, xy position, r rotation) error {
	return placeShip(s, p, xy, r, "Cruiser", 3)
}

func placeSubmarine(s *gameState, p player, xy position, r rotation) error {
	return placeShip(s, p, xy, r, "Submarine", 3)
}

func placeDestroyer(s *gameState, p player, xy position, r rotation) error {
	return placeShip(s, p, xy, r, "Destroyer", 2)
}

func overlaps(s1, s2 *ship) bool {
	for _, o1 := range s1.occupies {
		for _, o2 := range s2.occupies {
			if o1 == o2 {
				return true
			}
		}
	}
	return false
}

func main() {
	s := cannedGameState()
	fmt.Println("Player one:")
	printBoard(s, player1)
	fmt.Println("Player two:")
	printBoard(s, player2)
}
