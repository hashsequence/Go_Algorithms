func numDecodings(s string) int {  
    prev := 0
    prevPrev := 0
    count := 0
    for i,_ := range s {
        if i >= 2 {
            prevPrev = prev
            prev = count
            if s[i-1] != '2' && s[i-1] != '1' && s[i] == '0' {
                fmt.Println("edge case")
                return 0
            } 
            if s[i-1] == '0' {
                if s[i-2] != '1' && s[i-2] != '2' {
                    return 0
                }
                count = prev
            } else if s[i] == '0' && (s[i-1] == '1'|| s[i-1] == '2') {
                count = prevPrev
            } else if val, ok := strconv.Atoi(s[i-1:i+1]); ok == nil && val <= 26 {
                count = prev+prevPrev
            } 
        } else if i == 1 {
            prev = count
            if s[i-1] != '2' && s[i-1] != '1' && s[i] == '0' {
                fmt.Println("edge case")
                return 0
            } 
            if s[i] == '0' && (s[i-1] == '1'|| s[i-1] == '2') {
                count = prev
            } else if val, ok := strconv.Atoi(s[i-1:i+1]); ok == nil && val <= 26  {
                count = prev + 1
            }
        } else {
            if s[i] != '0' {
                count = 1
            } else {
                return 0
            }
        }
    }
    return count
}
