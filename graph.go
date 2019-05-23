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
	data      int
	status    VStatus
	inDegree  int
	outDegree int
	dTime     int
	fTime     int
	parent    int
	priority  int
}

type Edge struct {
	data   int
	weight int
	status EStatus
}

type GraphMatrix struct {
	vertex []Vertex
	edge   [][]Edge
}

func NewVertex(data int) *Vertex {
	return &Vertex{
		data:      data,
		inDegree:  0,
		outDegree: 0,
		status:    UNDISCOVERED,
		parent:    -1,
		fTime:     math.MaxInt32,
		dTime:     math.MaxInt32,
	}
}

func NewEdge(data int, weight int) *Edge {
	return &Edge{
		data:   data,
		weight: weight,
		status: UNDETERMINED,
	}
}

func NewGraphMatrix() *GraphMatrix {
	return &GraphMatrix{
		vertex: make([]Vertex, 0),
		edge:   make([][]Edge, 0),
	}
}

func (vertext *Vertex) exist(i int, j int) bool {
	return
}

func (vertex *Vertex) NextNbr(i int, j int) int {
for j != -1
}
