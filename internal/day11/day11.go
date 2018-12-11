package day11

import (
	"fmt"
	. "github.com/dr-horv/advent-of-code-2018/internal/pkg"
	"math"
)

type taskResult struct {
	X      int
	Y      int
	Value  int
	Size   int
}

type job func(Coordinate) taskResult

func worker(id int, jobs <-chan Coordinate, results chan<- taskResult, job job) {
	for j := range jobs {
		result := job(j)
		results <- result
	}
}

const SERIAL_NUMBER = 7315

func Solve(lines []string, partOne bool) string {

	grid := make(map[Coordinate]int, 300*300)
	size := 300

	for y := 1; y <= size; y++ {
		for x := 1; x <= size; x++ {
			c := Coordinate{X: x, Y: y}
			grid[c] = calculatePowerLevel(c)
		}
	}

	maxValue := math.MinInt64
	var maxCoord Coordinate
	if partOne {
		for y := 1; y <= (size - 2); y++ {
			for x := 1; x <= (size - 2); x++ {
				c := Coordinate{X: x, Y: y}
				value := calculateSquareValue(x, y, grid, 3)
				if value > maxValue {
					maxValue = value
					maxCoord = c
				}
			}
		}

		return maxCoord.String()
	} else {
		const TaskSize = 300 * 300
		messages := make(chan taskResult, TaskSize)
		jobs := make(chan Coordinate, TaskSize)

		task := func(c Coordinate) taskResult {
			values := CalculateSquareValues(c, grid, 300)
			max, i := FindMax(values)
			return taskResult{c.X, c.Y, max, i+1}
		}

		for w := 1; w <= 8; w++ {
			go worker(w, jobs, messages, task)
		}
		
		tasks := 0
		for y := 1; y <= size; y++ {
			for x := 1; x <= size; x++ {
				c := Coordinate{X: x, Y: y}
				jobs <- c
				tasks++
			}
		}

		finished := 0
		bestResult := taskResult{-1, -1, math.MinInt64, -1}

		for {
			taskResult := <-messages
			if taskResult.Value > bestResult.Value {
				bestResult = taskResult
			}
			finished++

			if finished%1000 == 0 {
				fmt.Printf("Progress: %v\n", float64(finished)/float64(tasks))
			}

			if finished == tasks {
				break
			}
		}

		fmt.Println(tasks)
		fmt.Println(bestResult.Value)
		return fmt.Sprintf("%v,%v,%v", bestResult.X, bestResult.Y, bestResult.Size)

	}
}

func calculateSquareValue(x int, y int, grid map[Coordinate]int, size int) int {
	sum := 0
	for dy := 0; dy < size; dy++ {
		for dx := 0; dx < size; dx++ {
			sum += grid[Coordinate{X: x + dx, Y: y + dy}]
		}
	}

	return sum
}

func CalculateSquareValues(c Coordinate, grid map[Coordinate]int, maxSize int) []int {
	results := make([]int, 1)
	results[0] = grid[Coordinate{X: c.X, Y: c.Y}]
	sum := results[0]
	size := 1
	for {
		if (c.X+size) > maxSize || (c.Y+size) > maxSize {
			break
		}

		for i := 0; i < size; i++ {
			sum += grid[Coordinate{X: c.X + size, Y: c.Y + i}]
			sum += grid[Coordinate{X: c.X + i, Y: c.Y + size}]
		}

		sum += grid[Coordinate{X: c.X + size, Y: c.Y + size}]
		results = append(results, sum)
		size++
	}

	return results
}

func calculatePowerLevel(c Coordinate) int {
	rackId := c.X + 10
	powerLevel := rackId * c.Y
	powerLevel += SERIAL_NUMBER
	powerLevel *= rackId
	powerLevel = (powerLevel / 100) % 10
	powerLevel -= 5

	return powerLevel
}
