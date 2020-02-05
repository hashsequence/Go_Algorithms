/*
69. Sqrt(x)
Easy


Implement int sqrt(int x).

Compute and return the square root of x, where x is guaranteed to be a non-negative integer.

Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.

Example 1:

Input: 4
Output: 2
Example 2:

Input: 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since 
             the decimal part is truncated, 2 is returned.
             
 solution :
 binary search
 remember l = med + 1
 r = med - 1 to avoid infinite loops with even numbers
 med = l + (r-l)/2
 if does not find exact return lower bound and since condition of loop ir l <= r
 r becomes lowerbound in the end so return r

*/
func mySqrt(x int) int {
    return Sqrt(x)
    //return squareRootHelper(x,1,x)  //recursive version
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


func squareRootHelper(a, left, right int) int {
 
  var med int
  med = (right+left)/2 //5 //3
  fmt.Println("a: ", a, " left: ", left, "right: ", right, "mid: ", med)
  
    if left > right {
    if left*left == a {
      return left
    } else {
      return right 
    }
  }
  

  if med*med == a {
    return med
  } else if med*med < a {
    return squareRootHelper(a,med+1,right) // 49, 5, 11 // 49, 3, 10
  } else {
    return squareRootHelper(a,left,med-1)// 49, 1, 24 // 49, 2, 12
  }
 
}
