package main

import (
	"fmt"
	"math"
)

// PointND represents a point in n-dimensional space using a slice
type PointND struct {
	coords []float64
}

// Hypercube represents an n-dimensional hypercube
type Hypercube struct {
	n         int       // number of dimensions
	vertices  []PointND // 2^n vertices
	edges     [][2]int  // edges between vertices
	rotations []float64 // rotation angles for each plane
}

// NewHypercube creates a new n-dimensional hypercube
func NewHypercube(dimensions int) *Hypercube {
	if dimensions < 1 {
		dimensions = 1 // minimum dimension
	}

	// Calculate number of vertices (2^n)
	numVertices := 1 << dimensions

	// Generate vertices
	vertices := make([]PointND, numVertices)
	for i := 0; i < numVertices; i++ {
		coords := make([]float64, dimensions)
		for d := 0; d < dimensions; d++ {
			if (i & (1 << d)) == 0 {
				coords[d] = -1.0
			} else {
				coords[d] = 1.0
			}
		}
		vertices[i] = PointND{coords}
	}

	// Generate edges
	edges := [][2]int{}
	for i := 0; i < numVertices; i++ {
		for d := 0; d < dimensions; d++ {
			j := i ^ (1 << d)
			if j > i {
				edges = append(edges, [2]int{i, j})
			}
		}
	}

	rotations := make([]float64, (dimensions*(dimensions-1))/2)
	return &Hypercube{
		n:         dimensions,
		vertices:  vertices,
		edges:     edges,
		rotations: rotations,
	}
}

// Rotate performs rotation in all possible planes
func (h *Hypercube) Rotate(t float64) {
	angleIdx := 0
	for i := 0; i < h.n-1; i++ {
		for j := i + 1; j < h.n; j++ {
			angle := t * float64(angleIdx+1) / float64(h.n)
			h.rotations[angleIdx] = angle
			angleIdx++

			cosA := math.Cos(angle)
			sinA := math.Sin(angle)

			for v := range h.vertices {
				vi := h.vertices[v].coords[i]
				vj := h.vertices[v].coords[j]
				h.vertices[v].coords[i] = vi*cosA - vj*sinA
				h.vertices[v].coords[j] = vi*sinA + vj*cosA
			}
		}
	}
}

// ProjectTo3D projects n-dimensional coordinates to 3D space
func (h *Hypercube) ProjectTo3D() [][3]float64 {
	projection := make([][3]float64, len(h.vertices))

	for i, v := range h.vertices {
		var x, y, z float64
		if h.n > 0 {
			wFactor := 1.0 / (v.coords[h.n-1] + 2.0)
			x = v.coords[0] * wFactor
			if h.n > 1 {
				y = v.coords[1] * wFactor
			}
			if h.n > 2 {
				z = v.coords[2] * wFactor
			}
		}
		projection[i] = [3]float64{x, y, z}
	}
	return projection
}

// PrintFrame prints a frame with a configurable number of vertices
func PrintFrame(h *Hypercube, t float64, numVerticesToShow int) {
	projection := h.ProjectTo3D()

	fmt.Printf("\nFrame at t=%.2f\n", t)
	fmt.Printf("Showing %d of %d total vertices:\n",
		min(numVerticesToShow, len(projection)), len(projection))

	for i := 0; i < numVerticesToShow && i < len(projection); i++ {
		fmt.Printf("Vertex %d: [%.2f, %.2f, %.2f]\n",
			i, projection[i][0], projection[i][1], projection[i][2])
	}
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	// Configurable parameters
	dimensions := 5     // Number of dimensions
	verticesToShow := 8 // Number of vertices to display

	hypercube := NewHypercube(dimensions)

	fmt.Printf("Created %d-dimensional hypercube\n", dimensions)
	fmt.Printf("Number of vertices: %d\n", len(hypercube.vertices))
	fmt.Printf("Number of edges: %d\n", len(hypercube.edges))

	// Simulate rotation
	for t := 0.0; t < 2*math.Pi; t += 0.2 {
		hypercube.Rotate(t)
		PrintFrame(hypercube, t, verticesToShow)
	}
}
