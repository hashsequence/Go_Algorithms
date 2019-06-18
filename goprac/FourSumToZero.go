/*********************************************************************
link to problem: https://www.geeksforgeeks.org/find-four-elements-sum-given-value-set-3-hashmap/
Find four elements that sum to a given value | Set 3 (Hashmap)
Given an array of integers, Check if there exist four elements at different indexes in the array whose sum is equal to a given value k.
*********************************************************************/
func FourSumToZero(nums []int) bool {
    dict := map[int][]Pair{}
    
    for i := 0; i < len(nums)-1; i++ {
        for j := i+1; j < len(nums); j++ {
            if _, ok := dict[nums[i]+nums[j]]; ok {
                if listOfPairs, ok := dict[-1*(nums[i]+nums[j])]; ok {
                    pair := Pair{i,j}
                    if pair.pairDoesNotShareIndexesWithExistingListOfPair(listOfPairs) {
                        return true
                    }
                } 
                dict[nums[i]+nums[j]] = append(dict[nums[i]+nums[j]], Pair{i,j})
            } else {
                dict[nums[i]+nums[j]] = []Pair{Pair{i,j}}
            }
        }
    }
    fmt.Println(dict)
    return false
}

type Pair [2]int

func (pair* Pair) pairDoesNotShareIndexesWithExistingListOfPair(arrOfPairs []Pair) bool {
    for _, currPair := range arrOfPairs {
        if ((*pair)[0] != currPair[0] && (*pair)[1] != currPair[1]) && ((*pair)[0] != currPair[1] && (*pair)[1] != currPair[0]) {
            fmt.Println(*pair,currPair)
           return true
       }
    }
    return false
}
