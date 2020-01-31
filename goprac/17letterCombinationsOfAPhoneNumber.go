/*
17. Letter Combinations of a Phone Number
Medium


Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.



Example:

Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
Note:

Although the above answer is in lexicographical order, your answer could be in any order you want.

solution:
superset type problem basically
start with base cases len = 0 and len = 1
initiaze arr for first digit 

double for loop iterate through digits and inner loop iterate through current combination arr
and for each string in the current combination array append each element out of the possible elements so if I see
2 then append a,b, and c to that current string and and put these new strings into the nextArr

*/
func letterCombinations(digits string) []string {
    mapDigit := map[byte]byte{'2':'a','3':'d','4':'g','5':'j','6':'m','7':'p','8':'t','9':'w'}
    if len(digits) == 0 {
        return []string{}
    } 
    
    var arr []string
    if digits[0] >= '2' && digits[0] <= '9' {
        arr = []string{string(mapDigit[digits[0]]), string(mapDigit[digits[0]] + 1) , string(mapDigit[digits[0]] + 2)}
        if digits[0] == '7' || digits[0] == '9' {
            arr = append(arr,string(mapDigit[digits[0]] + 3))
        }
    } 
    
    if len(digits) == 1 {
        return arr
    }
    
    for i := 1; i < len(digits); i++ {
        nextArr := make([]string,0)
        for j,_ := range arr {
             if digits[i] >= '2' && digits[i] <= '9' {
                 nextArr = append(nextArr, arr[j] + string(mapDigit[digits[i]]) , arr[j] + string(mapDigit[digits[i]] + 1) , arr[j] + string(mapDigit[digits[i]] + 2))
                 if digits[i] == '7' || digits[i] == '9' {
                     nextArr = append(nextArr, arr[j] + string(mapDigit[digits[i]] + 3))
                 }
              }
        }
        arr = nextArr
    }
    
    return arr
}

