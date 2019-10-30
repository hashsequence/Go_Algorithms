/*
idea is cycle detection in a directed graph
we need dp as the recursion stack as go through a path in the graph
we need visited to keep track of which paths we already visited for example 
if vertex 3 is pointing to 3 adjacent vertices 1,2,4,
and we already visited the path down 1 and 2 then we know to go through 4 

for example if we go down one path and there is no cycle we return back up
the path in the graph by reseting the recrsion stack (dp) to false for that vertex

*/

func canFinish(numCourses int, prerequisites [][]int) bool {
    visited := make([]bool,numCourses)
    dp := make([]bool,numCourses)
    g := newGraph(numCourses)
    for _,edge := range prerequisites {
        g.insert(edge)
    }
    for v := 0; v < numCourses; v++ {
        if g.detectCycle(v,visited,dp) {
            return false
        }
    }
    return true
}

type Graph []map[int]bool

func newGraph(numVertices int) Graph {
    g := make([]map[int]bool,numVertices)
    return g
}
func (g *Graph) insert(edge []int) {
    if len(edge) == 2 {
        if ok := (*g)[edge[0]]; ok == nil {
            (*g)[edge[0]] = map[int]bool{}
        }
        (*g)[edge[0]][edge[1]] = true
    }
}

func (g *Graph) remove(v1, v2 int) {
    
    delete((*g)[v1],v2)
}

func (g *Graph) detectCycle(v int, visited, dp []bool) bool {
    
    
    visited[v] = true
    dp[v] = true
    
    for vertex, _ := range (*g)[v] {
        if !visited[vertex] && g.detectCycle(vertex,visited,dp) {
            return true
        } else if dp[vertex] {
            return true
        }
    }
    dp[v] = false 
    return false
}

