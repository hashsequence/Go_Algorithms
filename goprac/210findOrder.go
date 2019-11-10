/*
idea: use Khans Algorithm to topologically sort the graph
make an adjacency list using []map[int]bool
make a Queue that holds all vertex with indegree 0
need an inDegree to track in degree of each vertex
remember an in degree is how many directed edges is pointed to the vertex
visitedCnt tracks how many vertices are visited
if the visitedCnt does not equal the numCourses then there is a cycle
res holds each vertex but remember to append each new element to the front



*/
func findOrder(numCourses int, prerequisites [][]int) []int {
    res := []int{}
    g := NewGraph(numCourses, prerequisites)
    inDegree := make([]int,numCourses)
    q := GetVerticesWithNoPreReqs(g, inDegree)
    visitedCnt := 0
   
  
    for !q.IsEmpty() {
        curr := q.Top()
        q.Pop()
        res = append([]int{curr}, res...)
        
        for vertex,_ := range g[curr] {
            inDegree[vertex]--
            if inDegree[vertex] == 0 {
                q.Push(vertex)
            }
        }
        visitedCnt++
    }
    if visitedCnt != numCourses {
          return []int{}
     }
    return res
        
    
}


type Queue struct {
    arr []int
    size int
    
}

func NewQueue() Queue {
    return Queue{[]int{},0}
}

func (q *Queue) Push(k int) {
    (*q).arr = append(((*q).arr),k)
    (*q).size++
}


func (q *Queue) Pop() {
    if ((*q).size <= 0) {
        return 
    }
     if ((*q).size == 1) {
         (*q).arr = []int{}
         (*q).size = 0
         return
    }
    
    (*q).arr[0] = 0
    (*q).arr = (*q).arr[1:]
    (*q).size--
}

func (q *Queue) Top() int {
     if ((*q).size <= 0) {
        return -1
    }
    return (*q).arr[0]
}


func (q *Queue) IsEmpty() bool {
    return (*q).size == 0
}

type Graph []map[int]bool 

func NewGraph(size int, edges [][]int) Graph {
    g := make(Graph ,size)
    for _,edge := range edges {
        g.Insert(edge)   
    }
    return g
}

func (g *Graph) Insert(edge []int) {
    if len(edge) == 2 {
        if (*g)[edge[0]] == nil {
            (*g)[edge[0]] = map[int]bool{}
        }
        (*g)[edge[0]][edge[1]] = true
    }
}

func GetVerticesWithNoPreReqs(g Graph, inDegree []int ) Queue {
    q := NewQueue()
    for i,_ := range g {
        for vertex, _ := range g[i] {
            inDegree[vertex]++
        }
    }
    for vertex := range inDegree {
        if inDegree[vertex] == 0 {
            q.Push(vertex)
        }
    }
    return q
}
