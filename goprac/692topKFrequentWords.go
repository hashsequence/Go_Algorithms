/*
692. Top K Frequent Words
Medium

Given a non-empty list of words, return the k most frequent elements.

Your answer should be sorted by frequency from highest to lowest. If two words have the same frequency, then the word with the lower alphabetical order comes first.

Example 1:
Input: ["i", "love", "leetcode", "i", "love", "coding"], k = 2
Output: ["i", "love"]
Explanation: "i" and "love" are the two most frequent words.
    Note that "i" comes before "love" due to a lower alphabetical order.
Example 2:
Input: ["the", "day", "is", "sunny", "the", "the", "the", "sunny", "is", "is"], k = 4
Output: ["the", "is", "sunny", "day"]
Explanation: "the", "is", "sunny" and "day" are the four most frequent words,
    with the number of occurrence being 4, 3, 2 and 1 respectively.
Note:
You may assume k is always valid, 1 ≤ k ≤ number of unique elements.
Input words contain only lowercase letters.
Follow up:
Try to solve it in O(n log k) time and O(n) extra space.

solution:
easy part is to know to use a frequency map
and a minheap

the hard part is implementing comparison
TRICK PART: when inserting words with same frequency we want the biggest lexographic word on top but when checking whether to pop out the
word with same freq on top and put in another word with same freq, we want the word we are replacing it with to be lexographically smaller
so we cant use the compare function for the case of when to pop out items from the heap when it exceeds size k since it is opposite of
how the heap sorts the words with the same frequency

to be sure next time implement two seperate compare functions one for freq and another for word
*/

func topKFrequent(words []string, k int) []string {
    freqMap := map[string]int{}
    mh := NewMinHeap(k)
    
    for i, _ := range words {
        if _, ok := freqMap[words[i]]; !ok {
            freqMap[words[i]] = 1
        } else {
            freqMap[words[i]]++
        }
    }
    
    for key, val := range freqMap {
        if mh.size < k {
            mh.InsertKey(&Item{key, val})
        } else {
            if mh.GetMin().freq < val || (mh.GetMin().freq == val && mh.GetMin().word > key) {
                mh.ExtractMin()
                mh.InsertKey(&Item{key, val})
            } 
        }
    }
    res := make([]string,k)
    for i := k-1; mh.size > 0; i-- {
      
        res[i] = mh.ExtractMin().word    
    }
    return res
}


type Item struct {
    word string
    freq int
}

func (this *Item) Compare(other *Item) int {
    if this.freq < other.freq {
        return -1
    } else if this.freq > other.freq {
        return 1
    } 
    if this.word > other.word {
        return -1
    } else if this.word < other.word {
        return 1
    }
    return 0 
}

type MinHeap struct {
    arr []*Item
    size int
    cap int
}

func NewMinHeap(k int) MinHeap {
    return MinHeap{make([]*Item,0,k),0,k}
}

func Parent(i int) int {
    return (i-1)/2
}

func Left(i int) int {
    return 2*i+1
}

func Right(i int) int {
    return 2*i+2
}

func Swap(a ,b **Item) {
    *a, *b = *b, *a
}

func (this *MinHeap) MinHeapify(root int) {
    left := Left(root)
    right := Right(root)
    smallest := root
    
    if left < this.size && this.arr[left].Compare(this.arr[smallest]) == -1 {
        smallest = left
    }
    if right < this.size && this.arr[right].Compare(this.arr[smallest]) == -1 {
        smallest = right
    }
    
    if smallest != root {
        Swap(&this.arr[smallest], &this.arr[root])
        this.MinHeapify(smallest)
    }
}

func (this *MinHeap) GetMin() *Item {
    if this.size == 0 {
        return nil
    }
    return this.arr[0]
}

func (this *MinHeap) ExtractMin() *Item {
    if this.size == 0 {
        return nil
    }
    currMin := this.arr[0]
    this.size--
    this.arr[0] = this.arr[this.size]
    this.arr[this.size] = nil
    this.arr = this.arr[:this.size]
    this.MinHeapify(0)
    return currMin
}

func (this *MinHeap) InsertKey(k *Item) {
    if this.size == this.cap {
        tempArr := make([]*Item,this.size+1,this.cap * 2 + 1)
        copy(tempArr, this.arr)
        this.arr = nil
        this.arr = tempArr
    }
    
    this.arr = append(this.arr,k)
    this.size++
    
    for i := this.size-1; i >= 0 && this.arr[i].Compare(this.arr[Parent(i)]) == -1; i = Parent(i) {
        Swap(&this.arr[i], &this.arr[Parent(i)])
    }
    
    
    
}
