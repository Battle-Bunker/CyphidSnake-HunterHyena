package main

import (
	"github.com/Battle-Bunker/cyphid-snake/agent"
)

/*
type GameSnapshot interface {
	GameID() string
	Rules() rules.Ruleset
	Turn() int
	Height() int
	Width() int
	Food() []rules.Point
	Hazards() []rules.Point
	Snakes() []SnakeSnapshot
	You() SnakeSnapshot
	Teammates() []SnakeSnapshot
	YourTeam() []SnakeSnapshot
	Opponents() []SnakeSnapshot
	ApplyMoves(moves []rules.SnakeMove) (GameSnapshot, error)
}

type SnakeSnapshot interface {
	ID() string
	Name() string
	Health() int
	Body() []rules.Point
	Head() rules.Point
	Length() int
	LastShout() string
	ForwardMoves() []rules.SnakeMove
}
*/

// TODO implement a heuristic that returns higher values when closer to food
func HeuristicFood(snapshot agent.GameSnapshot) float64 {
	if snapshot.You().Health() == 100 {
		return 10
	}
	foodDist := 0
	Headx := snapshot.You().Head().X 
	Heady := snapshot.You().Head().Y
	Neckx := snapshot.You().Body()[1].X
	Necky := snapshot.You().Body()[1].Y
	for i := 0; i < len(snapshot.Food()); i++ {
			foodx := snapshot.Food()[i].X
			foody := snapshot.Food()[i].Y
			foodPosx := Headx - foodx
			foodPosy := Heady - foody
			foodPospastx := Neckx - foodx
			foodPospasty := Necky - foody
			foodDistance := (foodPosx * foodPosx) + (foodPosy * foodPosy)
			foodDistancepast := (foodPospastx*foodPospastx)+(foodPospasty*foodPospasty)
			FoodClosenesses :=  (1/foodDistance - 1/foodDistancepast) *10
		
		
		if (FoodClosenesses > foodDist){
			foodDist = FoodClosenesses
		}
			
	}
	
	return float64(foodDist)
}
