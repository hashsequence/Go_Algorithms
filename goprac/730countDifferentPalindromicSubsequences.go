/*
730. Count Different Palindromic Subsequences
Hard

406

38

Add to List

Share
Given a string S, find the number of different non-empty palindromic subsequences in S, and return that number modulo 10^9 + 7.

A subsequence of a string S is obtained by deleting 0 or more characters from S.

A sequence is palindromic if it is equal to the sequence reversed.

Two sequences A_1, A_2, ... and B_1, B_2, ... are different if there is some i for which A_i != B_i.

Example 1:
Input: 
S = 'bccb'
Output: 6
Explanation: 
The 6 different non-empty palindromic subsequences are 'b', 'c', 'bb', 'cc', 'bcb', 'bccb'.
Note that 'bcb' is counted only once, even though it occurs twice.
Example 2:
Input: 
S = 'abcdabcdabcdabcdabcdabcdabcdabcddcbadcbadcbadcbadcbadcbadcbadcba'
Output: 104860361
Explanation: 
There are 3104860382 different non-empty palindromic subsequences, which is 104860361 modulo 10^9 + 7.
Note:

The length of S will be in the range [1, 1000].
Each character S[i] will be in the set {'a', 'b', 'c', 'd'}.

solution : 
use recusive top to bottom approach
let dp = number of diff palin subseq in S[i:j]
init all dp[i][j] = -1 to tell if it was set
cases:

1) i == j then dp[i][j] = 1
2) j - i == 1 then 2 subseq because ex. aa {a, aa} ab {a, b}
3) if dp[i][j] != -1 then value was set and so return
4) actual case:
if S[i] != [j]
then we know dp[i][j] is equal to the number of subsequence in dp[i+1][j-1] and the number of subsequence in dp[i+1][j-1] again with S[i]<subseq>S[j] we then subtract that number by dp[i+1][j-1] since we double counted\

if S[i] == S[j] there are three more cases:
let S[i] = S[j] = 'a'

so we have 
'a'<subseq[i+1:j-1] with no 'a' in it>'a' ====> which gives us dp[i+1][j-1] * 2, b/c one for original and again + 'a' appended to both ends of each subsequence in the original + 2 since we have 'a''a' and 'a' added
'a'<subseq[i+1:j-1] with one 'a' in the middle>'a' ===> which gives us dp[i+1][j-1] * 2, b/c one for original and again + 'a' appended to both ends of each subsequence in the original and + 1 since we only have 'aa' extra
'a'<>subseq[i+1:j-1] with 'a'<one or more characters>'a' in the middle> a'this is the most difficult, basically we have  + original with 'a' attached to each end, but we are double counting the ones in'a'<one or more characters>'a' 

*/
func countPalindromicSubsequences(S string) int {
    dp := make([][]int,len(S))
    for i, _ := range dp {
        dp[i] = make([]int, len(S))
    }
    for i, _:= range dp {
        for j,_ := range dp[i] {
            dp[i][j] = -1
        }
    }
    return countPalindromicSubsequencesHelper(S, dp ,0,len(S)-1)
}

func countPalindromicSubsequencesHelper(S string, dp [][]int, i, j int) int {
    if i < 0 || j >= len(S) || j < i  {
        return 0
    }
    if i == j {
        dp[i][j] = 1
        return 1
    }
    if j-i == 1 {
        dp[i][j] = 2
        return 2
    }
    if dp[i][j] != -1 {
        return dp[i][j] 
    }
    if S[i] == S[j] {
        left := i+1
        right := j-1
        currLetter := S[i]
        for left <= right && S[left] != currLetter {
            left++
        } 
        for left <= right && S[right] != currLetter {
            right--
        }
        if left > right {
            dp[i][j] = (countPalindromicSubsequencesHelper(S, dp, i + 1, j - 1) * 2) + 2
        } else if left == right {
            dp[i][j] = (countPalindromicSubsequencesHelper(S, dp, i + 1, j - 1) * 2) + 1
        } else {
            dp[i][j] = (countPalindromicSubsequencesHelper(S, dp, i + 1, j - 1) * 2) - (countPalindromicSubsequencesHelper(S, dp, left + 1, right - 1))
        }
    } else {
        dp[i][j] = (countPalindromicSubsequencesHelper(S, dp, i + 1, j)) + (countPalindromicSubsequencesHelper(S, dp, i, j - 1))  - (countPalindromicSubsequencesHelper(S, dp, i + 1, j - 1))
    }
    dp[i][j] = (dp[i][j] +  1000000007)  % 1000000007       
    return dp[i][j]
}
