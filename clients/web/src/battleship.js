import { BattleshipClient } from './proto/battleship_grpc_web_pb';
import * as pb from './proto/battleship_pb';

let client = new BattleshipClient('http://localhost:8080');

export function newGame() {
  let request = new pb.NewGameRequest();

  return new Promise((resolve, reject) => {
    client.newGame(request, {}, function (err, response) {
      if (err) {
        reject(err);
        return;
      }

      resolve({
        gameID: response.getGameid(),
        playerID: response.getPlayerid(),
      });
    });
  });
}

export function joinGame(gameID) {
  let request = new pb.JoinGameRequest();
  request.setGameid(gameID);

  return new Promise((resolve, reject) => {
    client.joinGame(request, {}, function (err, response) {
      if (err) {
        reject(err);
        return;
      }

      resolve({ playerID: response.getPlayerid() });
    });
  });
}

export function placeShip(gameID, playerID, name, x, y, vertical) {
  let request = new pb.PlaceShipRequest();
  request.setGameid(gameID);
  request.setPlayerid(playerID);
  request.setShipname(name);
  request.setX(x);
  request.setY(y);
  request.setVertical(vertical);

  return new Promise((resolve, reject) => {
    client.placeShip(request, {}, function (err, response) {
      if (err) {
        reject(err);
        return;
      }

      resolve({});
    });
  });
}
