package main

import "fmt"

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

type gridBuilder struct {
	steps []gridBuilderStep
}

type gridBuilderStep struct {
	s *Ship
	p Position
	o Orientation
}

func (b *gridBuilder) Reset() {
	b.steps = b.steps[:0]
}

func (b *gridBuilder) Place(s *Ship, p Position, o Orientation) {
	step := gridBuilderStep{s, p, o}
	b.steps = append(b.steps, step)
}

func (b *gridBuilder) Build() (*Grid, error) {
	var g Grid

	for _, step := range b.steps {
		if err := g.PlaceShip(step.s, step.p, step.o); err != nil {
			return &g, err
		}
	}

	return &g, nil
}

func cannedGrids() (*Grid, *Grid) {
	// player one

	var b gridBuilder

	b.Place(ShipCarrier, Position{2, 2}, Vertical)
	b.Place(ShipBattleship, Position{3, 2}, Horizontal)
	b.Place(ShipCruiser, Position{4, 4}, Horizontal)
	b.Place(ShipSubmarine, Position{0, 1}, Vertical)
	b.Place(ShipDestroyer, Position{8, 7}, Vertical)

	g1, err := b.Build()
	if err != nil {
		panic(err)
	}

	// player two

	b.Reset()

	b.Place(ShipCarrier, Position{0, 0}, Horizontal)
	b.Place(ShipBattleship, Position{2, 3}, Vertical)
	b.Place(ShipCruiser, Position{2, 2}, Horizontal)
	b.Place(ShipSubmarine, Position{5, 5}, Vertical)
	b.Place(ShipDestroyer, Position{3, 9}, Horizontal)

	g2, err := b.Build()
	if err != nil {
		panic(err)
	}

	return g1, g2
}

func cannedShots(g1, g2 *Grid) {
	g2.Fire(Position{0, 0}) // hit
	g2.Fire(Position{1, 0}) // hit
	g2.Fire(Position{2, 0}) // hit
	g2.Fire(Position{3, 0}) // hit
	g2.Fire(Position{4, 0}) // sunk

	g2.Fire(Position{0, 1}) // miss

	g2.Fire(Position{2, 2}) // hit

	fmt.Println("Player 1 grid")
	printGrid(g1)

	fmt.Println()

	fmt.Println("Player 2 grid")
	printGrid(g2)
}
