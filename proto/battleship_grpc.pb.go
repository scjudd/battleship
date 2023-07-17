// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: proto/battleship.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BattleshipClient is the client API for Battleship service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BattleshipClient interface {
	NewGame(ctx context.Context, in *NewGameRequest, opts ...grpc.CallOption) (*NewGameResponse, error)
	JoinGame(ctx context.Context, in *JoinGameRequest, opts ...grpc.CallOption) (*JoinGameResponse, error)
	PlaceShip(ctx context.Context, in *PlaceShipRequest, opts ...grpc.CallOption) (*PlaceShipResponse, error)
}

type battleshipClient struct {
	cc grpc.ClientConnInterface
}

func NewBattleshipClient(cc grpc.ClientConnInterface) BattleshipClient {
	return &battleshipClient{cc}
}

func (c *battleshipClient) NewGame(ctx context.Context, in *NewGameRequest, opts ...grpc.CallOption) (*NewGameResponse, error) {
	out := new(NewGameResponse)
	err := c.cc.Invoke(ctx, "/battleship.Battleship/NewGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *battleshipClient) JoinGame(ctx context.Context, in *JoinGameRequest, opts ...grpc.CallOption) (*JoinGameResponse, error) {
	out := new(JoinGameResponse)
	err := c.cc.Invoke(ctx, "/battleship.Battleship/JoinGame", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *battleshipClient) PlaceShip(ctx context.Context, in *PlaceShipRequest, opts ...grpc.CallOption) (*PlaceShipResponse, error) {
	out := new(PlaceShipResponse)
	err := c.cc.Invoke(ctx, "/battleship.Battleship/PlaceShip", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BattleshipServer is the server API for Battleship service.
// All implementations must embed UnimplementedBattleshipServer
// for forward compatibility
type BattleshipServer interface {
	NewGame(context.Context, *NewGameRequest) (*NewGameResponse, error)
	JoinGame(context.Context, *JoinGameRequest) (*JoinGameResponse, error)
	PlaceShip(context.Context, *PlaceShipRequest) (*PlaceShipResponse, error)
	mustEmbedUnimplementedBattleshipServer()
}

// UnimplementedBattleshipServer must be embedded to have forward compatible implementations.
type UnimplementedBattleshipServer struct {
}

func (UnimplementedBattleshipServer) NewGame(context.Context, *NewGameRequest) (*NewGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NewGame not implemented")
}
func (UnimplementedBattleshipServer) JoinGame(context.Context, *JoinGameRequest) (*JoinGameResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinGame not implemented")
}
func (UnimplementedBattleshipServer) PlaceShip(context.Context, *PlaceShipRequest) (*PlaceShipResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PlaceShip not implemented")
}
func (UnimplementedBattleshipServer) mustEmbedUnimplementedBattleshipServer() {}

// UnsafeBattleshipServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BattleshipServer will
// result in compilation errors.
type UnsafeBattleshipServer interface {
	mustEmbedUnimplementedBattleshipServer()
}

func RegisterBattleshipServer(s grpc.ServiceRegistrar, srv BattleshipServer) {
	s.RegisterService(&Battleship_ServiceDesc, srv)
}

func _Battleship_NewGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(NewGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BattleshipServer).NewGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/battleship.Battleship/NewGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BattleshipServer).NewGame(ctx, req.(*NewGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Battleship_JoinGame_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinGameRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BattleshipServer).JoinGame(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/battleship.Battleship/JoinGame",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BattleshipServer).JoinGame(ctx, req.(*JoinGameRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Battleship_PlaceShip_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PlaceShipRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BattleshipServer).PlaceShip(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/battleship.Battleship/PlaceShip",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BattleshipServer).PlaceShip(ctx, req.(*PlaceShipRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Battleship_ServiceDesc is the grpc.ServiceDesc for Battleship service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Battleship_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "battleship.Battleship",
	HandlerType: (*BattleshipServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "NewGame",
			Handler:    _Battleship_NewGame_Handler,
		},
		{
			MethodName: "JoinGame",
			Handler:    _Battleship_JoinGame_Handler,
		},
		{
			MethodName: "PlaceShip",
			Handler:    _Battleship_PlaceShip_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/battleship.proto",
}