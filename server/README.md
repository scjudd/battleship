```console
$ grpcurl -plaintext -proto ../proto/battleship.proto localhost:50051 battleship.Battleship/NewGame                                                         
{
  "gameID": "2b146b20-654a-439d-bf92-1de0b46a00fa",
  "playerID": "f7448ffd-aca9-4b03-a80c-afaaf7056f90"
}

$ grpcurl -plaintext -proto ../proto/battleship.proto -d '{"gameID": "grpcurl -plaintext -proto ../proto/battleship.proto -d '{"gameID": "52917399-f6a6-4efe-871e-b8f313b61d50", "playerID": "e23b11a1-50e3-4659-a5a7-f2a1d57bdad6", "shipName": "Cruiser", "x": 0, "y": 0, "vertical": false}' localhost:50051 battleship.Battleship/PlaceShip
{
  
}"}' localhost:50051 battleship.Battleship/JoinGame

$ grpcurl -plaintext -proto ../proto/battleship.proto -d '{"gameID": "2b146b20-654a-439d-bf92-1de0b46a00fa"}' localhost:50051 battleship.Battleship/JoinGame
{
  "playerID": "9ee088c6-d70e-4f17-96ae-dcae3fd240ee"
}

$ grpcurl -plaintext -proto ../proto/battleship.proto -d '{"gameID": "2b146b20-654a-439d-bf92-1de0b46a00fa", "playerID": "f7448ffd-aca9-4b03-a80c-afaaf7056f90", "shipName": "Carrier", "x": 0, "y": 0, "vertical": false}' localhost:50051 battleship.Battleship/PlaceShip
{
  
}

$ grpcurl -plaintext -proto ../proto/battleship.proto -d '{"gameID": "2b146b20-654a-439d-bf92-1de0b46a00fa", "playerID": "9ee088c6-d70e-4f17-96ae-dcae3fd240ee", "shipName": "Carrier", "x": 0, "y": 0, "vertical": false}' localhost:50051 battleship.Battleship/PlaceShip
{
  
}

$ grpcurl -plaintext -proto ../proto/battleship.proto -d '{"gameID": "2b146b20-654a-439d-bf92-1de0b46a00fa", "playerID": "f7448ffd-aca9-4b03-a80c-afaaf7056f90", "shipName": "Battleship", "x": 0, "y": 1, "vertical": false}' localhost:50051 battleship.Battleship/PlaceShip
{
  
}

$ grpcurl -plaintext -proto ../proto/battleship.proto -d '{"gameID": "2b146b20-654a-439d-bf92-1de0b46a00fa", "playerID": "9ee088c6-d70e-4f17-96ae-dcae3fd240ee", "shipName": "Battleship", "x": 0, "y": 1, "vertical": false}' localhost:50051 battleship.Battleship/PlaceShip
{
  
}

$ grpcurl -plaintext -proto ../proto/battleship.proto -d '{"gameID": "2b146b20-654a-439d-bf92-1de0b46a00fa", "playerID": "f7448ffd-aca9-4b03-a80c-afaaf7056f90", "shipName": "Cruiser", "x": 0, "y": 2, "vertical": false}' localhost:50051 battleship.Battleship/PlaceShip
{
  
}

$ grpcurl -plaintext -proto ../proto/battleship.proto -d '{"gameID": "2b146b20-654a-439d-bf92-1de0b46a00fa", "playerID": "9ee088c6-d70e-4f17-96ae-dcae3fd240ee", "shipName": "Cruiser", "x": 0, "y": 2, "vertical": false}' localhost:50051 battleship.Battleship/PlaceShip   
{
  
}

$ grpcurl -plaintext -proto ../proto/battleship.proto -d '{"gameID": "2b146b20-654a-439d-bf92-1de0b46a00fa", "playerID": "f7448ffd-aca9-4b03-a80c-afaaf7056f90", "shipName": "Submarine", "x": 0, "y": 3, "vertical": false}' localhost:50051 battleship.Battleship/PlaceShip
{
  
}

$ grpcurl -plaintext -proto ../proto/battleship.proto -d '{"gameID": "2b146b20-654a-439d-bf92-1de0b46a00fa", "playerID": "9ee088c6-d70e-4f17-96ae-dcae3fd240ee", "shipName": "Submarine", "x": 0, "y": 3, "vertical": false}' localhost:50051 battleship.Battleship/PlaceShip
{
  
}

$ grpcurl -plaintext -proto ../proto/battleship.proto -d '{"gameID": "2b146b20-654a-439d-bf92-1de0b46a00fa", "playerID": "f7448ffd-aca9-4b03-a80c-afaaf7056f90", "shipName": "Destroyer", "x": 0, "y": 4, "vertical": false}' localhost:50051 battleship.Battleship/PlaceShip
{
  
}

$ grpcurl -plaintext -proto ../proto/battleship.proto -d '{"gameID": "2b146b20-654a-439d-bf92-1de0b46a00fa", "playerID": "9ee088c6-d70e-4f17-96ae-dcae3fd240ee", "shipName": "Destroyer", "x": 0, "y": 4, "vertical": false}' localhost:50051 battleship.Battleship/PlaceShip
{
  
}

$ grpcurl -plaintext -proto ../proto/battleship.proto -d '{"gameID": "2b146b20-654a-439d-bf92-1de0b46a00fa", "playerID": "f7448ffd-aca9-4b03-a80c-afaaf7056f90", "x": 0, "y": 5}' localhost:50051 battleship.Battleship/Fire    
{
  "result": "missed"
}

$ grpcurl -plaintext -proto ../proto/battleship.proto -d '{"gameID": "2b146b20-654a-439d-bf92-1de0b46a00fa", "playerID": "f7448ffd-aca9-4b03-a80c-afaaf7056f90", "x": 0, "y": 5}' localhost:50051 battleship.Battleship/Fire
ERROR:
  Code: Unknown
  Message: it is the other player's turn

$ grpcurl -plaintext -proto ../proto/battleship.proto -d '{"gameID": "2b146b20-654a-439d-bf92-1de0b46a00fa", "playerID": "9ee088c6-d70e-4f17-96ae-dcae3fd240ee", "x": 0, "y": 0}' localhost:50051 battleship.Battleship/Fire    
{
  "result": "hit"
}

$ grpcurl -plaintext -proto ../proto/battleship.proto -d '{"gameID": "2b146b20-654a-439d-bf92-1de0b46a00fa", "playerID": "f7448ffd-aca9-4b03-a80c-afaaf7056f90", "x": 0, "y": 0}' localhost:50051 battleship.Battleship/Fire
{
  "result": "hit"
}

$ grpcurl -plaintext -proto ../proto/battleship.proto -d '{"gameID": "2b146b20-654a-439d-bf92-1de0b46a00fa", "playerID": "9ee088c6-d70e-4f17-96ae-dcae3fd240ee", "x": 1, "y": 0}' localhost:50051 battleship.Battleship/Fire
{
  "result": "hit"
}

$ grpcurl -plaintext -proto ../proto/battleship.proto -d '{"gameID": "2b146b20-654a-439d-bf92-1de0b46a00fa", "playerID": "f7448ffd-aca9-4b03-a80c-afaaf7056f90", "x": 0, "y": 0}' localhost:50051 battleship.Battleship/Fire
ERROR:
  Code: Unknown
  Message: a shot has already been fired at the given position

$ grpcurl -plaintext -proto ../proto/battleship.proto -d '{"gameID": "2b146b20-654a-439d-bf92-1de0b46a00fa", "playerID": "f7448ffd-aca9-4b03-a80c-afaaf7056f90", "x": 1, "y": 0}' localhost:50051 battleship.Battleship/Fire
{
  "result": "hit"
}

$ grpcurl -plaintext -proto ../proto/battleship.proto -d '{"gameID": "2b146b20-654a-439d-bf92-1de0b46a00fa", "playerID": "9ee088c6-d70e-4f17-96ae-dcae3fd240ee", "x": 2, "y": 0}' localhost:50051 battleship.Battleship/Fire
{
  "result": "hit"
}

$ grpcurl -plaintext -proto ../proto/battleship.proto -d '{"gameID": "2b146b20-654a-439d-bf92-1de0b46a00fa", "playerID": "f7448ffd-aca9-4b03-a80c-afaaf7056f90", "x": 2, "y": 0}' localhost:50051 battleship.Battleship/Fire
{
  "result": "hit"
}

$ grpcurl -plaintext -proto ../proto/battleship.proto -d '{"gameID": "2b146b20-654a-439d-bf92-1de0b46a00fa", "playerID": "9ee088c6-d70e-4f17-96ae-dcae3fd240ee", "x": 3, "y": 0}' localhost:50051 battleship.Battleship/Fire
{
  "result": "hit"
}

$ grpcurl -plaintext -proto ../proto/battleship.proto -d '{"gameID": "2b146b20-654a-439d-bf92-1de0b46a00fa", "playerID": "f7448ffd-aca9-4b03-a80c-afaaf7056f90", "x": 3, "y": 0}' localhost:50051 battleship.Battleship/Fire
{
  "result": "hit"
}

$ grpcurl -plaintext -proto ../proto/battleship.proto -d '{"gameID": "2b146b20-654a-439d-bf92-1de0b46a00fa", "playerID": "9ee088c6-d70e-4f17-96ae-dcae3fd240ee", "x": 4, "y": 0}' localhost:50051 battleship.Battleship/Fire
{
  "result": "sunk"
}
```
