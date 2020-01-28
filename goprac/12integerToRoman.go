/*
12. Integer To Roman
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9. 
X can be placed before L (50) and C (100) to make 40 and 90. 
C can be placed before D (500) and M (1000) to make 400 and 900.
Given an integer, convert it to a roman numeral. Input is guaranteed to be within the range from 1 to 3999.

idea:
must know how to iterate through the number
must know how to set up if else cases for each digit as you iterate
use mapping table to map numbers to roman numeral
*/
func intToRoman(num int) string {
    intToSymbol :=  map[int]string{1:"I", 5:"V", 10:"X", 50:"L", 100:"C", 500:"D", 1000:"M"}
    romanStr := ""
    
    for i := 1;num / i != 0 ; i*=10 {
        currNum := (num-(num % i)) - (num-(num % (i*10)))
        currPlace := i
        currDigit := currNum / currPlace
       // fmt.Println(currDigit, currPlace)
        if currDigit == 9{
            romanStr = intToSymbol[currPlace] + intToSymbol[currPlace*10] + romanStr
        } else if currDigit == 4 {
            romanStr = intToSymbol[currPlace] + intToSymbol[currPlace*5] + romanStr
        }else if currDigit > 5 && currDigit < 9 {
            romanStr = intToSymbol[5 * currPlace] + valueRepeatedXTimes(intToSymbol[currPlace],currDigit-5) + romanStr
        } else if currDigit < 5 && currDigit > 1 {
            romanStr = valueRepeatedXTimes(intToSymbol[currPlace],currDigit) + romanStr
        } else if currDigit == 1 || currDigit == 5 {
            romanStr = intToSymbol[currNum] + romanStr
        } 
    }
    return romanStr
}

func valueRepeatedXTimes(value string, x int) string {
    str := ""   
    for i := 0; i < x; i++ {
        str += value
    }
    return str
}
