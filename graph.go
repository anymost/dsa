package main

import "math"

type VStatus int
type EStatus int

const (
	UNDISCOVERED VStatus = iota
	DISCOVERED
	VISITED
)

const (
	UNDETERMINED EStatus = iota
	TREE
	CROSS
	FORWARD
	BACKWARD
)

type Vertex struct {
	Data      int
	Status    VStatus
	InDegree  int
	OutDegree int
	DTime     int
	FTime     int
	Parent    int
	Priority  int
}

type Edge struct {
	Data   int
	Weight int
	Status EStatus
}

type GraphMatrix struct {
	Vertex []*Vertex
	Edge   [][]*Edge
}

func NewVertex(data int) *Vertex {
	return &Vertex{
		Data:      data,
		InDegree:  0,
		OutDegree: 0,
		Status:    UNDISCOVERED,
		Parent:    -1,
		FTime:     math.MaxInt32,
		DTime:     math.MaxInt32,
	}
}

func NewEdge(data int, weight int) *Edge {
	return &Edge{
		Data:   data,
		Weight: weight,
		Status: UNDETERMINED,
	}
}

func NewGraphMatrix() *GraphMatrix {
	return &GraphMatrix{
		Vertex: make([]*Vertex, 0),
		Edge:   make([][]*Edge, 0),
	}
}

func (graph *GraphMatrix) IsExistEdge(i, j int) bool {
	if 0 <= i && i < len(graph.Vertex) && 0 <= j && j < len(graph.Vertex) {
		return graph.Edge[i][j] != nil
	}
	return false
}

func (graph *GraphMatrix) NextNeighbor(i, j int) int {
	for -1 < j {
		j--
		if graph.IsExistEdge(i, j) {
			return j
		}
	}
	return -1
}

func (graph *GraphMatrix) FirstNeighbor(i int) int {
	return graph.NextNeighbor(i, len(graph.Edge[i]))
}

func (graph *GraphMatrix) InsertEdge(i, j int, edge *Edge) {
	graph.Edge[i][j] = edge
	graph.Vertex[i].InDegree++
	graph.Vertex[j].InDegree++
}

func (graph *GraphMatrix) RemoveEdge(i, j int) {
	graph.Edge[i][j] = nil
	graph.Vertex[i].InDegree--
	graph.Vertex[j].InDegree++
}

func (graph *GraphMatrix) InsertVertex(v *Vertex) {
	for i := 0; i < len(graph.Edge); i++ {
		graph.Edge[i] = append(graph.Edge[i], nil)
	}
	newEdges := make([]*Edge, len(graph.Vertex)+1)
	graph.Vertex = append(graph.Vertex, v)
	graph.Edge = append(graph.Edge, newEdges)
}

func (graph *GraphMatrix) removeVertex(i int) {
	newVertex := make([]*Vertex, len(graph.Vertex)-1)
	for k := 0; k < i; k++ {
		newVertex[k] = graph.Vertex[k]
	}
	for k := i + 1; k < len(graph.Vertex); k++ {
		newVertex[k-1] = graph.Vertex[k]
	}
	graph.Vertex = newVertex

	for i := 0; i < len(graph.Edge); i++ {
		graph.Edge[i] = graph.Edge[i][0 : len(graph.Edge)-1]
	}
	graph.Edge = graph.Edge[0 : len(graph.Edge)-1]
}

func (graph *GraphMatrix) BFS(v int, clock *int, visitor func(v int)) {
	queue := NewQueue()
	graph.Vertex[v].Status = DISCOVERED
	queue.Enqueue(v)
	for !queue.Empty() {
		v := queue.Dequeue().(int)
		vertex := graph.Vertex[v]
		*clock++
		vertex.DTime = *clock
		visitor(vertex.Data)
		for i := graph.FirstNeighbor(v); -1 < i; i = graph.NextNeighbor(v, i) {
			vertex := graph.Vertex[i]
			edge := graph.Edge[v][i]
			if vertex.Status == UNDISCOVERED {
				vertex.Status = DISCOVERED
				queue.Enqueue(i)
				edge.Status = TREE
				vertex.Parent = v
			} else {
				edge.Status = CROSS
			}
		}
		vertex.Status = VISITED
	}
}

func (graph *GraphMatrix) DFS(v int, clock *int, visitor func(v int)) {
	vertex := graph.Vertex[v]
	*clock++
	vertex.DTime = *clock
	vertex.Status = DISCOVERED
	for i := graph.FirstNeighbor(v); -1 < i; i = graph.NextNeighbor(v, i) {
		status := graph.Vertex[i].Status
		switch status {
		case UNDISCOVERED:
			graph.Edge[v][i].Status = TREE
			graph.Vertex[i].Status = DISCOVERED
			graph.Vertex[i].Parent = v
			graph.DFS(i, clock, visitor)
		case DISCOVERED:
			graph.Edge[v][i].Status = BACKWARD
		default:
			vDTime := graph.Vertex[v].DTime
			iDTime := graph.Vertex[i].DTime
			if vDTime < iDTime {
				graph.Edge[v][i].Status = FORWARD
			} else {
				graph.Edge[v][i].Status = CROSS
			}
		}
	}
	visitor(vertex.Data)
	vertex.Status = VISITED
	*clock++
	vertex.FTime = *clock
}
