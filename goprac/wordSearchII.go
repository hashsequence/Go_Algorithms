func findWords(board [][]byte, words []string) []string {
//have a trie of all the words
    root := newTrie()
  //have a trie of all the words plus prefixes of each words
  //used to backtrack if prefix does not exist
    prefixes := newTrie()
    height := 0
    for i, _ := range words {
        insert(root, []byte(words[i]))
        for j := 0; j <= len([]byte(words[i])); j++ {
            insert(prefixes, []byte(words[i])[:j])
        }
        height = max(height,len(words[i]))
        
    }
    
    dir := [4]Dir{Dir{0,1}, Dir{1,0}, Dir{0,-1}, Dir{-1,0}}
    res := map[string]bool{}
    
    for i, _ := range board {
        for j, _ := range board[i] {
            str := board[i][j]
            board[i][j] = byte('#')
            findWordsDFS(&board, &dir, &res, root, prefixes, []byte{str}, height, len(words), Dir{i,j})
            board[i][j] = str
            if len(res) == len(words) {
                goto Exit
            }
            
        }
    }
    Exit:
    ret := make([]string, 0, len(words))
    for keys, _ := range res {
        ret = append(ret, keys)
    }
    return ret
}

func findWordsDFS(board *[][]byte, dir *[4]Dir, res *map[string]bool, root *TrieNode, prefixes *TrieNode, str []byte, height int, numWords int, coord Dir) {
    if !search(prefixes, []byte(str)) {
        return
    }
     fmt.Println(len(str), " ", string(str))
    if search(root, []byte(str)) {
        (*res)[string(str)] = true
     }
    if len(*res) == numWords || height == len(str){
        return
    }
    
    for k, _ := range (*dir) {
        newCoord := Dir{coord.x + (*dir)[k].x, coord.y + (*dir)[k].y}
        if  newCoord.x < 0 || newCoord.y < 0 || newCoord.x >= len((*board)) || newCoord.y >= len((*board)[0]) || (*board)[newCoord.x][newCoord.y] == '#' {
            continue
        }
        newStr := append(str,(*board)[newCoord.x][newCoord.y])
        temp := (*board)[newCoord.x][newCoord.y]
        (*board)[newCoord.x][newCoord.y] = byte('#')
        findWordsDFS(board, dir, res, root, prefixes, newStr, height, numWords, newCoord)
        (*board)[newCoord.x][newCoord.y] = temp
    }
    return 
}

type Dir struct {
    x int
    y int
}

type TrieNode struct {
    child []*TrieNode
    end bool 
}

func newTrie() *TrieNode {
    root := TrieNode{nil, false}
    root.child = make([]*TrieNode, 26)
    return &root
}

func insert(root *TrieNode, k []byte) {
    it := root
    for i := 0; i < len(k); i++ {
        index :=  k[i] - 97
        if it.child[index] == nil {
            it.child[index] = newTrie()
        }
        it = it.child[index]
    }
    it.end = true
}

func search(root *TrieNode, k []byte) bool {
    it := root
    for i := 0; i < len(k); i++ {
        index :=  k[i] - 97
        if it.child[index] == nil {
            return false
        }
        it = it.child[index]
    }
    return it != nil && it.end
}

func max(a int, b int) int {
	c := b
	if a > b {
		c = a
	}
	return c
}

