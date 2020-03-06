/*
140. Word Break II
Hard

Given a non-empty string s and a dictionary wordDict containing a list of non-empty words, add spaces in s to construct a sentence where each word is a valid dictionary word. Return all such possible sentences.

Note:

The same word in the dictionary may be reused multiple times in the segmentation.
You may assume the dictionary does not contain duplicate words.
Example 1:

Input:
s = "catsanddog"
wordDict = ["cat", "cats", "and", "sand", "dog"]
Output:
[
  "cats and dog",
  "cat sand dog"
]
Example 2:

Input:
s = "pineapplepenapple"
wordDict = ["apple", "pen", "applepen", "pine", "pineapple"]
Output:
[
  "pine apple pen apple",
  "pineapple pen apple",
  "pine applepen apple"
]
Explanation: Note that you are allowed to reuse a dictionary word.
Example 3:

Input:
s = "catsandog"
wordDict = ["cats", "dog", "sand", "and", "cat"]
Output:
[]

solution:
brute force:

we make a map to have O(1) search in the wordDict
the brute force approach would be simple bfs at each point we find a matching word in dict 
as we iterate over the string

optimzation with dp:
how can we remember the function calls for wordbreak but starting at position i

the trick is we return imediately if dp[i] already exist

*/
func wordBreak(s string, wordDict []string) []string {
    dict := map[string]bool{}
    for i,_ := range wordDict {
        dict[wordDict[i]] = true
    }
   // return wordBreakHelperBruteForce(s, dict)
    return wordBreakHelper(s, 0, dict, make([]*[]string,len(s)+1))
}

func wordBreakHelper(s string, l int, dict map[string]bool, dp []*[]string) []string {
    if dp[l] != nil {
        return *dp[l]
    }
    res := make([]string,0)
    // this is the end case, all paths that leads to this is breakable, othewise we return an empty res so we cant
    //iterate over in subSentences since it would also be empty and never get added to results
    if l == len(s) {
        res = append(res, "")
    }
    for i := l+1; i <= len(s); i++ {
        if dict[s[l:i]] == true {
            subSentences := wordBreakHelper(s, i, dict, dp)
            for j, _ := range subSentences {
                // we need to check if the subsentence is an empty string so we don't add " "
                res = append(res, s[l:i]  + 
                             func() string {
                                 if subSentences[j] == "" {
                                     return "" 
                                 } else {
                                     return " " + subSentences[j] 
                                 }
                             }())
            }
        }
    }
    // store the results into dp so we dont need to go over function call process again
    dp[l] = &res
    return res
}


func wordBreakHelperBruteForce(s string, dict map[string]bool) []string {
    res := make([]string,0)
    for i := 0; i <= len(s); i++ {
        if dict[s[:i]] == true {
            if i == len(s) {
                res = append(res, s[:i])
            } else {
                variations := wordBreakHelperBruteForce(s[i:], dict)
                for j, _ := range variations {
                    res = append(res, s[:i] + " " + variations[j])
                }   
            }
        }
    }
    return res
}
