func minWindow(s string, t string) string {
    
   //yes you can index string, but it will give you the ascii number
    //initialize count map for t 
    hash_t := [128]int{}
     for _, val := range t {
        hash_t[val]++
    }
 
    //initialize count map for s
    hash_s := [128]int{}
    if len(s) < len(t) {
        return ""
    }
    l := 0
    r := 0
    finalL := -1
    finalR := len(s) + 1
    count := 0
    
    for ; r < len(s); r++ {
        //fmt.Println(r)
        hash_s[s[r]]++
        if hash_t[s[r]] > 0 && hash_s[s[r]] <= hash_t[s[r]] {
            count++
        }
        
        if count == len(t) {
           // fmt.Println("count r is ", r, " and l is ", l)
            //get rid of useless ones and duplicates 
           
            for hash_t[s[l]]== 0 || hash_s[s[l]] > hash_t[s[l]] {
             
                if hash_s[s[l]] > hash_t[s[l]] {
                    hash_s[s[l]]--
                }
                l++
            }
            if r-l < finalR-finalL {
                finalR = r
                finalL = l
            }
        }
    }
    if finalL == -1 {
        return ""
    }
    return s[finalL:finalR+1]
 
}
