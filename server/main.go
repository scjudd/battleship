package main

import (
	"context"
	"log"
	"net"

	"github.com/google/uuid"
	"github.com/scjudd/battleship"
	pb "github.com/scjudd/battleship/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedBattleshipServer
	games map[string]*battleship.Game
}

func (s *server) NewGame(ctx context.Context, in *pb.NewGameRequest) (*pb.NewGameResponse, error) {
	return &pb.NewGameResponse{Id: uuid.New().String()}, nil
}

func (s *server) PlaceShip(ctx context.Context, in *pb.PlaceShipRequest) (*pb.PlaceShipResponse, error) {
	log.Printf("Recieved: %v", in.GetName())
	return &pb.PlaceShipResponse{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterBattleshipServer(s, &server{})

	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
