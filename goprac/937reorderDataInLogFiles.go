/*
937. Reorder Data in Log Files
Easy

You have an array of logs.  Each log is a space delimited string of words.

For each log, the first word in each log is an alphanumeric identifier.  Then, either:

Each word after the identifier will consist only of lowercase letters, or;
Each word after the identifier will consist only of digits.
We will call these two varieties of logs letter-logs and digit-logs.  It is guaranteed that each log has at least one word after its identifier.

Reorder the logs so that all of the letter-logs come before any digit-log.  The letter-logs are ordered lexicographically ignoring identifier, with the identifier used in case of ties.  The digit-logs should be put in their original order.

Return the final order of the logs.

solution:
need to implement a modified sort based on conditions to sort letter logs
sift down digits logs to the bottom, and you can preserver order of digit logs by iterating logs array backwards
*/

func reorderLogFiles(logs []string) []string {
    //sift down digits logs to bottom by iterating through logs array backwards to preserver order order of digits
    endOfLetterLogs := len(logs)-1
    for i := len(logs)-1; i >= 0; i-- {
        if doesStrContainOnlyDigits(logs[i][IndexOf(logs[i], ' ',1)+1:]) {
            Swap(&logs[i],&logs[endOfLetterLogs])
            endOfLetterLogs--
        }
    }
    //sort the letter logs partition lexographically
    qSort(logs[:endOfLetterLogs+1])
    return logs
    
}

func NthWord(s string, n int) string {
    fw := ""
    j := 0
    start := 0
    for i,_ := range s {
        if s[i] == ' ' {
            j++
            if j == n {
                fw = s[start:i]
                break
            }
            start = i+1
        }
    }
    return fw
}

func IndexOf(s string, delimeter byte, n int) int {
    j := 0
    for i, _ := range s {
        if s[i] == delimeter {
            j++
            if j == n {
                return i
            }
        }
    }
    return -1
}

func Swap(a, b *string) {
    *a, *b = *b, *a
}

func Reverse(strs []string) {
    for l, r := 0, len(strs)-1; l < r; l, r = l+1,r-1 {
        Swap(&strs[l], &strs[r])
    }
}

func qSort(A []string) {
    if len(A) <= 1 {
        return 
    }
    
    left := 0
    right := len(A)-1
    pivot := (len(A)-1)/2
    Swap(&A[right],&A[pivot])
    for i, _ := range A {
        if A[i][IndexOf(A[i],' ', 1)+1:] < A[right][IndexOf(A[right],' ', 1)+1:] {
            Swap(&A[i], &A[left])
            left++
        } else if A[i][IndexOf(A[i],' ', 1)+1:] == A[right][IndexOf(A[right],' ', 1)+1:] {
            if  NthWord(A[i],1) < NthWord(A[right],1) {
                 Swap(&A[i], &A[left])
                 left++
            }
        }
    }
    Swap(&A[left], &A[right])
    qSort(A[:left])
    qSort(A[left+1:])
}

func doesStrContainOnlyDigits(s string) bool {
    b := true
    for _, c := range s {
        if c != ' ' && (c < '0' || c > '9') {
	    	b = false
	    	break
	    }
    }
    return b
}
