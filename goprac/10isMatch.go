func isMatch(s string, p string) bool {
    dp := make([][]bool,len(s)+1)
    for i,_ := range dp {
        dp[i] = make([]bool,len(p)+1)
    }
    dp[0][0] = true
    
    //case when "" 
    for j := 1; j < len(p) + 1; j++ {
        //obviously at j-1 theres no j-2 so we have this case so it doesnt get an index error
        if j == 1  && p[j-1] == '*'{
            dp[0][j] = true
            //p[j-1] is "*" but only true if p[j-2] is also "*"
        } else if j > 1 && p[j-1] == '*' && dp[0][j-2] {
            dp[0][j] = true
        }
    }
    
    for i := 1; i < len(s)+1; i++ {
        for j:= 1; j < len(p)+1; j++ {
            // s[i-1] := "<curr>" and p[j-1] := "<curr>" or "."
            if s[i-1]==p[j-1] || p[j-1] == '.' {
                //then same as previous case without s[i-1] and without p[j-1]
                dp[i][j] = dp[i-1][j-1]
              //case when p[j-1] = "*"
            } else if p[j-1] == '*' {
                //if the character at s[i-1] matches p[j-2] so for example s := "<curr>" and we have p := "<curr>*" or ".*"
                if s[i-1] == p[j-2] || p[j-2] == '.' {
                    //three cases its the same as and s[i-1] = "<curr>": 
                    //1) so p[j-2:j] := "<curr>*"
                    //2) so p[j-2:j] := ".*"
                    //3) so p[j-3:j] := "<curr>.*"
                    dp[i][j] = dp[i-1][j] || dp[i][j-1] || dp[i][j-2]
                } else {
                    //p[j-1] is "*" but p[j-2] does not match s[i-1] but we can skip p[i-1] since its "*" since "*" is zero or more of preceding char
                    dp[i][j] = dp[i][j-2]
                }
            }
        }
    }
    return dp[len(s)][len(p)]
}
