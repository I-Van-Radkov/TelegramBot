package graph

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func (g *Graph) Dijkstra(channel chan string) {
	numVertices := len(g.Matrix)
	distances := make([]int, numVertices)
	visited := make([]bool, numVertices)
	previous := make([]int, numVertices)

	for i := range distances {
		distances[i] = math.MaxInt64
		previous[i] = -1
	}
	distances[g.StartVertex] = 0

	for i := 0; i < numVertices; i++ {
		u := minDistances(distances, visited)
		visited[u] = true

		for v := 0; v < numVertices; v++ {
			if !visited[v] && g.Matrix[u][v] != 0 && distances[u] != math.MaxInt64 && distances[u]+g.Matrix[u][v] < distances[v] {
				distances[v] = distances[u] + g.Matrix[u][v]
				previous[v] = u
			}
		}
	}

	resultDistance := distances[g.EndVertex]

	path := []int{}
	for at := g.EndVertex; at != -1; at = previous[at] {
		path = append([]int{at}, path...)
	}

	var resultPath string
	for at := g.EndVertex; at != -1; at = previous[at] {
		resultPath = " " + strconv.Itoa(at) + resultPath
	}
	channel <- fmt.Sprintf("Кратчайшее расстояние от вершины %d до вершины %d: %d\nПуть: %s", g.StartVertex, g.EndVertex, resultDistance, strings.TrimSpace(resultPath))
}

func minDistances(distances []int, visited []bool) int {
	min := math.MaxInt64
	minIndex := -1

	for v := 0; v < len(distances); v++ {
		if !visited[v] && distances[v] <= min {
			min = distances[v]
			minIndex = v
		}
	}

	return minIndex
}
