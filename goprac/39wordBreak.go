/*
39. Word Break
Medium


Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, determine if s can be segmented into a space-separated sequence of one or more dictionary words.

Note:

The same word in the dictionary may be reused multiple times in the segmentation.
You may assume the dictionary does not contain duplicate words.
Example 1:

Input: s = "leetcode", wordDict = ["leet", "code"]
Output: true
Explanation: Return true because "leetcode" can be segmented as "leet code".
Example 2:

Input: s = "applepenapple", wordDict = ["apple", "pen"]
Output: true
Explanation: Return true because "applepenapple" can be segmented as "apple pen apple".
             Note that you are allowed to reuse a dictionary word.
Example 3:

Input: s = "catsandog", wordDict = ["cats", "dog", "sand", "and", "cat"]
Output: false
Accepted

solution:

use a trie or a map to store the words
let dp[i] remember if s[0:i] is breakable


*/
func wordBreak(s string, wordDict []string) bool {
    dict := NewTrie()
    for i, _ := range wordDict {
        dict.Insert(wordDict[i])
    }
    //recursive way with true
    /*
    return WordBreakHelper(s, &dict)
    */
    
    //dp[i] means that s[0:i] is breakable, note "" is breakable by base case
    dp := make([]bool, len(s)+1)
    dp[0] = true
    for i := 1; i <= len(s); i++ {
        for j := 0; j < i; j++ {
            if dp[j] && dict.Search(s[j:i]) {
                dp[i] = true
                break;
            }
        }
    }
    return dp[len(s)]

}

func WordBreakHelper(s string, dict *Trie ) bool {
    if s == "" {
        return true
    }
     for i := 0; i <= len(s); i++ {
         if dict.Search(s[0:i]) && WordBreakHelper(s[i:], dict) {
             return true
         }
    }
    return false
}

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

func NewTrie() Trie {
    return Trie{NewTrieNode()}
}

func (this *Trie) Insert(s string) {
    curr := this.root
    for i,_ := range s {
        letter := s[i] - 'a'
        if curr.child[letter] == nil {
            curr.child[letter] = NewTrieNode()
        }
        curr = curr.child[letter]
    }
    curr.isEnd = true
}

func (this *Trie) Search(s string) bool {
    if s == "" {
        return true
    }
    curr := this.root
    for i,_ := range s {
        letter := s[i] - 'a'
        if curr.child[letter] == nil {
            return false
        }
        curr = curr.child[letter]
    }
    if curr != nil && curr.isEnd == true {
        return true
    }
    return false
}
