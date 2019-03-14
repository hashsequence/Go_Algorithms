func trapRainWater(heightMap [][]int) int {
    if len(heightMap) < 3 {
        return 0
    }'
    if heightMap != nil {
        if len(heightMap[0]) < 3 {
            return 0
        }
    }
    
    horizontalAcc := 0
    verticalAcc := 0
    uMax := 0
    dMax := 0
    lMax := 0
    rMax := 0
    
    waterMap := make([][]int, len(heightMap))
    for i, _ := range heightMap[0] {
        waterMap[i] = make([]int, len(heightMap[i]))
    } 
    
    u, d, l, r := 0, len(heightMap), 0, len(heightMap[0])
    for i := 0; i < len(heightMap[0]); i++ {
        l, r = 0, len(heightMap[i])
        for l <= r {
            if heightMap[l] < heightMap[r] {
                if heightMap[l] < lMax {
                     waterMap[i][l] = lMax - heightMap[l]
                 } else {
                  lMax = heightMap[l]
                 }
                l++
            } else {
                 if heightMap[r] < rMax {
                     waterMap[i][r] = rMax - heightMap[r]
                 } else {
                    rMax = heightMap[r]
                }
                r--
            }
        }
    }
    
   
    
    
    
}
