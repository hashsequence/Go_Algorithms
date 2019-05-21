func groupAnagrams(strs []string) [][]string {
    hashTable := map[string]*[]string{}
    for _, s := range strs {
        hashCode := hash(s)
        if hashTable[hashCode] == nil {
            arr := new([]string)
            hashTable[hashCode] = arr
        }
        *hashTable[hashCode] = append(*hashTable[hashCode],s)
    }
    
    res := make([][]string, len(hashTable))
    
    counter := 0
    for i, _ := range hashTable {
        res[counter] = *hashTable[i]
        counter++
    }
    return res
}

func hash(s string) string {
    countArr := [26]int{}
    
    for i, _ := range s {
        countArr[int(s[i])-int('a')]+=1
    }
    
    res := ""
    for i, _ := range countArr {
        res+= strconv.Itoa(countArr[i])+":" 
    }
    return res[:len(res)-1]
}
