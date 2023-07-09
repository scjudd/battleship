package main

import "fmt"

func printBoard(s *gameState, p player) {
	for y := 0; y < 10; y++ {
	nextPos:
		for x := 0; x < 10; x++ {
			for _, ship := range s.ships[int(p)] {
				for _, o := range ship.occupies {
					if o.x == x && o.y == y {
						fmt.Print(shipCharacter(ship.name) + " ")
						continue nextPos
					}
				}
			}
			fmt.Print(". ")
		}
		fmt.Println()
	}
}

func shipCharacter(name string) string {
	switch name {
	case "Carrier":
		return "C"
	case "Battleship":
		return "B"
	case "Cruiser":
		return "R"
	case "Submarine":
		return "S"
	case "Destroyer":
		return "D"
	}
	return ""
}

func cannedGameState() *gameState {
	s := &gameState{}

	// player one

	if err := placeCarrier(s, player1, position{2, 2}, vertical); err != nil {
		panic(err)
	}
	if err := placeBattleship(s, player1, position{3, 2}, horizontal); err != nil {
		panic(err)
	}
	if err := placeCruiser(s, player1, position{4, 4}, horizontal); err != nil {
		panic(err)
	}
	if err := placeSubmarine(s, player1, position{0, 1}, vertical); err != nil {
		panic(err)
	}
	if err := placeDestroyer(s, player1, position{8, 7}, vertical); err != nil {
		panic(err)
	}

	// player two

	if err := placeCarrier(s, player2, position{0, 0}, horizontal); err != nil {
		panic(err)
	}
	if err := placeBattleship(s, player2, position{2, 3}, vertical); err != nil {
		panic(err)
	}
	if err := placeCruiser(s, player2, position{2, 2}, horizontal); err != nil {
		panic(err)
	}
	if err := placeSubmarine(s, player2, position{5, 5}, vertical); err != nil {
		panic(err)
	}
	if err := placeDestroyer(s, player2, position{3, 9}, horizontal); err != nil {
		panic(err)
	}

	return s
}
