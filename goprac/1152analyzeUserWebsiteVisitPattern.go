/*
1152. Analyze User Website Visit Pattern
Medium

29

201

Add to List

Share
We are given some website visits: the user with name username[i] visited the website website[i] at time timestamp[i].

A 3-sequence is a list of websites of length 3 sorted in ascending order by the time of their visits.  (The websites in a 3-sequence are not necessarily distinct.)

Find the 3-sequence visited by the largest number of users. If there is more than one solution, return the lexicographically smallest such 3-sequence.

 

Example 1:

Input: username = ["joe","joe","joe","james","james","james","james","mary","mary","mary"], timestamp = [1,2,3,4,5,6,7,8,9,10], website = ["home","about","career","home","cart","maps","home","home","about","career"]
Output: ["home","about","career"]
Explanation: 
The tuples in this example are:
["joe", 1, "home"]
["joe", 2, "about"]
["joe", 3, "career"]
["james", 4, "home"]
["james", 5, "cart"]
["james", 6, "maps"]
["james", 7, "home"]
["mary", 8, "home"]
["mary", 9, "about"]
["mary", 10, "career"]
The 3-sequence ("home", "about", "career") was visited at least once by 2 users.
The 3-sequence ("home", "cart", "maps") was visited at least once by 1 user.
The 3-sequence ("home", "cart", "home") was visited at least once by 1 user.
The 3-sequence ("home", "maps", "home") was visited at least once by 1 user.
The 3-sequence ("cart", "maps", "home") was visited at least once by 1 user.
 

Note:

3 <= N = username.length = timestamp.length = website.length <= 50
1 <= username[i].length <= 10
0 <= timestamp[i] <= 10^9
1 <= website[i].length <= 10
Both username[i] and website[i] contain only lowercase characters.
It is guaranteed that there is at least one user who visited at least 3 websites.
No user visits two websites at the same time.

solution:

first find a way to get the possible three sequence:
you can iterate through all possible iterations of the three sequence of each user 
by using a triple for loop i := 0 j := i+1 k:=j+1  
you can map the frequency of each triple sequence 

must sort the original triple array
must sort the triple sequences array

return the biggest value in the sequences array


next time do not implement interface sort, too many errors and pitfalls
*/
func mostVisitedPattern(username []string, timestamp []int, website []string) []string {
    seqMap:= map[string]*TripleSeq{}
    triplesArr := make(Triples,0)
    
    for i := 0; i < Min(len(username),len(timestamp),len(website)); i++ {
        triplesArr = append(triplesArr, Triple{username[i], timestamp[i], website[i]})
    }
    
    //sorting triples
    QSort(&triplesArr, 0, triplesArr.Len()-1)
 
    //getting all the unique triple sequences for each user
    currUser := triplesArr[0].username
    prevUser := triplesArr[0].username
    start := 0
    for i,_ := range triplesArr { 
        currUser = triplesArr[i].username
        if currUser != prevUser {
            GetThreeSequences(triplesArr, seqMap, start, i)
            start = i
        }
         prevUser = triplesArr[i].username
    }
    
    
    
    //get sequences for final user
    GetThreeSequences(triplesArr, seqMap, start, len(triplesArr))
    //tripleSeqs := make(TripleSeqs, 0)
    
    //get the most frequent one
    var mostCommon *TripleSeq
    for i,_ := range seqMap {
        if mostCommon == nil {
            mostCommon = seqMap[i]
        } else {
            curr := TripleSeqs{mostCommon, seqMap[i]}
            if curr.Less(0,1) {
                mostCommon = seqMap[i]
            }
        }
    
    }
    
    return mostCommon.seq
    
      
}


func GetThreeSequences(triples []Triple, seqMap map[string]*TripleSeq, start, end int) {
    if end-start+1 < 3 {
        return
    }
    isSeen := map[string]bool{}
    for i := start; i < end-2; i++ {
        for j := i+1; j < end-1; j++ {
            for k := j+1; k < end; k++ {
                val := "(" + triples[i].website + "," + triples[j].website + "," + triples[k].website  + ")"
                if _, ok := seqMap[val]; !ok {
                    seqMap[val] = &TripleSeq{[]string{triples[i].website, triples[j].website, triples[k].website}, 1, val}
                    isSeen[val] = true
                } else if _, ok := seqMap[val]; ok && !isSeen[val] {
                    seqMap[val].count++ 
                    isSeen[val] = true
                }
                
            }
        }
    }
}

type Sortable interface {
    Less(i, j int) bool
    Swap(i, j int)
    Len() int
    
}

type TripleSeq struct {
    seq []string
    count int
    key string
}

type Triple struct {
    username string
    timestamp int
    website string
}

type Triples []Triple

type TripleSeqs []*TripleSeq


func (triples *Triples) Less(i ,j int) bool {
    return ((*triples)[i].username < (*triples)[j].username) || ((*triples)[i].username == (*triples)[j].username && (*triples)[i].timestamp < (*triples)[j].timestamp) 
}

func (triples *Triples) Swap(i ,j int) {
    (*triples)[i], (*triples)[j] = (*triples)[j], (*triples)[i]
}

func (triples *Triples) Len() int {
    return len(*triples)
}


func (tripleSeqs *TripleSeqs) Less(i ,j int) bool {
    return ((*tripleSeqs)[i].count < (*tripleSeqs)[j].count) || 
    ((*tripleSeqs)[i].count == (*tripleSeqs)[j].count && (*tripleSeqs)[i].seq[0] > (*tripleSeqs)[j].seq[0]) || 
    ((*tripleSeqs)[i].count == (*tripleSeqs)[j].count && (*tripleSeqs)[i].seq[0] == (*tripleSeqs)[j].seq[0] && (*tripleSeqs)[i].seq[1] > (*tripleSeqs)[j].seq[1]) || 
    ((*tripleSeqs)[i].count == (*tripleSeqs)[j].count && (*tripleSeqs)[i].seq[0] == (*tripleSeqs)[j].seq[0] && (*tripleSeqs)[i].seq[1] == (*tripleSeqs)[j].seq[1] && (*tripleSeqs)[i].seq[2] > (*tripleSeqs)[j].seq[2])
}


func (tripleSeqs *TripleSeqs) Swap(i ,j int) {
    (*tripleSeqs)[i], (*tripleSeqs)[j] = (*tripleSeqs)[j], (*tripleSeqs)[i]
}

func (tripleSeqs *TripleSeqs) Len() int {
    return len(*tripleSeqs)
}



func QSort(triples Sortable, left, right int) {
    if right-left+1 < 2 {
        return
    }
    oLeft := left
    rand.Seed(time.Now().UnixNano())
    mid := rand.Intn(right-left+1) + left
    triples.Swap(right, mid)
    for i := left; i <= right; i++ {
        if triples.Less(i, right) {
            triples.Swap(i , left)
            left++
        } 
    }
    
    triples.Swap(left, right)
    mid = left
    QSort(triples, oLeft, mid-1)
    QSort(triples, mid+1, right)
    
}  

func Min(a ...int) int {
    m := a[0]
    for i, _ := range a {
        if a[i] < m {
            m = a[i]
        }
    }
    return m
}
