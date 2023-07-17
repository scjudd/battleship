```console
$ grpcurl -plaintext -proto ../proto/battleship.proto localhost:50051 battleship.Battleship/NewGame
{
  "gameID": "52917399-f6a6-4efe-871e-b8f313b61d50",
  "playerID": "e23b11a1-50e3-4659-a5a7-f2a1d57bdad6"
}

$ grpcurl -plaintext -proto ../proto/battleship.proto -d '{"gameID": "52917399-f6a6-4efe-871e-b8f313b61d50"}' localhost:50051 battleship.Battleship/JoinGame
{
  "playerID": "1c476478-daa0-4319-bd57-96aec884294c"
}

$ grpcurl -plaintext -proto ../proto/battleship.proto -d '{"gameID": "52917399-f6a6-4efe-871e-b8f313b61d50", "playerID": "e23b11a1-50e3-4659-a5a7-f2a1d57bdad6", "shipName": "Cruiser", "x": 0, "y": 0, "vertical": false}' localhost:50051 battleship.Battleship/PlaceShip
{
  
}
```
