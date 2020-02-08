/*
380. Insert Delete GetRandom O(1)
Medium

Design a data structure that supports all following operations in average O(1) time.

insert(val): Inserts an item val to the set if not already present.
remove(val): Removes an item val from the set if present.
getRandom: Returns a random element from current set of elements. Each element must have the same probability of being returned.
Example:

// Init an empty set.
RandomizedSet randomSet = new RandomizedSet();

// Inserts 1 to the set. Returns true as 1 was inserted successfully.
randomSet.insert(1);

// Returns false as 2 does not exist in the set.
randomSet.remove(2);

// Inserts 2 to the set, returns true. Set now contains [1,2].
randomSet.insert(2);

// getRandom should return either 1 or 2 randomly.
randomSet.getRandom();

// Removes 1 from the set, returns true. Set now contains [2].
randomSet.remove(1);

// 2 was already in the set, so return false.
randomSet.insert(2);

// Since 2 is the only number in the set, getRandom always return 2.
randomSet.getRandom();

solution:
an easy solution is to use a map to keep track of what has been inserted and deleted
use the array to store the values and the maps value to store the index of the array
when you delete an element, to preserve index integrity swap the element at the end of 
the array with the val you want to remove in the array, set the index of the element originally
at the end of the array to the val you want to remove's index and decrement the index 

*/


type RandomizedSet struct {
    set map[int]*int
    arr []*int
}


/** Initialize your data structure here. */
func Constructor() RandomizedSet {
    set := map[int]*int{}
    arr := make([]*int,0)
    return RandomizedSet{set, arr}
}


/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
    if v, ok := this.set[val]; (ok && v == nil) || !ok {
        this.arr = append(this.arr, &val)
        index :=  len(this.arr)-1
        this.set[val] = &index
        return true
    }
    return false
}


/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
    if v, ok := this.set[val]; ((ok && v != nil)) && len(this.arr) > 0 {
        //take the element from the end and put it into val's current place in array
        this.arr[*this.set[val]] = this.arr[len(this.arr)-1] 
        this.set[*this.arr[len(this.arr)-1]] = this.set[val]
        this.set[val] = nil
        this.arr[len(this.arr)-1] = nil
        this.arr = this.arr[:len(this.arr)-1]
        return true
    }
    
    return false
}


/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
    if len(this.arr) > 0 {
        //fmt.Println(this.arr)
        return *this.arr[rand.Intn(len(this.arr))]
    }
    return -1
}


/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
