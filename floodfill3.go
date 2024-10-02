package main

import (
  "fmt"
  //"math"
  "github.com/Battle-Bunker/cyphid-snake/agent"
  "github.com/BattlesnakeOfficial/rules"
)

// HeuristicFreeSpace calculates a score based on the amount of free space accessible to each snake
func HeuristicFreeSpace(snapshot agent.GameSnapshot) float64 {
  aliveSnakes := snapshot.YourTeam()

  if len(aliveSnakes) == 0 {
    fmt.Println("No alive snakes in our team")
    return 0
  }

  occupancyGrid := buildOccupancyGrid(snapshot)

  var totalFreeSpace float64
  for _, allySnake := range aliveSnakes {
    if !allySnake.Alive() {
      continue
    }
    freeSpace := calculateSnakeFreeSpace(snapshot, allySnake, occupancyGrid)
    fmt.Printf("Snake at %v has free space: %f\n", allySnake.Head(), freeSpace)
    totalFreeSpace += freeSpace
  }

  fmt.Printf("Total free space: %f\n", totalFreeSpace)
  return totalFreeSpace
}

func calculateSnakeFreeSpace(snapshot agent.GameSnapshot, snake agent.SnakeSnapshot, occupancyGrid [][]bool) float64 {
  head := snake.Head()
  if isOutOfBounds(head, snapshot.Width(), snapshot.Height()) {
    fmt.Printf("Snake head is out of bounds: %v\n", head)
    return 0
  }

 freeSpaceCount := floodFill(snapshot, head, occupancyGrid)
 // totalBoardSize := snapshot.Width() * snapshot.Height()
  //freeSpaceRatio := float64(freeSpaceCount) / float64(totalBoardSize)

  // Return the raw free space count, not a scaled score
  return float64(freeSpaceCount)
}

func floodFill(snapshot agent.GameSnapshot, start rules.Point, occupancyGrid [][]bool) int {
  width, height := snapshot.Width(), snapshot.Height()
  visited := make([][]bool, width)
  for i := range visited {
    visited[i] = make([]bool, height)
  }

  queue := []rules.Point{start}
  freeSpaceCount := 0

  for len(queue) > 0 {
    current := queue[0]
    queue = queue[1:]

    if isOutOfBounds(current, width, height) || occupancyGrid[current.X][current.Y] || visited[current.X][current.Y] {
      continue
    }

    visited[current.X][current.Y] = true
    freeSpaceCount++

    // Add neighboring cells to the queue
    neighbors := []rules.Point{
      {X: current.X + 1, Y: current.Y},
      {X: current.X - 1, Y: current.Y},
      {X: current.X, Y: current.Y + 1},
      {X: current.X, Y: current.Y - 1},
    }

    for _, neighbor := range neighbors {
      queue = append(queue, neighbor)
    }
  }

  return freeSpaceCount
}

// Assume buildOccupancyGrid and isOutOfBounds are defined elsewhere in your project