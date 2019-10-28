/**********************************************
idea:
Remove leaf nodes one level at a time until there are less than or equal to 2 nodes remaininig. The remaining nodes will be roots for min height trees

Why can't more than 2 nodes have minimum height?

  Z
  |
X---A---Y
 |    |
 V	 W

 We can prove that any path that is tied with XY in length, must go through A, and A will be the midpoint of each of them. Why? Suppose WZ=XY, but A is not the midpoint of WZ. In that case WA != AZ, which means that either WX or YZ is longer than XY (contradicting our premise). We can use the same technique to prove that VZ cannot be the same length as XY.

This proof also works if A is actually 2 nodes in the middle of an even length path. Either way, A is the solution to the problem.
***********************************************/

func findMinHeightTrees(n int, edges [][]int) []int {
	edgeSet := make(EdgeSet, n)
	for i := 0; i < n-1; i++ {
		edgeSet.insertDirectedEdge(edges[i][0], edges[i][1])
		edgeSet.insertDirectedEdge(edges[i][1], edges[i][0])
	}
	// fmt.Println("edgeSet",edgeSet)
	leafNodes := []int{} //get leaf nodes of the tree
	for i := 0; i < n; i++ {
		if len(edgeSet[i]) == 1 || len(edgeSet[i]) == 0 {
			leafNodes = append(leafNodes, i)
		}
	}
	//fmt.Println("leafNodes", leafNodes)
	for n > 2 { //terminates when we have only two non leaf nodes left
		temp := []int{} //will hold the new leaf nodes
		n -= len(leafNodes)
		for _, v1 := range leafNodes { //iterate through current leaf nodes
			v2 := edgeSet.popDirectedEdges(v1) //remove edge v1<---v2 from graph, and return the vertex that is still in the tree
			// fmt.Println(v2, len(edgeSet[v2]) )
			if len(edgeSet[v2]) == 1 { //if the vertex v2 you got is only connected to one other v1 then its a leaf
				temp = append(temp, v2) //we add this new leafnode to our new set
			}
		}
		leafNodes = temp //set the new list of leaf nodes
	}
	return leafNodes

}

type EdgeSet []map[int]bool

func (es *EdgeSet) popDirectedEdges(v1 int) int {
	v2 := 0
	for key, _ := range (*es)[v1] {
		v2 = key
		break
	}
	delete((*es)[v2], v1)
	return v2
}

func (es *EdgeSet) insertDirectedEdge(v1 int, v2 int) {
	if (*es)[v1] == nil {
		(*es)[v1] = map[int]bool{}
	}
	(*es)[v1][v2] = true
}

/*******************************Brute force***********************************************/
func findMinHeightTreesBruteForce(n int, edges [][]int) []int {
	min := 0
	heightMap := map[int][]int{}
	for i := 0; i < n; i++ {
		//	fmt.Println("i", i, "edges", edges)
		h := findHeightAtRoot(-1, i, edges)

		if i == 0 {
			min = h
		} else {
			if h < min {
				min = h
			}
		}

		if _, exists := heightMap[h]; exists {
			heightMap[h] = append(heightMap[h], i)
		} else {
			heightMap[h] = []int{i}
		}
	}
	//fmt.Println(heightMap)
	return heightMap[min]
}

func findHeightAtRoot(prevRoot int, root int, edges [][]int) int {
	maxHeight := 0
	if len(edges) == 1 && (root == edges[0][0] || root == edges[0][1]) && prevRoot != root {
		//fmt.Println(prevRoot, root, edges, "return 1")
		return 1
	}
	if len(edges) <= 1 || root == prevRoot {
		//fmt.Println(prevRoot, root, edges, "return 0")
		return 0
	}
	visited := false
	for i, _ := range edges {
		if root == edges[i][0] || root == edges[i][1] {
			nextRoot := edges[i][0]
			if edges[i][0] == root {
				nextRoot = edges[i][1]
			}
			curr := 0
			//fmt.Println("currRoot is ", root, " and nextRoot is ", nextRoot, " and edges are ", edges, " and edges length is", len(edges), " and i is ", i)
			if i == 0 {
				arr := myCopy(edges[i+1:])
				curr = findHeightAtRoot(root, nextRoot, arr) + 1
				maxHeight = curr
			} else if i == len(edges)-1 {
				arr := myCopy(edges[:len(edges)-1])
				curr = findHeightAtRoot(root, nextRoot, arr) + 1
			} else {
				arr := myCopy(append(append([][]int{}, edges[:i]...), edges[i+1:]...))
				curr = findHeightAtRoot(root, nextRoot, arr) + 1
			}
			if !visited {
				maxHeight = curr
				visited = true
			} else {
				if curr > maxHeight && curr != 0 {
					maxHeight = curr
				}
			}
		}
	}
	return maxHeight
}

func myCopy(src [][]int) [][]int {
	arr := make([][]int, len(src))
	for i, _ := range src {
		arr[i] = []int{src[i][0], src[i][1]}
	}
	return arr
}

