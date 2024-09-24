package main

import (
	"github.com/Battle-Bunker/cyphid-snake/agent"
	"github.com/Battle-Bunker/cyphid-snake/server"
	"github.com/BattlesnakeOfficial/rules/client"
)

func main() {

	metadata := client.SnakeMetadataResponse{
		APIVersion: "1",
		Author:     "",
		Color:      "#0480d9",
		Head:       "default",
		Tail:       "default",
	}

	portfolio := agent.NewPortfolio(

		agent.NewHeuristic(1.0, "team-health", HeuristicHealth), 
		agent.NewHeuristic(2.0, "food-distance", HeuristicFood),
		
	)
//agent.NewHeuristic(10.0, "snakes", Heuristicsnakes),
	snakeAgent := agent.NewSnakeAgent(portfolio, metadata)
	server := server.NewServer(snakeAgent)

	server.Start()
}

