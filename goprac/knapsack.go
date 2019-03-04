package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {

	val := []int{60, 100, 120}
	wt := []int{10, 20, 30}
	w := 50
	m, s := knapsack(wt, val, w)
	fmt.Println(m)
	fmt.Println(s)
}

type Sack struct {
	w int
	v int
}

func knapsack(weights []int, values []int, W int) (int, []Sack) {
	//initializing Value Array
	N := len(values)
	V := make([][]int, N+1)
	for i, _ := range V {
		V[i] = make([]int, W+1)
	}

	//of there are no items then the value would be 0
	for j := 0; j < W+1; j++ {
		V[0][j] = 0
	}

	//if the max capacity is 0 then the value would be 0
	for i := 0; i < N+1; i++ {
		V[i][0] = 0
	}

	for item := 1; item < N+1; item++ {
		for weight := 1; weight < W+1; weight++ {
			//Given a weight, check if the value of the current
			//item + value of the item that we could afford
			//with the remaining weight is greater than the value
			//without the current item itself
			if weights[item-1] <= weight {
				V[item][weight] = max(values[item-1]+V[item-1][weight-weights[item-1]], V[item-1][weight])
			} else {
				//If the current item's weight is more than the
				//running weight, just carry forward the value
				//without the current item
				V[item][weight] = V[item-1][weight]
			}
		}
	}

	res := V[N][W]
	s := make([]Sack, 0)

	currWeight := W
	for i := N; i > 0 && res > 0; i-- {
		if res != V[i-1][currWeight] {
			s = append(s, Sack{weights[i-1], values[i-1]})
			res -= values[i-1]
			currWeight -= weights[i-1]
		}
	}

	//return max value
	return V[N][W], s

}

func max(vals ...int) int {

	maxVal := vals[0]
	for _, val := range vals {
		if maxVal < val {
			maxVal = val
		}
	}
	return maxVal
}
