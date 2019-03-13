func trap(height []int) int {
    //horizontal displacement
    x := 0
    //total water accumulated 
    total := 0
    //max height
    maxH := 0
    //tabs of spaces we filled with blocks or rain
    tab := 0
    
    for i, h := range height {
        //check if there is a valley for rainwater 
        if (height[i] > height[i-1] && maxH > h && i > 0) {
            //calculate raintotal
            calculateRainTotal
        }
        if h >= maxH {
            x = 0
            tab = 0
        }
        x++
        
    }
    
}

func calculateRainTotal(maxH, h, x, tab int) int {
    return min(maxh, h) * x - tab
}

func min(a int, b int) int {
	c := b
	if a < b {
		c = a
	}
	return c
}

func max(a int, b int) int {
	c := b
	if a > b {
		c = a
	}
	return c
}
