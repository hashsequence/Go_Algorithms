/*
14. Longest Common Prefix

Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".
*/
func longestCommonPrefix(strs []string) string {
    longestCommonPrefix := ""
    //base cases
    if len(strs) == 1 {
        return strs[0]
    } else if len(strs) == 0 {
        return ""
    }
    
    done := false
    //idea" check each character at the same position and check if that character changes
    //if doesnt change after check each character for each string then append to longestCommonPrefix
    for j := 0; done != true; j++ {
        init := false
        var character byte
        same := true
        for i,_ := range strs {
            if j >= len(strs[i]) {
                done = true
                break
            }
            if init == false {
                character = strs[i][j]
                init = true
            } else {
                if strs[i][j] != character {
                    same = false
                    done = true
                    break
                }
            }
        }
        if same && !done {
            longestCommonPrefix += string(character)
        }
    }
    return longestCommonPrefix
  
}
