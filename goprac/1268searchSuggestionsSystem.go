/*
1268. Search Suggestions System
Medium

161

58

Add to List

Share
Given an array of strings products and a string searchWord. We want to design a system that suggests at most three product names from products after each character of searchWord is typed. Suggested products should have common prefix with the searchWord. If there are more than three products with a common prefix return the three lexicographically minimums products.

Return list of lists of the suggested products after each character of searchWord is typed. 

 

Example 1:

Input: products = ["mobile","mouse","moneypot","monitor","mousepad"], searchWord = "mouse"
Output: [
["mobile","moneypot","monitor"],
["mobile","moneypot","monitor"],
["mouse","mousepad"],
["mouse","mousepad"],
["mouse","mousepad"]
]
Explanation: products sorted lexicographically = ["mobile","moneypot","monitor","mouse","mousepad"]
After typing m and mo all products match and we show user ["mobile","moneypot","monitor"]
After typing mou, mous and mouse the system suggests ["mouse","mousepad"]
Example 2:

Input: products = ["havana"], searchWord = "havana"
Output: [["havana"],["havana"],["havana"],["havana"],["havana"],["havana"]]
Example 3:

Input: products = ["bags","baggage","banner","box","cloths"], searchWord = "bags"
Output: [["baggage","bags","banner"],["baggage","bags","banner"],["baggage","bags"],["bags"]]
Example 4:

Input: products = ["havana"], searchWord = "tatiana"
Output: [[],[],[],[],[],[],[]]
 

Constraints:

1 <= products.length <= 1000
There are no repeated elements in products.
1 <= Σ products[i].length <= 2 * 10^4
All characters of products[i] are lower-case English letters.
1 <= searchWord.length <= 1000
All characters of searchWord are lower-case English letters.

solution:
sort product list first
iterate over searchwords prefixes
ex. mouse
1st iteration: m
2nd: mo
3rd: mou
...etc...
if we see a match put it into the set for the current prefix
append top 3 to the result set
set the next set to be the pruned current set


*/
func suggestedProducts(products []string, searchWord string) [][]string {
    
    QSort(products)
    res := make([][]string, 0)
    currSet := make([]string, 0)
    for i, _ := range searchWord {
        if i == 0 {
            for j, _ := range products {
                if i < len(products[j]) && products[j][0] == searchWord[0] {
                    currSet = append(currSet, products[j])
                }
            }
            if len(currSet) > 3 {
                res = append(res, currSet[:3])
            } else {
                res = append(res, currSet)
            }
        } else {
            nSet := make([]string, 0)
            for j,_ := range currSet {
                if  i < len(currSet[j]) && currSet[j][i] == searchWord[i] {
                    nSet = append(nSet, currSet[j]) 
                }
            }
            if len(nSet) > 3 {
                res = append(res, nSet[:3])
            } else {
                res = append(res, nSet)
            }
            currSet = nSet
            
        }
    }
    return res
}

func QSort(A []string) {
    if len(A) < 2 {
        return 
    }
    left := 0
    right := len(A) -  1
    pivot := left + (right - left)/2 // rand.Int() % len(A)
    
    //move pivot value to end
    Swap(&A[right], &A[pivot])
    
    for i := 0; i < right; i++ {
        if A[i] < A[right] {
            Swap(&A[i], &A[left])
            left++
        }
    }
    //put back pivot value right after last smallest element
    Swap(&A[left], &A[right])
    QSort(A[:left])
    QSort(A[left+1:])
}

func Swap(a, b *string) {
    *a, *b = *b, *a
}


