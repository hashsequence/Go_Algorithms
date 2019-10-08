func plusOne(digits []int) []int {
    n := len(digits)
    carry := 0
    for i := n-1; i >= 0; i-- {
        if i == n-1 {
            if digits[i] == 9 {
                carry = 1 
                digits[i] = 0
            } else {
                digits[i] += 1
            }
            continue
        }
        if carry == 1 {
             if digits[i] == 9 {
                carry = 1 
                digits[i] = 0
             } else {
                 carry = 0
                 digits[i] += 1
             }
        } else {
             carry = 0
        }
    }
    if carry == 1 {
        digits = append([]int{1},digits...)
    }
    return digits
}
