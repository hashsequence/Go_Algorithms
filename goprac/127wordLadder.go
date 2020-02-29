/*
127. Word Ladder
Medium

2454

1000

Add to List

Share
Given two words (beginWord and endWord), and a dictionary's word list, find the length of shortest transformation sequence from beginWord to endWord, such that:

Only one letter can be changed at a time.
Each transformed word must exist in the word list. Note that beginWord is not a transformed word.
Note:

Return 0 if there is no such transformation sequence.
All words have the same length.
All words contain only lowercase alphabetic characters.
You may assume no duplicates in the word list.
You may assume beginWord and endWord are non-empty and are not the same.
Example 1:

Input:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

Output: 5

Explanation: As one shortest transformation is "hit" -> "hot" -> "dot" -> "dog" -> "cog",
return its length 5.
Example 2:

Input:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

Output: 0

Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.
solution:
make an adjacency matrix
use a visited array to store depth at each visit
run a bfs traversal 

optimzation:
run bfs from beginWord and endWord to meet in the middle

actual solution:

My adjaceny matrix was wrong should have done a preprocessing step like:
map[*og:[dog log cog] *ot:[hot dot lot] c*g:[cog] co*:[cog] d*g:[dog] d*t:[dot] do*:[dot dog] h*t:[hot] ho*:[hot] l*g:[log] l*t:[lot] lo*:[lot log]]


*/

func ladderLength(beginWord string, endWord string, wordList []string) int {
    adjMatrix := make(map[string][]string)  
    isEndWordInList := false
    for i := 0; i < len(wordList); i++ {
        if endWord == wordList[i] {
            isEndWordInList = true
        }
        for j := 0; j < len(beginWord); j++ {
            adjMatrix[wordList[i][:j] + "*" + wordList[i][j+1:]] = append(adjMatrix[wordList[i][:j] + "*" + wordList[i][j+1:]], wordList[i])
        }
    } 
    if isEndWordInList == false {
        return 0
    }
    //fmt.Println(adjMatrix)
    
    //making 2 visited maps for the two traversal from beginWord and from endWord
    visited := map[string]int{}
    otherVisited := map[string]int{}
    //fmt.Println(adjMatrix)
   
    //make a queue for beginWord and push in and mark visited for beginWord
    q := make(Queue, 0)
    q.Push(&Item{beginWord, 1})
    visited[beginWord] = 1
    
    //make a queue and push and mark visited for endWord
    qOther := make(Queue, 0)
    qOther.Push(&Item{endWord, 1})
    otherVisited[endWord] = 1
    
    for len(q) > 0 && len(qOther) > 0{
        
        //traverse neighbors for the first queue
        currWord := q.Top()
        q.Pop()
        //loop through each iteration <firstPart>*<secondPart>
        for i := 0; i < len(beginWord); i++ {
            intermediate_word := currWord.value[:i] + "*" + currWord.value[i+1:]
            for _, neighbor := range adjMatrix[intermediate_word] {
                if otherVisited[neighbor] > 0 {
                    ans :=otherVisited[neighbor] + currWord.depth
                    return ans
                }   
            
                if visited[neighbor] == 0 {
                    q.Push(&Item{neighbor, currWord.depth+1})
                    visited[neighbor] = currWord.depth+1
                }    
            }
        }
        
        //traverse neighbors for second queue 
        currWordOther := qOther.Top()
        qOther.Pop()
        for i := 0; i < len(beginWord); i++ {
            intermediate_word := currWordOther.value[:i] + "*" + currWordOther.value[i+1:]
            for _, neighbor := range adjMatrix[intermediate_word] {
                if visited[neighbor] > 0 {
                    ans := visited[neighbor] + currWordOther.depth
                    return ans
                }   
            
                if otherVisited[neighbor] == 0 {
                    qOther.Push(&Item{neighbor, currWordOther.depth+1})
                    otherVisited[neighbor] = currWordOther.depth+1
        
                } 
            }
         }
    }
    return 0
    
}
func ladderLengthFirstWay(beginWord string, endWord string, wordList []string) int {
    //init adjacency matrix n^2
    adjMatrix := map[string][]string{}
    adjMatrix[beginWord] = make([]string,0)
    for i := 0; i < len(wordList) ;i++ {
        adjMatrix[wordList[i]] = make([]string,0)
    }
    if _, ok := adjMatrix[endWord]; !ok {
        return 0 
    }
    for i := 0; i < len(wordList); i++ {
        if wordDistance(beginWord, wordList[i]) == 1 {
            adjMatrix[beginWord] = append(adjMatrix[beginWord], wordList[i])
        }
    }
    for i := 0; i < len(wordList); i++ {
        for j := 0; j < len(wordList);j++ {
            if i != j && wordDistance(wordList[i], wordList[j]) == 1 {
                adjMatrix[wordList[i]] = append(adjMatrix[wordList[i]], wordList[j])
            }
        }
    }
   
    //making 2 visited maps for the two traversal from beginWord and from endWord
    visited := map[string]int{}
    otherVisited := map[string]int{}
    //fmt.Println(adjMatrix)
   
    //make a queue for beginWord and push in and mark visited for beginWord
    q := make(Queue, 0)
    q.Push(&Item{beginWord, 1})
    visited[beginWord] = 1
    
    //make a queue and push and mark visited for endWord
    qOther := make(Queue, 0)
    qOther.Push(&Item{endWord, 1})
    otherVisited[endWord] = 1
    
    for len(q) > 0 && len(qOther) > 0{
        
        //traverse neighbors for the first queue
        currWord := q.Top()
        q.Pop()
        
        for _, neighbor := range adjMatrix[currWord.value] {
            if otherVisited[neighbor] > 0 {
                ans :=otherVisited[neighbor] + currWord.depth
                return ans
            }   
            
            if visited[neighbor] == 0 {
                q.Push(&Item{neighbor, currWord.depth+1})
                visited[neighbor] = currWord.depth+1
            } 
        }
        
        //traverse neighbors for second queue 
        currWordOther := qOther.Top()
        qOther.Pop()
        
        for _, neighbor := range adjMatrix[currWordOther.value] {
            if visited[neighbor] > 0 {
                ans := visited[neighbor] + currWordOther.depth
                return ans
            }   
            
            if otherVisited[neighbor] == 0 {
                qOther.Push(&Item{neighbor, currWordOther.depth+1})
                otherVisited[neighbor] = currWordOther.depth+1
        
            } 
        }
    }
    return 0
    
}

func wordDistance(w1 string, w2 string) int {
    dist := 0
    for i := 0 ; i < len(w1); i++ {
        if w1[i] != w2[i] {
            dist++
        }
    }
    return dist
}

type Item struct {
    value string
    depth int
}

type Queue []*Item

func (q *Queue) Top() *Item {
    if len(*q) > 0 {
        return (*q)[0]
    }
    return nil
}

func(q *Queue) Pop() {
    if len(*q) > 0 {
        (*q) = (*q)[1:]
    }
}

func (q *Queue) Push(i *Item) {
    (*q) = append(*q,i)
}
