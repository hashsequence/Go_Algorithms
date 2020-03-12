/*
472. Concatenated Words
Hard

Given a list of words (without duplicates), please write a program that returns all concatenated words in the given list of words.
A concatenated word is defined as a string that is comprised entirely of at least two shorter words in the given array.

Example:
Input: ["cat","cats","catsdogcats","dog","dogcatsdog","hippopotamuses","rat","ratcatdogcat"]

Output: ["catsdogcats","dogcatsdog","ratcatdogcat"]

Explanation: "catsdogcats" can be concatenated by "cats", "dog" and "cats"; 
 "dogcatsdog" can be concatenated by "dog", "cats" and "dog"; 
"ratcatdogcat" can be concatenated by "rat", "cat", "dog" and "cat".
Note:
The number of elements of the given array will not exceed 10,000
The length sum of elements in the given array will not exceed 600,000.
All the input string will only include lower case letters.
The returned elements order does not matter.

solution:

I can check if each word is breakable from the words in the list

time complexity O(k*N^2) where N is the length of the word with the greatest length and k is the number of words in the wordlist
since we need to check if each word is breakable which is O(N^2)
*/

func findAllConcatenatedWordsInADict(words []string) []string {
    dict := map[string]bool{}

    for i,_ := range words {
        dict[words[i]] = true
    }
    
    res := make([]string, 0)
    
    for i,_ := range words {
        
        dict[words[i]] = false
        dp := make([]bool, len(words[i])+1)
        if WordBreak(words[i], dict, dp) {
            res = append(res, words[i])
        }
        
        dict[words[i]] = true
        
    }
    
    return res
}

func WordBreak(word string, dict map[string]bool, dp []bool) bool {
    if word == "" {
        return false
    }
    dp[0] = true
    //dp[i] means word[0:i] is breakable
    //iterate over substrings word[0:i]
    for i := 0; i <= len(word); i++ {
        //check if substring word[0:i] is breakable
        for j := 0; j < i; j++ {
            if dp[j] && dict[word[j:i]] {
                dp[i] = true
                break
            }
        }
    }
    return dp[len(word)]
}
