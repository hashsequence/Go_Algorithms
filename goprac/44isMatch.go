/*
problem :
Given an input string (s) and a pattern (p), implement wildcard pattern matching with support for '?' and '*'.

'?' Matches any single character.
'*' Matches any sequence of characters (including the empty sequence).
The matching should cover the entire input string (not partial).

Note:

s could be empty and contains only lowercase letters a-z.
p could be empty and contains only lowercase letters a-z, and characters like ? or *.


idea: its dp 
base case check dp[0][j]
only case is the '*' since s := ""
so single for loop over j for i := 0
next have a double for loop i := 1 and j := 1 
if ?
check [i-1][j-1]
if * check
check [i-1][j-1] [i-1][j] [i][j-1]
*/

func isMatch(s string, p string) bool {
    dp := make([][]bool,len(s)+1)
    for i,_ := range dp {
        dp[i] = make([]bool,len(p)+1)
    }
    dp[0][0] = true
    for j := 1; j < len(p)+1; j++ {
        if j == 1 {
            if p[j-1] == '*' {
                dp[0][j] = true
            }  
        } else {
            if p[j-1] == '*' {
                dp[0][j] = dp[0][j-1]
            }
        }
    }
    
    for i := 1; i < len(s)+1; i++ {
        for j := 1; j < len(p)+1; j++ {
            if p[j-1] == '?' || s[i-1] == p[j-1] {
                dp[i][j] = dp[i-1][j-1]
            } else if p[j-1] == '*' {
                dp[i][j] = dp[i-1][j-1] || dp[i-1][j] || dp[i][j-1]
            }
        }
    }
    return dp[len(s)][len(p)]
}
