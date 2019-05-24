/*
run a for loop such that we start with x and successively divide by 10 to iterate through each place
hav an accummulator where we add the current x mod 10 to (get the ones place) to the current accumulator (the previous value * 10 since we need to shift it 
right)


*/
func reverse(x int) int {
    res := 0
    for x != 0 {
        res =  res* 10 + x%10
        x /= 10
    }
    fmt.Println(bInt(-1<<31 ),-1<<31)
    fmt.Println(bInt(1 << 31 - 1),1 << 31 - 1)
     
    if res > (-1<<31) && res < (1 << 31 - 1) { 
        return res
    }
    return 0
}

func bInt(n int64) string {
    return strconv.FormatUint(*(*uint64)(unsafe.Pointer(&n)), 2)
}

