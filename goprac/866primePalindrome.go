/*
866. Prime Palindrome
Medium

Find the smallest prime palindrome greater than or equal to N.

Recall that a number is prime if it's only divisors are 1 and itself, and it is greater than 1. 

For example, 2,3,5,7,11 and 13 are primes.

Recall that a number is a palindrome if it reads the same from left to right as it does from right to left. 

For example, 12321 is a palindrome.

 

Example 1:

Input: 6
Output: 7
Example 2:

Input: 8
Output: 11
Example 3:

Input: 13
Output: 101
 

Note:

1 <= N <= 10^8
The answer is guaranteed to exist and be less than 2 * 10^8.

solution:
brute force + math shortcut:
make isPrime and isPalindrome and pow and numDigits function
must notice all even digit palindromes cannot be prime
and iterate until 2 * 10^8




*/
func primePalindrome(N int) int {
    if N <= 1 {
        return 2
    }
    if N <= 10 && N > 7 {
        return 11
    }
    i := N
    for ; i < 200000000;  {
        if isPalindrome(i) {
            if isPrime(i) {
                return i
            }
        }
        num := NumDigits(i)
        if num % 2 == 0  {
            i = Pow(10,num)
        } else {
            i++
        }
    }
    return i
}

func Pow(base int, n int) int {
    if n == 0 {
        return 1
    }
    if n == 1 {
        return base
    }
    if n == -1 {
        return 1/base
    }
    return Pow(base,func() int {
        if n % 2 == 0 {
        return n/2 
        } else {
            if n < 0 {
                 return n/2 - 1   
            } 
            return n/2 + 1   
        }
    }()) * Pow(base,n/2)
}

func NumDigits(N int) int {
    if N < 10 {
        return 1
    }
    count := 0
    i := 1
    for i <=  N  {
        count++
        i*=10
    }
    return count
}

func isPalindrome(N int) bool {
    if N < 10 {
        return true
    }
    right := 1
    for right <  N  {
        right*= 10
    }
    left := 10
    
    for right > 1 && left >  1{
        if right == 1 || left == 1 {
            break
        }
        leftDigit := ((N%left) - (N%(left/10)))/(left/10)
        rightDigit := ((N%right) - (N%(right/10)))/(right/10)
        if leftDigit != rightDigit {
            return false
        }
        right /= 10
        left *= 10
    }
    return true
}

func Sqrt(N int) int {
    if N == 1 {
        return 1
    }
    if N == 0 {
        return 0
    }
    l := 0
    r := N-1
    x := l + (r-l)/2
    
    for l <= r {
       // fmt.Println(l, r, x)
        if x * x == N {
            return x
        }
        if x * x > N {
            r = x - 1 
        } else if x * x < N {
            l = x + 1 
        }
        x = l + (r-l)/2
    }
    //r becomes lower than l so return lower bound
    return r 
    
}

func isPrime(N int) bool {
    if N <= 1 {
        return false
    } 
    bound := Sqrt(N)
    for i := 2; i<=bound; i++ {
        if N%i == 0 {
            return false
        }
    }  
    return true; 
}

