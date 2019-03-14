func trap(height []int) int {
    //we are going to increment from the left and the right, so intiate two increment variables
    if len(height) < 3 {
        return 0
    }
    
    l := 0
    r := len(height)-1
    //we need to keep track of the left and right max, since we need to know know how much rain water can we collect at each point
    lMax := height[l]
    rMax := height[r]
    water := 0
    for l <= r {
        //we just need to check which one is taller
        if height[l] < height[r] {
            //since the right height, which is  currently set to the max on the right, is taller than the left max
            //we just need to compare with the left max height to determine how much water is on the current height
            if height[l] < lMax {
                water += lMax - height[l]
            } else {
                lMax = height[l]
            }
            l++
        } else {
            //since the left height, which is currently set to the max on the left, is taller than the right max
            //we just need to compare with the right max height to see how much water is on current height
            if height[r] < rMax {
                water += rMax - height[r]
            } else {
                rMax = height[r]
            }
            r--
        }
    }
    return water
}


