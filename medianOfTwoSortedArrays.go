func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    n := len(nums1)
    m := len(nums2)
    var k int = (n + m)/2
    if (m + n) % 2 != 0 {
       return kthelement(nums1, n, nums2, m, k+1) 
    } else {
       return (kthelement(nums1, n, nums2, m, k) + kthelement(nums1, n, nums2, m, k + 1))/2
    }

}



func kthelement(nums1 []int, n int, nums2 []int, m int , k int) float64 {
    if k <= 0 || k > (m + n) {
        return -1
    }
    if m < n {
        return kthelement(nums2, m, nums1, n, k)
    }
    
    if n == 0 {
        return float64(nums2[k - 1])
    }
    
    if k == 1 {
        return float64(min(nums1[0],nums2[0]))
    }
    
    i := min(k/2, n) 
    j := min(k/2, m)
    fmt.Println(i,j)
    
    if nums1[i-1] > nums2[j-1] {
        return kthelement(nums1, n, nums2[j:] , m - j, k - j )
    } else {
        return kthelement(nums1[i:], n - i, nums2, m, k - i )
    }
    
}
      

func min(a int, b int) int {
    var c int
    if c = b; a < b {
    c = a
    }
    return c
}
