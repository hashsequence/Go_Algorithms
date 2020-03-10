/*
273. Integer to English Words
Hard


Convert a non-negative integer to its english words representation. Given input is guaranteed to be less than 231 - 1.

Example 1:

Input: 123
Output: "One Hundred Twenty Three"
Example 2:

Input: 12345
Output: "Twelve Thousand Three Hundred Forty Five"
Example 3:

Input: 1234567
Output: "One Million Two Hundred Thirty Four Thousand Five Hundred Sixty Seven"
Example 4:

Input: 1234567891
Output: "One Billion Two Hundred Thirty Four Million Five Hundred Sixty Seven Thousand Eight Hundred Ninety One"

solution:

make a dictionary for each single word

iterate through digit backwards and store it an array, we need to do this so we can access previous and next digit 

iterate array backwards so from the billions place to the ones places

_,_ _ _, _ _ _, _ _ _

notices for every triple we just append thousand, million, or billion after ever three set
and we only append the suffix if for every _ _ _ there is at least one nonzero digit

so we need to only translate _ _ _ and append a suffix


if place is between 1000 and 1000000 we divide place by 1000
if place is between 1000000 and 1000000000 we divid by 1000000

now we can evaluate _ _ _ so three cases for every digit:

ones place
tens place
hundreds place
*/

func numberToWords(num int) string {
    if num == 0 {
        return "Zero"
    }
    //create dictionary 
    dict := map[int]string {
        0:"",
        1:"One",
        2:"Two",
        3:"Three",
        4:"Four",
        5:"Five",
        6:"Six",
        7:"Seven",
        8:"Eight",
        9:"Nine",
        10:"Ten",
        11:"Eleven",
        12:"Twelve",
        13:"Thirteen",
        14:"Fourteen",
        15:"Fifteen",
        16:"Sixteen",
        17:"Seventeen",
        18:"Eighteen",
        19:"Nineteen",
        20:"Twenty",
        30:"Thirty",
        40:"Forty",
        50:"Fifty",
        60:"Sixty",
        70:"Seventy",
        80:"Eighty",
        90:"Ninety",
        100:"Hundred",
        1000:"Thousand",
        1000000:"Million",
        1000000000:"Billion",
    }
    number := ""
    place := 1
    arr := []int{}
    for i := 1 ; ; i*=10 {
        //getting digit at 
        digit := ((num % (i*10)) - (num % (i)))/(i)
        arr = append(arr, digit)
        place *= 10
        if i > num {
            //before breaking the loop we need to divide place by 10 to go back to the leftmost place of number
            place /= 10
            break
        }
    }
    prevDigit :=  0
    nextDigit :=  0
    countNonZeroes := 0
    for i := len(arr)-1; i>= 0; i-- {
        //get the nextDigit from left to right
        nextDigit = func() int {
            if i-1 >= 0 {
                return arr[i-1]
            }
            return 0
        }() 
        //get the prevDigit from left to right
        prevDigit = func() int {
            if i + 1 < len(arr) {
                return arr[i+1]
            }
            return 0
        }() 
        // we need to count the non zeroes
        if arr[i] != 0 {
            countNonZeroes++
        }
        part := Translate3Set(dict, arr[i], prevDigit, nextDigit, place)
        //we have an anonymous functions to add a space if the part is nonempty or number is nonempty
        number =  number + func() string {
            if part == "" || number == "" {
                return ""
            }
            return " "
        }() + part
        
        //if count of nonzeroes is greater than zero we add a suffix
        //since the bound is 2^31-1 we just need to account for billions
        //in case it is bounded by something higher than 
        //notice we only need to add a suffix if the number of zeroes in the place variable is multiple of 3 
        if (place == 1000) || place == 1000000 || place == 1000000000 {
            if countNonZeroes > 0 {
                 number = number + " " + dict[place]
            }
            //reset count to 0 for next 3set
            countNonZeroes = 0
        }
        //divide place by 10 to move to next place
        place /= 10
    }
    return number
}

func Translate3Set(dict map[int]string, digit int, prevDigit int,  nextDigit int, place int) string {
    if place >= 1000 && place < 1000000 {
        place /= 1000
    } else if place >= 1000000 && place < 1000000000 {
        place /= 1000000
    } else if place >= 1000000000 {
        place /= 1000000000
    }
    if place == 1 {
        if prevDigit != 1 {
            return dict[digit]
        }
    } else if place == 10 {
        if digit == 1 {
            return dict[nextDigit+10]
        } else if digit > 1 {
            return dict[digit*10]
        }
    } else if place == 100 {
        if digit != 0 {
            return dict[digit] + " " + dict[place]   
        }
    }
    return ""
}

