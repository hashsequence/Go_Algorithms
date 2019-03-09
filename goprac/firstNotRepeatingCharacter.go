func firstNotRepeatingCharacter(s string) string {
    bitArr := [26]Letter{}
    for i, val := range s {
        bitArr[mapRuneToInt(val)].count++
        if  bitArr[mapRuneToInt(val)].count == 1 {
             bitArr[mapRuneToInt(val)].index = i
        }
    }
    
    minIndex := len(s)
    minLetter := '_'
    for i, val := range bitArr {
        if val.count == 1 && val.index < minIndex{
            minIndex = val.index
            minLetter =  mapIntToRune(i)    
        }
    }
    return  string(minLetter)
}

type Letter struct{
    count int
    index int
}

func mapRuneToInt(l rune) int {
    return int(l)-97
}

func mapIntToRune(i int) rune {
    return rune(i+97)
}
