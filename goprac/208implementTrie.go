/*
208. Implement Trie (Prefix Tree)
Medium

2378

42

Add to List

Share
Implement a trie with insert, search, and startsWith methods.

Example:

Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // returns true
trie.search("app");     // returns false
trie.startsWith("app"); // returns true
trie.insert("app");   
trie.search("app");     // returns true
Note:

You may assume that all inputs are consist of lowercase letters a-z.
All inputs are guaranteed to be non-empty strings.

solution:

remember to make a TrieNode with

 child [26]*TrieNode
    isEnd bool
*/
type Trie struct {
    root *TrieNode
}

type TrieNode struct {
    child [26]*TrieNode
    isEnd bool
}

func NewTrieNode() *TrieNode {
    return &TrieNode{[26]*TrieNode{}, false}
}
/** Initialize your data structure here. */
func Constructor() Trie {
    return Trie{NewTrieNode()}
}


/** Inserts a word into the trie. */
func (this *Trie) Insert(word string)  {
    curr := this.root
    for i := 0; i < len(word); i ++ {
        letter := word[i] - 'a'
        if curr.child[letter] == nil {
            curr.child[letter] = NewTrieNode()
        } 
        curr = curr.child[letter]
    }
    curr.isEnd = true
}


/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
    curr := this.root
    for i := 0; i < len(word); i++ {
        letter := word[i] - 'a'
        if curr.child[letter] != nil {
            curr = curr.child[letter] 
        } else {
            return false
        }
    }
    return curr.isEnd
}


/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
    curr := this.root
    for i := 0; i < len(prefix); i++ {
        letter := prefix[i] - 'a'
        if curr.child[letter] != nil {
            curr = curr.child[letter]  
        } else {
            return false
        }
    }
    return true
}


/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
