package day8

import (
	"fmt"
	"math"
	"robinlanderloos/aoc2025/io"
	"sort"
	"strconv"
	"strings"
)

type edge struct {
	from, to int
	distance float64
}

type unionFind struct {
	parent []int
	size   []int
}

func newUnionFind(length int) *unionFind {
	unionFind := &unionFind{
		parent: make([]int, length),
		size:   make([]int, length),
	}

	for i := range length {
		unionFind.parent[i] = i
		unionFind.size[i] = 1
	}

	return unionFind
}

func (unionFind *unionFind) find(x int) int {
	for unionFind.parent[x] != x {
		x = unionFind.parent[x]
	}

	return x
}

func (unionFind *unionFind) union(x, y int) {
	rootX := unionFind.find(x)
	rootY := unionFind.find(y)

	if rootX == rootY {
		return
	}

	if unionFind.size[rootX] < unionFind.size[rootY] {
		unionFind.size[rootY] += unionFind.size[rootX]
		unionFind.parent[rootX] = rootY
	} else {
		unionFind.size[rootX] += unionFind.size[rootY]
		unionFind.parent[rootY] = rootX
	}
}

type junction struct {
	x int
	y int
	z int
}

func Main() {
	// solve("day-8/example-input.txt", 10)
	// solve("day-8/input.txt")
	solveP2("day-8/input.txt")
}

func solveP2(path string) {
	junctions := parseInput(path)
	length := len(junctions)

	edges := []edge{}
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			distance := calculateDistance(junctions[i], junctions[j])
			edges = append(edges, edge{from: i, to: j, distance: distance})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distance < edges[j].distance
	})

	unionFind := newUnionFind(length)

	for _, edge := range edges {
		unionFind.union(edge.from, edge.to)
		root := unionFind.find(edge.from)
		if unionFind.size[root] == length {
			fmt.Printf("Result: %d\n", junctions[edge.from].x*junctions[edge.to].x)
			break
		}
	}

	// sizes := []int{}
	// for i := range junctions {
	// 	if unionFind.parent[i] == i {
	// 		sizes = append(sizes, unionFind.size[i])
	// 	}
	// }

	// sort.Slice(sizes, func(i, j int) bool {
	// 	return sizes[i] > sizes[j]
	// })

	// fmt.Print(sizes[0])
}

func solve(path string) {
	junctions := parseInput(path)
	length := len(junctions)

	edges := []edge{}
	for i := 0; i < length; i++ {
		for j := i + 1; j < length; j++ {
			distance := calculateDistance(junctions[i], junctions[j])
			edges = append(edges, edge{from: i, to: j, distance: distance})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distance < edges[j].distance
	})

	unionFind := newUnionFind(length)

	edgesProcessed := 0
	for _, edge := range edges {
		unionFind.union(edge.from, edge.to)
		edgesProcessed++
		if edgesProcessed == length {
			break
		}
	}

	sizes := []int{}
	for i := range junctions {
		if unionFind.parent[i] == i {
			sizes = append(sizes, unionFind.size[i])
		}
	}

	sort.Slice(sizes, func(i, j int) bool {
		return sizes[i] > sizes[j]
	})

	fmt.Print(sizes[0] * sizes[1] * sizes[2])
}

func calculateDistance(coordinate, other junction) float64 {
	dx := float64(coordinate.x - other.x)
	dy := float64(coordinate.y - other.y)
	dz := float64(coordinate.z - other.z)
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func parseInput(path string) []junction {
	result := make([]junction, 0)
	for line := range io.EnumerateFile(path) {
		coordinates := strings.Split(line, ",")
		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])
		z, _ := strconv.Atoi(coordinates[2])
		result = append(result, junction{
			x: x,
			y: y,
			z: z,
		})
	}

	return result
}
