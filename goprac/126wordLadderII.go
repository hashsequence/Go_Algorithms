/*
126. Word Ladder II
Hard


Given two words (beginWord and endWord), and a dictionary's word list, find all shortest transformation sequence(s) from beginWord to endWord, such that:

Only one letter can be changed at a time
Each transformed word must exist in the word list. Note that beginWord is not a transformed word.
Note:

Return an empty list if there is no such transformation sequence.
All words have the same length.
All words contain only lowercase alphabetic characters.
You may assume no duplicates in the word list.
You may assume beginWord and endWord are non-empty and are not the same.
Example 1:

Input:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

Output:
[
  ["hit","hot","dot","dog","cog"],
  ["hit","hot","lot","log","cog"]
]
Example 2:

Input:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

Output: []

Explanation: The endWord "cog" is not in wordList, therefore no possible transformation.

solution:

the idea is to use a bfs but we want to use a queue of Items defined as
type Item struct {
    s string
    visited map[string]bool
    parent *Item
    depth int
}

we need to send a copy of the visited map for every level of the traversal
we need to keep track of parent nodes to backtrack and get results when done
we need to break if got all the results for minDepth 

edge case:
if beginWord is in the wordList we must remember to construct the adjacency list without the duplicate neighbors

note: there would never be a cycle in our traversal if we dont deep copy the visited array because we are using bfs, so we will always hit the end node of the shortest path first, so we never need to account for the path that needs to go back to a already visited node in our visited array, otherwise it wont be the shortest path first, which means it is not bfs in the first place
*/

func findLadders(beginWord string, endWord string, wordList []string) [][]string {
    adj := map[string][]string{}
    
    for i, _ := range wordList {
       adj[wordList[i]] = make([]string, 0)
    }
    if _, ok := adj[beginWord]; !ok {
        for i, _ := range wordList {
             if Distance(beginWord,wordList[i]) == 1 {
                adj[beginWord] = append(adj[beginWord],wordList[i])
            }
        }
    }
    
     for i, _ := range wordList {
         for j, _ := range wordList {
             if i != j {
                 if Distance(wordList[i], wordList[j]) == 1 && wordList[j] != beginWord {
                     adj[wordList[i]] = append(adj[wordList[i]], wordList[j])
                 }
             }
         }
    }
    
    if _, ok := adj[endWord]; !ok {
        return [][]string{}
    }
    return Bfs(beginWord, endWord, wordList, adj)
    
    
}

func Distance(s1, s2 string) int {
    diff := 0
    for i, _ := range s1 {
        if s1[i] != s2[i] {
            diff++
        }
    }
    return diff
}

func Bfs(beginWord string, endWord string, wordList []string, adj map[string][]string) [][]string {
    q := make(Queue, 0)
    q.Push(&Item{beginWord, make(map[string]bool, len(wordList)+1), nil, 0})
    
    var endNodes []*Item
    minDepth := -1
    for len(q) > 0 {
        curr := q.Front()
        q.Pop()
        curr.visited[curr.s] = true
        if curr.s == endWord {
            if minDepth == -1 {
                minDepth = curr.depth
            } else if minDepth != -1 && curr.depth > minDepth {
                break
            }
            endNodes = append(endNodes, curr)
        } else {
            cVisited := curr.visited
            for _, next := range adj[curr.s] {
                if !curr.visited[next] {
                    q.Push(&Item{next, cVisited, curr, curr.depth+1})
                }
            } 
        }
    }
    res := make([][]string, 0)
    for _, endNode := range endNodes {
        seq := make([]string, 0)
        for it := endNode; it != nil; it = it.parent {
            seq = append(seq, it.s)
        }
        res = append(res, Reverse(seq))
    }
    return res
}

func Reverse(s []string) []string {
    l := 0
    r := len(s)-1
    for l < r {
        s[l], s[r] = s[r], s[l]
        l++
        r--
    }
    return s
}

func DeepCopyMap(original map[string]bool) map[string]bool {
    c := map[string]bool{}
    for key, value := range original {
        c[key] = value
    }
    return c
}
type Item struct {
    s string
    visited map[string]bool
    parent *Item
    depth int
}

type Queue []*Item


func (q *Queue) Front() *Item {
    if len(*q) == 0 {
        return nil
    }
    return (*q)[0]
}

func (q *Queue) Pop() {
    if len(*q) > 0 {
        (*q) = (*q)[1:]
    }
}

func (q *Queue) Push(k *Item) {
    (*q) = append(*q, k)
}
