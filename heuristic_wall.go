package main

import (
  "github.com/Battle-Bunker/cyphid-snake/agent"
)
/*
type GameSnapshot interface {
  GameID() string
  Rules() .Ruleset
  Turn() int
  Height() int
  Width() int
  Food() [].Point
  Hazards() [].Point
  Snakes() []SnakeSnapshot
  You() SnakeSnapshot
  Teammates() []SnakeSnapshot
  YourTeam() []SnakeSnapshot
  Opponents() []SnakeSnapshot
  ApplyMoves(moves [].SnakeMove) (GameSnapshot, error)
}

type SnakeSnapshot interface {
  ID() string
  Name() string
  Health() int
  Body() [].Point
  Head() .Point
  Length() int
  LastShout() string
  ForwardMoves() [].SnakeMove
}
*/

func HeuristicWall(snapshot agent.GameSnapshot) float64 {
  dead := 0
  
  Headx := snapshot.You().Head().X 
  Heady := snapshot.You().Head().Y
  wallx := snapshot.Height()
  wally := snapshot.Width()
  if Headx == -1 {
    dead = 1
  }
  if Heady == -1 {
    dead = 1
  }
  if Headx == wallx {
    dead = 1
  }
  if Heady == wally {
    dead = 1
  }
  death := (dead * -1000)
  return float64(death)
}