/*
idea is to use dfs on the graph using a stack and color the graph
use a colorArr to keep track of it being colored and visited
use a set of vertices so that we can add vertices that are disconnected from the other
vertices
*/
func isBipartite(graph [][]int) bool {
	colorArr := make([]int, len(graph))
	stack := newStack()
	vertices := Set{}
	for i, _ := range graph {
		vertices.insert(i)
	}

	stack.push(0)
	colorArr[0] = 1
	for !stack.isEmpty() {
		curr := stack.top()
		fmt.Println(curr)
		stack.pop()
		vertices.remove(curr)
		for _, child := range graph[curr] {
			if colorArr[child] == 0 {
				colorArr[child] = colorChild(colorArr[curr])
				stack.push(child)
			} else {
				if colorArr[child] != colorChild(colorArr[curr]) {
					return false
				}
			}
		}
		if stack.isEmpty() {
			v := vertices.getAKey()
			if v != -1 {
				stack.push(v)
			}
		}
	}
	return true
}

type Stack struct {
	arr  []int
	size int
}

func newStack() Stack {
	return Stack{
		arr:  []int{},
		size: 0,
	}
}

func (s *Stack) push(e int) {
	s.arr = append(s.arr, e)
	s.size++
}

func (s *Stack) pop() {
	s.arr[s.size-1] = 0
	s.size--
	s.arr = s.arr[:s.size]
}

func (s *Stack) isEmpty() bool {
	return s.size == 0
}

func (s *Stack) top() int {
	return s.arr[s.size-1]
}

type Set map[int]bool

func (s *Set) insert(i int) {
	(*s)[i] = true
}

func (s *Set) remove(i int) {
	delete(*s, i)
}

func (s *Set) getAKey() int {
	for key, _ := range *s {
		return key
	}
	return -1
}
func (s *Set) isEmpty() bool {
	return len(*s) == 0
}

func (s *Set) exists(i int) bool {
	return (*s)[i]
}

func colorChild(i int) int {
	if i == 1 {
		return 2
	}
	return 1
}
