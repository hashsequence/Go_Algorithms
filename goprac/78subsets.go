func subsetsBetterSolution(nums []int) [][]int {
        sets := [][]int{[]int{}}
        for _, n := range nums {
                for _, set := range sets {
                        l := len(set)
                        ns := make([]int, l+1)
                        copy(ns, set)
                        ns[l] = n
                        sets = append(sets, ns)
                }
        }
        return sets
}



/*
my original solution
*/
func subsets(nums []int) [][]int {
    arrayOfSets := newArrOfSet()
    arrayOfSets.add(newSet())
    for _, val := range nums{
        s := newSet()
        s.add(val)
        arrayOfSets.add(s)
    }
    
   fmt.Println(arrayOfSets.toArray())
    
    for _, val := range nums {
        addingSets := copyArrOfSets(arrayOfSets)
        for _, currSet := range addingSets.arr {
            set := copySet(currSet)
            set.add(val)
            arrayOfSets.add(set)
        }
    }
    
    return arrayOfSets.toArray()
}

type Set struct {
    arr map[int]bool
}

func (this *Set) contains(value int) bool {
    return this.arr[value] 
}

func (this *Set) add(value int) {
    if !this.contains(value) {
        this.arr[value] = true
    } 
}

func newSet() Set {
    return Set{map[int]bool{}}
}

func copySet(original Set) Set {
    cop := newSet()
    for val, _ := range original.arr{
        cop.add(val)
    }
    return cop
}


type ArrOfSet struct {
    arr []Set
}

func (this *ArrOfSet) contains(set Set) bool {
    for _, currSet := range this.arr {
        flag := true
        for val, _:= range set.arr {
            if !currSet.contains(val) {
                flag = false
            }
        }   
        if flag && len(currSet.arr) == len(set.arr) {
            return true
        }
    }
    return false
}


func (this *ArrOfSet) add(set Set) {
    if !this.contains(set) {
        this.arr = append(this.arr, set)
    } 
}

func (this *ArrOfSet) toArray() [][]int {
    twoDArray := make([][]int,0)
    for i, _ := range this.arr {
        oneDArray := make([]int,0)
        for j, _ := range this.arr[i].arr {
            oneDArray = append(oneDArray, j)
        }
        twoDArray = append(twoDArray, oneDArray)
    }
    return twoDArray
}


func newArrOfSet() ArrOfSet {
    s := ArrOfSet{[]Set{}}
    return s
}

func copyArrOfSets(original ArrOfSet) ArrOfSet {
    cop := newArrOfSet()
    for i, _ := range original.arr {
        cop.add(original.arr[i])
    }
    return cop
}

