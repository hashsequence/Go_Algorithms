/*
460. LFU Cache
Hard

Design and implement a data structure for Least Frequently Used (LFU) cache. It should support the following operations: get and put.

get(key) - Get the value (will always be positive) of the key if the key exists in the cache, otherwise return -1.
put(key, value) - Set or insert the value if the key is not already present. When the cache reaches its capacity, it should invalidate the least frequently used item before inserting a new item. For the purpose of this problem, when there is a tie (i.e., two or more keys that have the same frequency), the least recently used key would be evicted.

Note that the number of times an item is used is the number of calls to the get and put functions for that item since it was inserted. This number is set to zero when the item is removed.

 

Follow up:
Could you do both operations in O(1) time complexity?

 

Example:

LFUCache cache = new LFUCache( 2 );

cache.put(1, 1);
cache.put(2, 2);
cache.get(1);       // returns 1
cache.put(3, 3);    // evicts key 2
cache.get(2);       // returns -1 (not found)
cache.get(3);       // returns 3.
cache.put(4, 4);    // evicts key 1.
cache.get(1);       // returns -1 (not found)
cache.get(3);       // returns 3
cache.get(4);       // returns 4*

solution:

use a map to map the cacheNode for O(1) access
use doubly linkedlist
use a frequencylist
for each have a cacheList where LRU is at the tail of the list
each cachenode should have pointer to frequency node 

easy to explain but TONS of edge cases
my advice is to implement FreqList and CacheList really well
*/
type LFUCache struct {
    lookup map[int]*CacheNode 
    freqList *FreqList
    capacity int
    size int
}


func Constructor(capacity int) LFUCache {
    return LFUCache{
        lookup : map[int]*CacheNode{},
        freqList : &FreqList{nil, nil, 0},
        capacity : capacity,
        size : 0,
    }
}


func (this *LFUCache) Get(key int) int {
    if this.capacity == 0 {
        return -1
    }
    cacheNode,_ := this.lookup[key]
    if cacheNode == nil {
        return -1
    }
    this.Put(cacheNode.key, cacheNode.value)
    return cacheNode.value
}

func (this* LFUCache) Evict()  {
    if this.size == 0 {
        return 
    }
    if this.size == this.capacity {
        this.lookup[this.freqList.head.cacheList.tail.key] = nil
        this.freqList.head.cacheList.PopBack()
        if this.freqList.head.cacheList.size == 0 {
            this.freqList.PopFront()
        }
        this.size--   
    }
}
func (this *LFUCache) Put(key int, value int)  {
    defer func() {
		if err := recover(); err != nil {
            this.PrintFreqList()
            os.Exit(3)
		}
	}()
    if this.capacity == 0 {
        return
    }
    cacheNode, _ := this.lookup[key]
    if cacheNode == nil {
        this.Evict()
        cacheNode = &CacheNode {
            key : key,
            value : value,
            prev : nil,
            next : nil,
        }
        if this.freqList.head == nil || (this.freqList.head != nil && this.freqList.head.freq > 1 ){
            newFreqNode := &FreqNode {
                freq : 1,
                cacheList : &CacheList{nil, nil, 0},
                next : nil,
                prev : nil,
            }
            this.freqList.PushFront(newFreqNode)
        } 
        cacheNode.freqNode = this.freqList.head
        this.freqList.head.cacheList.PushFront(cacheNode)
        this.size++
        this.lookup[key] = cacheNode
    } else {
        freqNode := cacheNode.freqNode
        cacheNodeFreq := freqNode.freq
        freqNode.cacheList.Remove(cacheNode)
        this.lookup[cacheNode.key] = nil
        this.size--
        if freqNode.next == nil || (freqNode.next != nil && freqNode.next.freq > cacheNodeFreq+1) {
            this.freqList.Insert(freqNode, &FreqNode{
                freq : cacheNodeFreq+1,
                cacheList : &CacheList{nil, nil, 0},
                next : nil,
                prev : nil,
            })
        } 
        cacheNode.freqNode = freqNode.next
        if freqNode.cacheList.size == 0 {
            this.freqList.Remove(freqNode)
        }
        cacheNode.value = value
        cacheNode.freqNode.cacheList.PushFront(cacheNode)
        this.lookup[key] = cacheNode
        this.size++
    }
}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

type FreqNode struct {
    freq int
    cacheList *CacheList
    next *FreqNode
    prev *FreqNode
}

type CacheNode struct {
    key int
    value int
    freqNode *FreqNode
    next *CacheNode
    prev *CacheNode
}

type FreqList struct {
    head *FreqNode
    tail *FreqNode
    size int
}

func (this *FreqList) PushFront(newFreqNode *FreqNode) {
    if newFreqNode == nil {
        return
    }
    newFreqNode.prev = nil
    newFreqNode.next = nil
    if this.head == nil {
        this.head = newFreqNode
        this.tail = newFreqNode
    } else {
        newFreqNode.next = this.head
        this.head.prev = newFreqNode
        newFreqNode.prev = nil
        this.head = newFreqNode
    }
    this.size++
}

