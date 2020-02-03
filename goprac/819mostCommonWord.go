/*
819. Most Common Word
Easy

Given a paragraph and a list of banned words, return the most frequent word that is not in the list of banned words.  It is guaranteed there is at least one word that isn't banned, and that the answer is unique.

Words in the list of banned words are given in lowercase, and free of punctuation.  Words in the paragraph are not case sensitive.  The answer is in lowercase.

 

Example:

Input: 
paragraph = "Bob hit a ball, the hit BALL flew far after it was hit."
banned = ["hit"]
Output: "ball"
Explanation: 
"hit" occurs 3 times, but it is a banned word.
"ball" occurs twice (and no other word does), so it is the most frequent non-banned word in the paragraph. 
Note that words in the paragraph are not case sensitive,
that punctuation is ignored (even if adjacent to words, such as "ball,"), 
and that "hit" isn't the answer even though it occurs more because it is banned.
 

Note:

1 <= paragraph.length <= 1000.
0 <= banned.length <= 100.
1 <= banned[i].length <= 10.
The answer is unique, and written in lowercase (even if its occurrences in paragraph may have uppercase symbols, and even if it is a proper noun.)
paragraph only consists of letters, spaces, or the punctuation symbols !?',;.
There are no hyphens or hyphenated words.
Words only consist of letters, never apostrophes or other punctuation symbols.

solution: 
make a banned map
make a frequency map
use a Lower function to have all the words in paragraph lowercase 
have isLetter function to distinguish letter characters 
iterate over paragraphs parsing out each words (this is the tedious part)

HAVE an if statent to check last word after loop ends
check the last word if paragraph ends with a word with no punctuation or if the paragraph was only one word
*/
func mostCommonWord(paragraph string, banned []string) string {

    bannedMap := map[string]bool{}
    frequencyMap := map[string]int{}
    
    for i, _ := range banned {
        bannedMap[banned[i]] = true
    }
    
    startOfWord:= 0    
    //regular case
    for i, _ := range paragraph {
        if isLetter(paragraph[i]) {
            if i > 0 && !isLetter(paragraph[i-1]) {
                startOfWord = i
            }
        } else {
            if i > 0 && i != startOfWord && isLetter(paragraph[i-1]) {
                //fmt.Println(Lower(paragraph[startOfWord:i]))
                if val, _ := bannedMap[Lower(paragraph[startOfWord:i])]; !val {
                    if _,ok := frequencyMap[Lower(paragraph[startOfWord:i])]; !ok {
                        frequencyMap[Lower(paragraph[startOfWord:i])] = 1
                    } else {
                        frequencyMap[Lower(paragraph[startOfWord:i])]++
                    }
                }
            } 
        }
    }
    //check the last word if paragraph ends with a word with no punctuation or if the paragraph was only one word
    if startOfWord >= 0 && len(paragraph) > 1 && isLetter(paragraph[len(paragraph)-1]) {
        //fmt.Println("checking last word")
        if val, _ := bannedMap[Lower(paragraph[startOfWord:len(paragraph)])]; !val {
            if _,ok := frequencyMap[Lower(paragraph[startOfWord:len(paragraph)])]; !ok {
                frequencyMap[Lower(paragraph[startOfWord:len(paragraph)])] = 1
            } else {
                frequencyMap[Lower(paragraph[startOfWord:len(paragraph)])]++
            }
        }
    }
    //fmt.Println(frequencyMap)
    return Max(frequencyMap)
}

func Lower(s string) string {
    temp := make([]byte, len(s))
    copy(temp,[]byte(s))
    for i,_ := range temp {
        if temp[i] >= 'A' && temp[i] <= 'Z' {
            temp[i] += ('a' - 'A')
        }
    }
    return string(temp)
}

func isLetter(c byte) bool {
    return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func Max(a map[string]int) string {
    c := ""
    for i, _ := range a {
        if a[i] >= a[c] {
            c = i
        }
    }
    return c
   
}
