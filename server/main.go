package main

import (
	"context"
	"errors"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/scjudd/battleship"
	pb "github.com/scjudd/battleship/proto"
	"google.golang.org/grpc"
)

var (
	errGameNotFound       = errors.New("game not found")
	errGameNotReady       = errors.New("game not ready")
	errGameAlreadyStarted = errors.New("game has already started")
	errInvalidPlayerID    = errors.New("player ID not associated with game")
	errInvalidShipName    = errors.New("ship name invalid")
)

type server struct {
	pb.UnimplementedBattleshipServer
	games map[string]*game
}

func NewServer() *server {
	return &server{
		games: make(map[string]*game),
	}
}

func (s *server) NewGame(ctx context.Context, in *pb.NewGameRequest) (*pb.NewGameResponse, error) {
	g := &game{
		gameID:      uuid.New().String(),
		playerOneID: uuid.New().String(),
	}
	s.games[g.gameID] = g
	return &pb.NewGameResponse{GameID: g.gameID, PlayerID: g.playerOneID}, nil
}

func (s *server) JoinGame(ctx context.Context, in *pb.JoinGameRequest) (*pb.JoinGameResponse, error) {
	g, ok := s.games[in.GameID]
	if !ok {
		return nil, errGameNotFound
	}
	if g.playerTwoID != "" {
		return nil, errGameAlreadyStarted
	}
	g.playerTwoID = uuid.New().String()
	return &pb.JoinGameResponse{PlayerID: g.playerTwoID}, nil
}

func (s *server) PlaceShip(ctx context.Context, in *pb.PlaceShipRequest) (*pb.PlaceShipResponse, error) {
	g, ok := s.games[in.GameID]
	if !ok {
		return nil, errGameNotFound
	}

	if !g.Ready() {
		return nil, errGameNotReady
	}

	var player battleship.Player
	if in.PlayerID == g.playerOneID {
		player = battleship.PlayerOne
	} else if in.PlayerID == g.playerTwoID {
		player = battleship.PlayerTwo
	} else {
		return nil, errInvalidPlayerID
	}

	var ship *battleship.Ship
	switch in.ShipName {
	case "Carrier":
		ship = battleship.ShipCarrier
	case "Battleship":
		ship = battleship.ShipBattleship
	case "Cruiser":
		ship = battleship.ShipCruiser
	case "Submarine":
		ship = battleship.ShipSubmarine
	case "Destroyer":
		ship = battleship.ShipDestroyer
	default:
		return nil, errInvalidShipName
	}

	p := battleship.Position{X: int(in.X), Y: int(in.Y)}

	var o battleship.Orientation
	if in.Vertical {
		o = battleship.Vertical
	} else {
		o = battleship.Horizontal
	}

	err := g.PlaceShip(player, ship, p, o)
	if err != nil {
		return nil, err
	}

	return &pb.PlaceShipResponse{}, nil
}

func (s *server) Fire(ctx context.Context, in *pb.FireRequest) (*pb.FireResponse, error) {
	g, ok := s.games[in.GameID]
	if !ok {
		return nil, errGameNotFound
	}

	if !g.Ready() {
		return nil, errGameNotReady
	}

	var player battleship.Player
	if in.PlayerID == g.playerOneID {
		player = battleship.PlayerOne
	} else if in.PlayerID == g.playerTwoID {
		player = battleship.PlayerTwo
	} else {
		return nil, errInvalidPlayerID
	}

	p := battleship.Position{X: int(in.X), Y: int(in.Y)}

	result, err := g.Fire(player, p)
	if err != nil {
		return nil, err
	}

	var resultString string
	switch result {
	case battleship.Missed:
		resultString = "missed"
	case battleship.Hit:
		resultString = "hit"
	case battleship.Sunk:
		resultString = "sunk"
	case battleship.Won:
		resultString = "won"
	}

	return &pb.FireResponse{Result: resultString}, nil
}

type game struct {
	battleship.Game
	gameID      string
	playerOneID string
	playerTwoID string
}

func (g *game) Ready() bool {
	return g.gameID != "" && g.playerOneID != "" && g.playerTwoID != ""
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBattleshipServer(s, NewServer())

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