func (this *FreqList) PushBack(newFreqNode *FreqNode) {
    if newFreqNode == nil {
        return
    }
    newFreqNode.prev = nil
    newFreqNode.next = nil
    if this.tail == nil {
        this.head = newFreqNode
        this.tail = newFreqNode
    } else {
        newFreqNode.prev = this.tail
        this.tail.next = newFreqNode
        newFreqNode.next = nil
        this.tail = newFreqNode
    }
    this.size++
}

func (this *FreqList) Insert(posNode *FreqNode, newFreqNode *FreqNode) {
    if newFreqNode == nil || posNode == nil {
        return
    }
    newFreqNode.prev = nil
    newFreqNode.next = nil
    if posNode == this.tail {
        this.PushBack(newFreqNode)
    } else {
        leftNode := posNode
        rightNode := posNode.next
        leftNode.next = newFreqNode
        newFreqNode.prev = leftNode
        newFreqNode.next = rightNode
        if rightNode != nil {
            rightNode.prev = newFreqNode
        }
        this.size++
    }
}


func (this *FreqList) Remove(currNode *FreqNode) {
    if currNode == nil || this.head == nil || this.tail == nil || this.size == 0 {
        return
    }
    if currNode == this.head {
        this.PopFront()
    } else if currNode == this.tail {
        this.PopBack()
    } else {
        leftNode := currNode.prev
        rightNode := currNode.next
        leftNode.next = rightNode
        rightNode.prev = leftNode
        this.size--
    }
}

func (this *FreqList) PopFront() {
   if this.head != nil {
        this.head = this.head.next
        if this.head != nil {
            this.head.prev = nil
        }
        this.size--
        if this.size == 0 {
            this.tail = nil
        }
    }
    
}

func (this *FreqList) PopBack() {
    if this.tail != nil {
        this.tail = this.tail.prev
        if this.tail != nil {
           this.tail.next = nil 
        }
        this.size--
        if this.size == 0 {
            this.head = nil
        }
    }
}

//********************************************************************************//
type CacheList struct {
    head *CacheNode
    tail *CacheNode
    size int
}

func (this *CacheList) PushFront(newCacheNode *CacheNode) {
    if newCacheNode == nil {
        return
    }
    newCacheNode.prev = nil
    newCacheNode.next = nil
    if this.head == nil {
        this.head = newCacheNode
        this.tail = newCacheNode
    } else {
        newCacheNode.next = this.head
        this.head.prev = newCacheNode
        newCacheNode.prev = nil
        this.head = newCacheNode
    }
    this.size++
}

func (this *CacheList) PushBack(newCacheNode *CacheNode) {
    if newCacheNode == nil {
        return
    }
    newCacheNode.prev = nil
    newCacheNode.next = nil
    if this.tail == nil {
        this.head = newCacheNode
        this.tail = newCacheNode
    } else {
        newCacheNode.prev = this.tail
        this.tail.next = newCacheNode
        newCacheNode.next = nil
        this.tail = newCacheNode
    }
    this.size++
}

func (this *CacheList) Insert(posNode *CacheNode, newCacheNode *CacheNode) {
    if newCacheNode == nil || posNode == nil {
        return
    }
    newCacheNode.prev = nil
    newCacheNode.next = nil
    if posNode == this.tail {
        this.PushBack(newCacheNode)
    } else {
        leftNode := posNode
        rightNode := posNode.next
        leftNode.next = newCacheNode
        newCacheNode.prev = leftNode
        newCacheNode.next = rightNode
        if rightNode != nil {
            rightNode.prev = newCacheNode
        }
        this.size++
    }
    
    
}


func (this *CacheList) Remove(currNode *CacheNode) {
    if currNode == nil || this.head == nil || this.tail == nil || this.size == 0 {
        return
    }
    if currNode == this.head {
        this.PopFront()
    } else if currNode == this.tail {
        this.PopBack()
    } else {
        leftNode := currNode.prev
        rightNode := currNode.next
        leftNode.next = rightNode
        rightNode.prev = leftNode
        this.size--
    }
}

func (this *CacheList) PopFront() {
    if this.head != nil {
        this.head = this.head.next
        if this.head != nil {
            this.head.prev = nil
        }
        this.size--
        if this.size == 0 {
            this.head = nil
        }
    }
}

func (this *CacheList) PopBack() {
    if this.tail != nil {
        this.tail = this.tail.prev
        if this.tail != nil {
           this.tail.next = nil 
        }
        this.size--
        if this.size == 0 {
            this.tail = nil
        }
    }
}

func (this* LFUCache) PrintFreqList() {
    fmt.Println("++++++++++++++++++++++++++++++++")
    fmt.Println("freqList size ", this.size, " freqHead: ", this.freqList.head," freqTail: " ,this.freqList.tail)
     for it := this.freqList.head; it != nil; it = it.next {
         fmt.Println("frequency: ",it.freq, " cacheList size: ", it.cacheList.size)
         for it2 := it.cacheList.head; it2 != nil; it2 = it2.next {
             fmt.Println("-- ", it2.value)
         }
    }
    fmt.Println("--------------------------------")
}
