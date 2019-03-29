//obvious O(n^2) solution would be iterate over every possible interval and calculate the max level
/*
there is however a O(N) solution
there are n^2 possible combination of intervals,
since half of the intervals are just reflexive, we can get rid of half
we can also get rid of N intervals for i==j

ideas:
1)any the thinner the containers the heigher the heights have to 
be to compensate for width

2)the smaller of the two pillars can never support water of a higher level, so we can get rid of it,
and shrink the container, so we iterate by incrementing the left pillar or decrementing the right pillar



*/
func maxArea(height []int) int {
    area := 0
    l, r := 0, len(height)-1
    for l < r {
        area = max(area, min(height[l],height[r])*(r-l))
        if height[l] > height[r] {
            r--
        } else {
            l++
        }
    }
    return area
}


func min(a int, b int) int {
    c := a
    if a > b {
        c = b
    }
    return c
}

func max(a int, b int) int {
    c := a
    if a < b {
        c = b
    }
    return c
}
