/*
1044. Longest Duplicate Substring
Hard

Given a string S, consider all duplicated substrings: (contiguous) substrings of S that occur 2 or more times.  (The occurrences may overlap.)

Return any duplicated substring that has the longest possible length.  (If S does not have a duplicated substring, the answer is "".)

 

Example 1:

Input: "banana"
Output: "ana"
Example 2:

Input: "abcd"
Output: ""
 

Note:

2 <= S.length <= 10^5
S consists of lowercase English letters.


solution:

brute force is iterate over every substring (non-unique)
so N^2 substrings possible 
we can use map too keep count of each substring
we can keep max score everytime

optimized:

we keep the hashset except our sub tasks is finding every sustring of size i between 1 and n 

we need a binary search type method 


to make it faster we need a better hash

like rolling hash
also dependent on language use a mutable array
*/
func longestDupSubstring(S string) string {
    freq := map[string]int{}
    l := 0
    r := len(S)
    pivot := (r+l)/2
    dup := false
    dupStr := ""
    for l <= r {
        //fmt.Println(pivot)
        str := Search(S, pivot, freq)
        if str != "" {
           dup = true
           dupStr = str
        } else {
            dup = false
        }
        if !dup {
            r = pivot-1
        } else {
            l = pivot+1
        }
        pivot = (r+l)/2
    }
    return dupStr
    
}

func Search(S string, windowSize int, freq map[string]int) string {
    for i := 0; i <= len(S)-windowSize; i++ {
        word := S[i:i+windowSize]
        //fmt.Println(word)
        if _, ok := freq[word]; !ok {
            freq[word] = 1
        } else {
            return word
        }
    }
    return ""
}
