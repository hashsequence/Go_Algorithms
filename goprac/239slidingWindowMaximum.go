/*
239. Sliding Window Maximum
Hard


Given an array nums, there is a sliding window of size k which is moving from the very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right by one position. Return the max sliding window.

Example:

Input: nums = [1,3,-1,-3,5,3,6,7], and k = 3
Output: [3,3,5,5,6,7] 
Explanation: 

Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       3
 1 [3  -1  -3] 5  3  6  7       3
 1  3 [-1  -3  5] 3  6  7       5
 1  3  -1 [-3  5  3] 6  7       5
 1  3  -1  -3 [5  3  6] 7       6
 1  3  -1  -3  5 [3  6  7]      7
Note: 
You may assume k is always valid, 1 ≤ k ≤ input array's size for non-empty array.

Follow up:
Could you solve it in linear time?

*/
func maxSlidingWindow(nums []int, k int) []int {
    if len(nums) == 0  || k == 0{
        return nums
    }

    res:= make([]int, 0)
    q := make(Queue,0)
    //queue always have k elements and is used as the window
    for i := 0; i < len(nums); i++ {
        //when queue is bigger than k pop the front of the queue since its no longer in window, and also we already added it into the results
        //so we already accounted the fact it was the biggest in one or more windows
        fmt.Println(q)
        for len(q) > 0 && isFrontOfQueueOutsideWindow(q[0], i, k) {
            q.PopFront()
        }
        //when current number we are looking at is bigger than the end of the queue, we pop the end of the queue since we know
        //that it is smaller it can be the biggest in the current window size
        for len(q) > 0 && isCurrentNumBiggerThanNumOfTheEndOfQueue(nums[i], nums[q.Back()]) {
            q.PopBack()
        }
        //push the current index into the back of the queue
        q.Push(i)
        //append the element in the front of the queue since we know that the element s the biggest in the current window since it wasnt
        //pushed off yet
        res = append(res, nums[q.Front()])
    }
    //we discount the first k-1 elements in nums since the window has not yet reached the first set of k yet and so it will be duplicates
    //fmt.Println(res)
    return res[k-1:]
    
}

type Queue []int 

func (q *Queue) Push(k int) {
    (*q) = append(*q,k)
}

func (q *Queue) Front() int {
    if len(*q) > 0 {
        return (*q)[0]
    }
    return -1
}

func (q *Queue) Back() (int) {
    if len(*q) > 0 {
        return (*q)[len(*q)-1]
    } 
    return -1
}

func (q *Queue) PopFront() {
    if len(*q) > 0 {
        (*q)[0] = 0
        (*q) = (*q)[1:]
    }
}

func (q *Queue) PopBack() {
    if len(*q) > 0 {
        (*q)[len(*q)-1] = 0
        (*q) = (*q)[:len(*q)-1]
    }
}


func isFrontOfQueueOutsideWindow(indexOfFrontOfQueue, currentIndex, windowSize int) bool {
    return currentIndex - indexOfFrontOfQueue >= windowSize
}

func isCurrentNumBiggerThanNumOfTheEndOfQueue(currNum, numOfEndOfQueue int) bool {
    return currNum > numOfEndOfQueue
}

