package main

import (
	"container/heap"
	"fmt"
	"os"
	"strings"
)

type State struct {
	pos Position
	dir Direction
}

type Position struct {
	x, y int
}

type Direction struct {
	dx, dy int
}

type PriorityQueue []Node

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].distance < pq[j].distance }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Node))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) AddNode(newState State, newDistance int, dMap map[State]int) {
	if dist, exists := dMap[newState]; !exists || dist > newDistance {
		dMap[newState] = newDistance
		heap.Push(pq, Node{state: newState, distance: newDistance})
	}
}

type Node struct {
	state    State
	distance int
}

func solve(start State, opSign bool, grid [][]rune) map[State]int {
	directions := []Direction{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	n, m := len(grid), len(grid[0])
	dMap := make(map[State]int)
	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, Node{state: start, distance: 0})
	dMap[start] = 0

	for pq.Len() > 0 {
		cur := heap.Pop(pq).(Node)

		if dist, exists := dMap[cur.state]; exists && dist < cur.distance {
			continue
		}

		curPos, curDir := cur.state.pos, cur.state.dir
		newSign := 1
		if opSign {
			newSign = -1
		}
		newPos := Position{curPos.x + newSign*curDir.dx, curPos.y + newSign*curDir.dy}
		if newPos.x >= 0 && newPos.x < n && newPos.y >= 0 && newPos.y < m && grid[newPos.x][newPos.y] != '#' {
			newState := State{pos: newPos, dir: curDir}
			pq.AddNode(newState, cur.distance+1, dMap)
		}

		var orthogonalDirs []Direction
		if curDir == directions[0] || curDir == directions[1] {
			orthogonalDirs = directions[2:]
		} else {
			orthogonalDirs = directions[:2]
		}
		for _, od := range orthogonalDirs {
			newState := State{pos: curPos, dir: od}
			pq.AddNode(newState, cur.distance+1000, dMap)
		}
	}

	return dMap
}

func main() {
	input, _ := os.ReadFile("solutions/2024/Day16/input.txt")
	lines := strings.Split(strings.TrimSpace(string(input)), "\n")
	var grid [][]rune
	for _, line := range lines {
		grid = append(grid, []rune(line))
	}

	n, m := len(grid), len(grid[0])
	var startPos, endPos Position
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			switch grid[i][j] {
			case 'S':
				startPos = Position{i, j}
			case 'E':
				endPos = Position{i, j}
			}
		}
	}

	directions := []Direction{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}
	initSolve := solve(State{startPos, directions[0]}, false, grid)
	minAns := int(^uint(0) >> 1) // Maximum integer
	for _, dir := range directions {
		if d, exists := initSolve[State{endPos, dir}]; exists {
			if d < minAns {
				minAns = d
			}
		}
	}
	fmt.Println(minAns)

	var ndSolves []map[State]int
	for _, dir := range directions {
		ndSolves = append(ndSolves, solve(State{endPos, dir}, true, grid))
	}

	ansSet := make(map[Position]struct{})
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			for _, dir := range directions {
				curState := State{Position{i, j}, dir}
				good := false
				for _, ndSolve := range ndSolves {
					if initD, exists1 := initSolve[curState]; exists1 {
						if ndD, exists2 := ndSolve[curState]; exists2 {
							if initD+ndD == minAns {
								good = true
								break
							}
						}
					}
				}
				if good {
					ansSet[curState.pos] = struct{}{}
				}
			}
		}
	}
	fmt.Println(len(ansSet))
}
