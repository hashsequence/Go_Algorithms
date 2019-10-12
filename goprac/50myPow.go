func myPow(x float64, n int) float64 {
    var acc float64 = 1
    if n == 0 {
        return 1
    }
    acc = multiply(x,abs(n))
    if n < 0 {
        acc = 1/acc
    } 
    return acc
}

func abs(n int) int{
    if n < 0 {
        return n * -1
    } 
    return n
}

func multiply(x float64, n int) float64 {
    if n == 1 {
        return x
    }
    v := multiply(x,n/2) 
    if n % 2 == 0 {
        v *=v
    } else {
        v = v * v * x
    }
    return v
}
