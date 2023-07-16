package main

import (
	"errors"
	"fmt"
	"math/rand"
)

var (
	ShipCarrier    = &Ship{name: "Carrier", length: 5}
	ShipBattleship = &Ship{name: "Battleship", length: 4}
	ShipCruiser    = &Ship{name: "Cruiser", length: 3}
	ShipSubmarine  = &Ship{name: "Submarine", length: 3}
	ShipDestroyer  = &Ship{name: "Destroyer", length: 2}
)

var (
	ErrAllShipsPlaced    = errors.New("all ships have been placed")
	ErrAlreadyPlaced     = errors.New("ship has already been placed")
	ErrAlreadyShot       = errors.New("a shot has already been fired at the given position")
	ErrGameNotOver       = errors.New("the game has not yet finished")
	ErrGameOver          = errors.New("the game has finished")
	ErrInvalidPosition   = errors.New("position is out of bounds")
	ErrNotAllShipsPlaced = errors.New("all ships have not yet been placed")
	ErrOutOfTurn         = errors.New("it is the other player's turn")
	ErrOverlapping       = errors.New("ship cannot be placed on top of an existing ship")
)

func NewGame() *GameSetup {
	return &GameSetup{}
}

type GameSetup [2]playerSetup

func (setup *GameSetup) PlaceShip(player Player, s *Ship, p Position, o Orientation) error {
	ps := &setup[int(player)]

	if ps.grid == nil {
		ps.grid = &Grid{}
	}

	if ps.placedShips == 5 {
		return ErrAllShipsPlaced
	}

	err := ps.grid.PlaceShip(s, p, o)
	if err != nil {
		return err
	}

	ps.placedShips += 1
	return nil
}

func (setup *GameSetup) StartGame() (*RunningGame, error) {
	for i := range setup {
		if setup[i].placedShips != 5 {
			return nil, ErrNotAllShipsPlaced
		}
	}

	return &RunningGame{
		finished: false,
		turn:     PlayerOne,
		grids: [2]*Grid{
			setup[0].grid,
			setup[1].grid,
		},
	}, nil
}

type playerSetup struct {
	grid        *Grid
	placedShips int
}

type RunningGame struct {
	finished bool
	turn     Player
	grids    [2]*Grid
}

func (game *RunningGame) Fire(player Player, p Position) (ShotResult, error) {
	if game.finished {
		return NotFired, ErrGameOver
	}

	if game.turn != player {
		return NotFired, ErrOutOfTurn
	}

	g := game.grids[int(otherPlayer(player))]
	result, err := g.Fire(p)
	if err != nil {
		return result, err
	}

	if result != Won {
		game.turn = otherPlayer(player)
	}

	return result, nil
}

func (game *RunningGame) Winner() (Player, error) {
	if !game.finished {
		return game.turn, ErrGameNotOver
	}
	return game.turn, nil
}

type Player int

const (
	PlayerOne Player = iota
	PlayerTwo
)

func otherPlayer(p Player) Player {
	if p == PlayerOne {
		return PlayerTwo
	}
	return PlayerOne
}

const (
	gridWidth  = 10
	gridHeight = 10
)

type Grid struct {
	cells          [gridHeight][gridWidth]cell
	remainingShips int
}

