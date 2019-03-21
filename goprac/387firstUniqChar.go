func firstUniqChar(s string) int {
    countArr := [26]int{}
    for i, _ := range s {
        countArr[s[i]-97]++
    }
    for i, _ := range s {
        if countArr[s[i]-97] == 1 {
          return i
        }
    }
    return -1
}

