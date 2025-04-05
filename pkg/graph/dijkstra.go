package graph

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func (g *Graph) Dijkstra() <-chan string {
	out := make(chan string)

	go func() {
		defer close(out)

		var textResult string

		numVertices := len(g.Matrix)
		distances := make([]int, numVertices)
		visited := make([]bool, numVertices)
		previous := make([]int, numVertices)

		textResult += "Инициализация:\n"
		for i := range distances {
			previous[i] = -1
			if i != g.StartVertex {
				distances[i] = math.MaxInt64
				textResult += fmt.Sprintf("'%d': ∞; ", i)
				continue
			}
			distances[g.StartVertex] = 0
			textResult += fmt.Sprintf("'%d': 0; ", i)
		}

		textResult += "\n\n"
		for i := 0; i < numVertices; i++ {
			textResult += fmt.Sprintf("Шаг №%d\n", i+1)
			u := minDistances(distances, visited)
			textResult += fmt.Sprintf("Текущая вершина: '%d'\n", u)
			visited[u] = true

			textResult += "Обновляем расстояния для вершин.\n"
			for v := 0; v < numVertices; v++ {
				textResult += fmt.Sprintf(" - Вершина '%d': ", v)

				if !visited[v] && g.Matrix[u][v] != 0 && distances[u] != math.MaxInt64 {
					textResult += fmt.Sprintf("min(%d + %d, %d) = ", distances[u], g.Matrix[u][v], distances[v])
					if distances[u]+g.Matrix[u][v] < distances[v] {
						distances[v] = distances[u] + g.Matrix[u][v]
						previous[v] = u
					}
				}

				textResult += fmt.Sprintf("%d\n", distances[v])
			}

			textResult += "Посещенные вершины:"
			for vert, isVisited := range visited {
				if isVisited {
					textResult += fmt.Sprintf(" '%d'", vert)
				}
			}
			textResult += ";\n\n"
		}

		textResult += "Итог:\n"

		resultDistance := distances[g.EndVertex]

		path := []int{}
		for at := g.EndVertex; at != -1; at = previous[at] {
			path = append([]int{at}, path...)
		}

		var resultPath string
		for at := g.EndVertex; at != -1; at = previous[at] {
			resultPath = strconv.Itoa(at) + "->" + resultPath
		}
		resultPath = strings.TrimSpace(resultPath)

		textResult += fmt.Sprintf("Кратчайшее расстояние от вершины %d до вершины %d: %d\nПуть: %s", g.StartVertex, g.EndVertex, resultDistance, resultPath[:len(resultPath)-2])
		textResult = strings.Replace(textResult, fmt.Sprintf("%d", math.MaxInt64), "∞", -1)

		out <- textResult
	}()

	return out
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