func (g *Grid) CheckPlacement(s *Ship, p Position, o Orientation) error {
	if p.X < 0 || p.Y < 0 ||
		o == Horizontal && p.X+s.length >= gridWidth ||
		o == Vertical && p.Y+s.length >= gridHeight {
		return ErrInvalidPosition
	}

	if g.remainingShips == 5 {
		return ErrAllShipsPlaced
	}

	for i := 0; i < s.length; i++ {
		p2 := p.offset(i, o)
		c := g.cells[p2.Y][p2.X]
		if c.shot {
			return ErrAlreadyShot
		}
		if c.placedShip != nil {
			return ErrOverlapping
		}
	}

	for y := 0; y < gridHeight; y++ {
		for x := 0; x < gridWidth; x++ {
			ps := g.cells[y][x].placedShip
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

	ps := &placedShip{ship: s, health: s.length}
	for i := 0; i < s.length; i++ {
		p2 := p.offset(i, o)
		g.cells[p2.Y][p2.X].placedShip = ps
	}

	g.remainingShips += 1

	return nil
}

func (g *Grid) Fire(p Position) (ShotResult, error) {
	c := &g.cells[p.Y][p.X]

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
		if err == nil {
			g.remainingShips -= 1
		}
		if g.remainingShips == 0 {
			return Won, err
		}
		return Sunk, err
	}

	return Hit, err
}

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

func randomPosition() Position {
	return Position{
		rand.Intn(gridWidth),
		rand.Intn(gridHeight),
	}
}

func (p Position) offset(n int, o Orientation) Position {
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
	NotFired ShotResult = iota
	Missed
	Hit
	Sunk
	Won
)

type cannedGame struct {
	steps []interface{}
}

type cannedGamePlacement struct {
	player Player
	s      *Ship
	p      Position
	o      Orientation
}

type cannedGameShot struct {
	player Player
	p      Position
}

func (cg *cannedGame) PlaceShip(player Player, s *Ship, p Position, o Orientation) {
	cg.steps = append(cg.steps, cannedGamePlacement{player, s, p, o})
}

func (cg *cannedGame) Fire(player Player, p Position) {
	cg.steps = append(cg.steps, cannedGameShot{player, p})
}

func (cg *cannedGame) Play() {
	setup := NewGame()
	var game *RunningGame
	var err error

	for _, i := range cg.steps {
		switch v := i.(type) {
		case cannedGamePlacement:
			if game != nil {
				panic("game already started")
			}
			err = setup.PlaceShip(v.player, v.s, v.p, v.o)
			if err != nil {
				panic(err)
			}
		case cannedGameShot:
			if game == nil {
				game, err = setup.StartGame()
				if err != nil {
					panic(err)
				}
			}
			_, err = game.Fire(v.player, v.p)
			if err != nil {
				panic(err)
			}
		}
	}

	fmt.Println("PlayerOne grid")
	printGrid(game.grids[int(PlayerOne)])

	fmt.Println()

	fmt.Println("PlayerTwo grid")
	printGrid(game.grids[int(PlayerTwo)])
}

func (cg *cannedGame) PlayRandom() {
	setup := NewGame()
	var game *RunningGame
	var err error

	for _, i := range cg.steps {
		switch v := i.(type) {
		case cannedGamePlacement:
			if game != nil {
				panic("game already started")
			}
			err = setup.PlaceShip(v.player, v.s, v.p, v.o)
			if err != nil {
				panic(err)
			}
		case cannedGameShot:
			if game == nil {
				game, err = setup.StartGame()
				if err != nil {
					panic(err)
				}
			}
			_, err = game.Fire(v.player, v.p)
			if err != nil {
				panic(err)
			}
		}
	}

	if game == nil {
		game, err = setup.StartGame()
		if err != nil {
			panic(err)
		}
	}

	player := PlayerTwo
	for {
		player = otherPlayer(player)
		result, err := game.Fire(player, randomPosition())
		if err != nil {
			continue
		}
		if result == Won {
			break
		}
	}

	fmt.Println("PlayerOne grid")
	printGrid(game.grids[int(PlayerOne)])

	fmt.Println()

	fmt.Println("PlayerTwo grid")
	printGrid(game.grids[int(PlayerTwo)])
}

func printGrid(g *Grid) {
	for y := 0; y < gridHeight; y++ {
		for x := 0; x < gridHeight; x++ {
			c := g.cells[y][x]

			if c.shot {
				if c.placedShip == nil {
					fmt.Print("\x1b[1;5;36m") // miss
				} else if c.placedShip.health > 0 {
					fmt.Print("\x1b[1;5;33m") // hit
				} else {
					fmt.Print("\x1b[1;5;31m") // sunk
				}
			}

			if c.placedShip == nil {
				fmt.Print(". \x1b[0m")
				continue
			}

			switch c.placedShip.ship.name {
			case "Carrier":
				fmt.Print("C ")
			case "Battleship":
				fmt.Print("B ")
			case "Cruiser":
				fmt.Print("R ")
			case "Submarine":
				fmt.Print("S ")
			case "Destroyer":
				fmt.Print("D ")
			}

			fmt.Print("\x1b[0m")
		}

		fmt.Println()
	}
}

func main() {
	cg := &cannedGame{}

	cg.PlaceShip(PlayerOne, ShipCarrier, Position{2, 2}, Vertical)
	cg.PlaceShip(PlayerOne, ShipBattleship, Position{3, 2}, Horizontal)
	cg.PlaceShip(PlayerOne, ShipCruiser, Position{4, 4}, Horizontal)
	cg.PlaceShip(PlayerOne, ShipSubmarine, Position{0, 1}, Vertical)
	cg.PlaceShip(PlayerOne, ShipDestroyer, Position{8, 7}, Vertical)

	cg.PlaceShip(PlayerTwo, ShipCarrier, Position{0, 0}, Horizontal)
	cg.PlaceShip(PlayerTwo, ShipBattleship, Position{2, 3}, Vertical)
	cg.PlaceShip(PlayerTwo, ShipCruiser, Position{2, 2}, Horizontal)
	cg.PlaceShip(PlayerTwo, ShipSubmarine, Position{5, 5}, Vertical)
	cg.PlaceShip(PlayerTwo, ShipDestroyer, Position{3, 9}, Horizontal)

	cg.PlayRandom()
}
