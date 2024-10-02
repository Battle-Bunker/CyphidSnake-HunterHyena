package main

import (
	"container/heap"
	"math"
	"fmt"
	"github.com/Battle-Bunker/cyphid-snake/agent"
	"github.com/BattlesnakeOfficial/rules"
)

// HeuristicFood calculates a score based on food accessibility and survival potential
func HeuristicFood(snapshot agent.GameSnapshot) float64 {
	aliveSnakes := snapshot.YourTeam()

	if len(aliveSnakes) == 0 {
		return 0 // No alive snakes in our team, return 0
	}

	// Check if any of our snakes are out of bounds
	for _, snake := range aliveSnakes {
		if isOutOfBounds(snake.Head(), snapshot.Width(), snapshot.Height()) {
			return 0 // Snake is out of bounds, return 0
		}
	}
	for _, snake := range aliveSnakes {
			if isOutOfBounds(snake.Head(), snapshot.Width(), snapshot.Height()) {
					fmt.Printf("Snake out of bounds: Head=%v, Width=%d, Height=%d\n", snake.Head(), snapshot.Width(), snapshot.Height())
					return 0 // Snake is out of bounds, return 0
			}
	}

	// Pre-compute the occupancy grid once for the entire board
	occupancyGrid := buildOccupancyGrid(snapshot)

	var totalScore float64
	for _, allySnake := range aliveSnakes {
		if !allySnake.Alive() {
			continue
		}
		snakeScore := calculateSnakeScore(snapshot, allySnake, occupancyGrid)
		totalScore += snakeScore
	}
if math.IsNaN(totalScore) || math.IsInf(totalScore, 0) {
	return 0
}

return totalScore
	
}

// ... (rest of the code remains the same)

func isOutOfBounds(point rules.Point, width, height int) bool {
		return point.X < 0 || point.X >= width || point.Y < 0 || point.Y >= height
}

func calculateSnakeScore(snapshot agent.GameSnapshot, snake agent.SnakeSnapshot, occupancyGrid [][]bool) float64 {
	if snake.Health() == 100 {
		return 100 // Snake has just eaten, return a flat value
	}

	nearestFoodDistance := nearestFoodDistanceAStar(snapshot, snake.Head(), occupancyGrid)

	// Only calculate opponent distances if there are opponents
	var opponentDistanceSum float64
	opponents := snapshot.Opponents()
	if len(opponents) > 0 {
		opponentDistanceSum = sumOpponentDistancesToFood(snapshot, occupancyGrid)
		opponentDistanceSum /= float64(len(opponents))
	}

	foodAccessibilityScore := 100 - nearestFoodDistance + opponentDistanceSum
	healthFactor := float64(100-snake.Health()) / 100
	return foodAccessibilityScore * healthFactor
}

func buildOccupancyGrid(snapshot agent.GameSnapshot) [][]bool {
	width := snapshot.Width()
	height := snapshot.Height()
	grid := make([][]bool, width)
	for i := range grid {
		grid[i] = make([]bool, height)
	}

	for _, snake := range snapshot.AllSnakes() {
		for _, bodyPart := range snake.Body() {
			if bodyPart.X >= 0 && bodyPart.X < width && bodyPart.Y >= 0 && bodyPart.Y < height {
				grid[bodyPart.X][bodyPart.Y] = true
			}
		}
	}
	return grid
}

func nearestFoodDistanceAStar(snapshot agent.GameSnapshot, start rules.Point, occupancyGrid [][]bool) float64 {
	food := snapshot.Food()
	if len(food) == 0 {
		return 0 // No food on the board
	}

	minDistance := math.Inf(1)
	for _, foodItem := range food {
		distance := aStarDistance(snapshot, start, foodItem, occupancyGrid)
		if distance < minDistance {
			minDistance = distance
		}
	}
	return minDistance
}

func sumOpponentDistancesToFood(snapshot agent.GameSnapshot, occupancyGrid [][]bool) float64 {
	var totalDistance float64
	for _, opponent := range snapshot.Opponents() {
		if !opponent.Alive() {
			continue
		}
		totalDistance += nearestFoodDistanceAStar(snapshot, opponent.Head(), occupancyGrid)
	}
	return totalDistance
}

func aStarDistance(snapshot agent.GameSnapshot, start, goal rules.Point, occupancyGrid [][]bool) float64 {
	openSet := make(priorityQueue, 0)
	heap.Init(&openSet)
	heap.Push(&openSet, &node{point: start, fScore: heuristic(start, goal)})

	gScore := make(map[rules.Point]float64)
	gScore[start] = 0

	for openSet.Len() > 0 {
		current := heap.Pop(&openSet).(*node)

		if current.point == goal {
			return gScore[goal]
		}

		for _, neighbor := range getNeighbors(snapshot, current.point, occupancyGrid) {
			tentativeGScore := gScore[current.point] + 1

			if _, exists := gScore[neighbor]; !exists || tentativeGScore < gScore[neighbor] {
				gScore[neighbor] = tentativeGScore
				fScore := tentativeGScore + heuristic(neighbor, goal)
				heap.Push(&openSet, &node{point: neighbor, fScore: fScore})
			}
		}
	}

	return math.Inf(1) // No path found
}

func getNeighbors(snapshot agent.GameSnapshot, p rules.Point, occupancyGrid [][]bool) []rules.Point {
		neighbors := make([]rules.Point, 0, 4)
		directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

		for _, dir := range directions {
				newX, newY := p.X+dir[0], p.Y+dir[1]
				newPoint := rules.Point{X: newX, Y: newY}
				if !isOutOfBounds(newPoint, snapshot.Width(), snapshot.Height()) && !occupancyGrid[newX][newY] {
						neighbors = append(neighbors, newPoint)
				}
		}
		return neighbors
}

func heuristic(a, b rules.Point) float64 {
	dx, dy := float64(a.X-b.X), float64(a.Y-b.Y)
	return math.Sqrt(dx*dx + dy*dy)
}

type node struct {
	point  rules.Point
	fScore float64
}

type priorityQueue []*node

func (pq priorityQueue) Len() int            { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool  { return pq[i].fScore < pq[j].fScore }
func (pq priorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *priorityQueue) Push(x interface{}) { *pq = append(*pq, x.(*node)) }
func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
