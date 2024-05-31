package HashSet

//Design a HashSet without using any built-in hash table libraries.
//
//Implement MyHashSet class:
//
//void add(key) Inserts the value key into the HashSet.
//bool contains(key) Returns whether the value key exists in the HashSet or not.
//void remove(key) Removes the value key in the HashSet. If key does not exist in the HashSet, do nothing.

type MyHashSet struct {
	set map[int]struct{}
}

func Constructor() MyHashSet {
	return MyHashSet{set: make(map[int]struct{})}
}

func (this *MyHashSet) Add(key int) {
	this.set[key] = struct{}{}
}

func (this *MyHashSet) Remove(key int) {
	delete(this.set, key)
}

func (this *MyHashSet) Contains(key int) bool {
	_, ok := this.set[key]
	return ok
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
