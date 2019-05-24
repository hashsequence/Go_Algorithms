func mySqrt(x int) int {
  if x == 0 {
    return 0
  }
  return squareRootHelper(x,1,x) 
}

func squareRootHelper(a, left, right int) int {
 
  var med int
  med = (right+left)/2 //5 //3
 // fmt.Println("a: ", a, " left: ", left, "right: ", right, "mid: ", med)
  
    if left > right {
    if left*left == a {
      return left
    } else {
      return right 
    }
  }
  
  //if right > left {
    if med*med == a {
      return med
    } else if med*med < a {
      return squareRootHelper(a,med+1,right) // 49, 5, 11 // 49, 3, 10
    } else {
      return squareRootHelper(a,left,med-1)// 49, 1, 24 // 49, 2, 12
    }
  //} else {
    //if med*med < a {
      //return squareRootHelper(a,left+1, right+1) 
    //} else {

  
 
 
}
