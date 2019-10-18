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

func myPowV2(base float64, n int) float64 {
    if n == 0 {
        return 1
    }
    if n == 1 {
        return base
    }
    if n == -1 {
        return 1/base
    }
    return myPowV2(base,func() int {
        if n % 2 == 0 {
        return n/2 
        } else {
            if n < 0 {
                 return n/2 - 1   
            } 
            return n/2 + 1   
        }
    }()) * myPowV2(base,n/2)
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
