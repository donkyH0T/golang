package main

import (
	"fmt"
	"math"
)
type Node struct {
	value string
	index int
}
type ItemGraph struct {
	nodes []*Node
	edges map[*Node][]*Node
	matSmej [7][7]float64
}
func (n *Node) String() string {
	return fmt.Sprintf("%v", n.value)
}
// AddNode adds a node to the graph
func (g *ItemGraph) AddNode(n *Node) {
	g.nodes = append(g.nodes, n)
}
// AddEdge adds an edge to the graph
func (g *ItemGraph) AddEdge(n1, n2 *Node, trail float64) {
	if g.edges == nil {
		g.edges = make(map[*Node][]*Node)
	}
	g.matSmej[n1.index][n2.index] = trail
	g.matSmej[n2.index][n1.index] = trail
	g.edges[n1] = append(g.edges[n1], n2)
	g.edges[n2] = append(g.edges[n2], n1)
}
func (g *ItemGraph) String() {
	s := ""
	for i := 0; i < len(g.nodes); i++ {
		s += g.nodes[i].String() + " -> "
		near := g.edges[g.nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
}
// NodeQueue очередь узлов
type NodeQueue struct {
	items []*Node
}
// Enqueue добавляет узел в конец очереди
func (s *NodeQueue) Enqueue(t *Node) {
	s.items = append(s.items, t)
}
// BeginDelqueue удаляет узел из начала очереди
func (s *NodeQueue) BeginDelqueue() *Node {
	item := s.items[0]
	s.items = s.items[1:len(s.items)]
	return item
}
// Front возвращает следующий в очереди элемент, не удаляя его
func (s *NodeQueue) Front() *Node {
	item := s.items[0]
	return item
}
// IsEmpty возвращает true, если очередь пуста
func (s *NodeQueue) IsEmpty() bool {
	return len(s.items) == 0
}
// Size возвращает количество узлов в очереди
func (s *NodeQueue) Size() int {
	return len(s.items)
}
func (g *ItemGraph) Traverse() {
	q := NodeQueue{}
	q.items=[]*Node{}
	n := g.nodes[0]
	q.Enqueue(n)
	visited := make(map[*Node]bool)
	for {
		if q.IsEmpty() {
			break
		}
		node := q.BeginDelqueue()
		fmt.Println(node)
		near := g.edges[node]
		visited[node] = true
		for  h := 0; h< len(g.edges[node]); h++{
			if !visited[near[h]] {
				q.Enqueue(near[h])
				visited[near[h]] = true
			}
		}
	}
}
func (g *ItemGraph) Deycstra(start int,stop int) {
	weight := make([]float64,len(g.nodes))
	w:=0
	for w=0;w< len(weight);w++{
		weight[w]=math.Inf(1)
	}
	q := NodeQueue{}
	q.items=[]*Node{}
	n := g.nodes[start]
	fmt.Print(n," - ")
	weight[start] = 0
	steep:=g.nodes[stop]
	q.Enqueue(n)
	visited := make(map[*Node]bool)
	for {
		if q.IsEmpty() {
			break
		}
		node := q.BeginDelqueue()
		near := g.edges[node]
		visited[node] = true
		for  h := 0; h< len(g.edges[node]); h++{
			if !visited[near[h]] {
				q.Enqueue(near[h])
				w := weight[node.index] + g.matSmej[node.index][near[h].index]
				if w<weight[near[h].index]{
					weight[near[h].index] = w
				}
			}
		}
		if node==steep{
			fmt.Print(node," - ")
			break
		}
	}
	fmt.Println(weight[stop])
}
var g ItemGraph
func main() {
	var ind = 0
	nA := Node{"A", ind}
	ind++
	nB := Node{"B", ind}
	ind++
	nC := Node{"C", ind}
	ind++
	nD := Node{"D", ind}
	ind++
	nE := Node{"E", ind}
	ind++
	nF := Node{"F", ind}
	ind++
	nS := Node{"S", ind}
	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)
	g.AddNode(&nD)
	g.AddNode(&nE)
	g.AddNode(&nF)
	g.AddNode(&nS)
	g.AddEdge(&nA, &nB, 3)
	g.AddEdge(&nA, &nC, 10)
	g.AddEdge(&nB, &nE, 6)
	g.AddEdge(&nB, &nD, 5)
	g.AddEdge(&nB, &nF, 7)
	g.AddEdge(&nF, &nD, 4)
	g.AddEdge(&nA,&nS,12)
	g.String()
	g.Traverse()
	fmt.Println(g.matSmej[0])
	fmt.Println(g.matSmej[1])
	fmt.Println(g.matSmej[2])
	fmt.Println(g.matSmej[3])
	fmt.Println(g.matSmej[4])
	fmt.Println(g.matSmej[5])
	fmt.Println("A = 0")
	fmt.Println("B = 1")
	fmt.Println("C = 2")
	fmt.Println("D = 3")
	fmt.Println("E = 4")
	fmt.Println("F = 5")
	fmt.Println("S = 6")
	g.Deycstra(0,4)
}
