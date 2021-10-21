package leetcode

import "math/rand"

/*
Runtime: 180 ms, faster than 64.25% of Go online submissions for Insert Delete GetRandom O(1).
Memory Usage: 50.2 MB, less than 74.30% of Go online submissions for Insert Delete GetRandom O(1).

*/

type RandomizedSet struct {
	ObjMap map[int]int
	Obj    []int
}

func Constructor() RandomizedSet {
	return RandomizedSet{Obj: []int{}, ObjMap: make(map[int]int)}
}

/*
Insert bool insert(int val) Inserts an item val into the set if not present.
Returns true if the item was not present, false otherwise.
*/
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.ObjMap[val]; ok {
		return false
	}
	this.ObjMap[val] = len(this.Obj)
	this.Obj = append(this.Obj, val)
	return true
}

/*
Remove bool remove(int val) Removes an item val from the set if present.
Returns true if the item was present, false otherwise.
*/
func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.ObjMap[val]
	if !ok {
		return false
	}

	last := this.Obj[len(this.Obj)-1]
	if val == last { // if we are removing the last element in the array, we don't need to swap and update index in map
		this.Obj = this.Obj[:len(this.Obj)-1]
	} else {
		// swap and remove the last element makes the delete of the element O(1)
		this.Obj[idx] = last
		this.Obj = this.Obj[:len(this.Obj)-1]
		// note: we also need to update the index of the last element in the map
		this.ObjMap[this.Obj[idx]] = idx
	}
	delete(this.ObjMap, val)
	return true

	return true
}

/*
GetRandom Returns a random element from the current set of elements (it's guaranteed that at least one element exists when this method is called).
Each element must have the same probability of being returned.
*/
func (this *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(this.Obj))
	return this.Obj[index]
}
