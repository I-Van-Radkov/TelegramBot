package graph

type Graph struct {
	Matrix      [][]int
	StartVertex int
	EndVertex   int
}

func (g *Graph) NewGraph(matrix [][]int, start, end int) *Graph {
	return &Graph{
		Matrix:      matrix,
		StartVertex: start,
		EndVertex:   end,
	}
}
