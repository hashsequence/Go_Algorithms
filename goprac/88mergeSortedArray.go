func merge(nums1 []int, m int, nums2 []int, n int) {
    //zero array cases
    if n == 0 {
        return 
    }
    if m == 0 {
        copy(nums1,nums2[:n])
        return
    }
    //iterate backwards so you dont overwrite the elements in nums1
    k := m+n-1
    i := m-1
    j := n-1
    for ; k>=0; k-- {
        if i >= 0 && j >= 0 {
            if nums1[i] < nums2[j] {
                nums1[k] = nums2[j]
                j--
            } else {
                nums1[k] = nums1[i]
                i--
            }
        } else if j >= 0 {
            nums1[k] = nums2[j]
            j--
        }
    }
    
    
}

//inplace merge O((m+n)log(n+m)) using sort
func inplaceMergeWithSort(nums1 []int, m int, nums2 []int, n int)  {
    //zero array cases
    if n == 0 {
        return 
    }
    if m == 0 {
        copy(nums1,nums2[:n])
        return
    }
    copy(nums1[m:],nums2)
    sort.Slice(nums1, func(i,j int) bool {
        if nums1[i] < nums1[j] {
            return true
        }
        return false
    })

}

