func numDecodings(s string) int {
    if len(s) == 0 {
        return 0
    }
    curr := singles(s[0])
    prev := 1
    
    for i := 1; i < len(s); i++ {
        currCount := curr * singles(s[i]) + prev * doubles(s[i-1],s[i])
        prev = curr
        curr = currCount % (int(math.Pow(10,9)) + 7)
    }
    return curr 
}

func singles(first byte) int {
    if first == '*' {
        return 9
    } else if first == '0'{
        return 0
    } else {
        return 1
    }
}

func doubles(first, second byte) int {
    if first == '*' {
        if second == '*' {
            return 15
        } else if second >= '0' && second <= '6' {
            return 2
        } else if second >= '7' && second <= '9' {
            return 1
        }
    } 
    if first == '1' {
        if second == '*' {
            return 9
        } else if second >= '0' && second <= '9' {
            return 1
        }
    } 
    if first == '2' {
        if second == '*' {
            return 6
        } else if second >= '0' && second <= '6' {
            return 1
        }  else if second >= '7' && second <= '9' {
            return 0
        }
    }
    return 0
}
