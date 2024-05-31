package HashMap

//Design a HashMap without using any built-in hash table libraries.
//
//Implement the MyHashMap class:
//
//MyHashMap() initializes the object with an empty map.
//void put(int key, int value) inserts a (key, value) pair into the HashMap. If the key already exists in the map, update the corresponding value.
//int get(int key) returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key.
//void remove(key) removes the key and its corresponding value if the map contains the mapping for the key.

const BucketSize = 1000

type pair struct {
	key   int
	value int
}

type bucket struct {
	pairs []pair
}

type MyHashMap struct {
	buckets []bucket
}

func Constructor() MyHashMap {
	return MyHashMap{buckets: make([]bucket, BucketSize)}
}

func (this *MyHashMap) hash(key int) int {
	return key % BucketSize
}

func (this *MyHashMap) Put(key int, value int) {
	b := &this.buckets[this.hash(key)]
	for i, p := range b.pairs {
		if p.key == key {
			b.pairs[i].value = value
			return
		}
	}
	b.pairs = append(b.pairs, pair{key: key, value: value})
}

func (this *MyHashMap) Get(key int) int {
	b := &this.buckets[this.hash(key)]
	for _, p := range b.pairs {
		if p.key == key {
			return p.value
		}
	}
	return -1
}

func (this *MyHashMap) Remove(key int) {
	b := &this.buckets[this.hash(key)]
	for i, p := range b.pairs {
		if p.key == key {
			b.pairs = append(b.pairs[:i], b.pairs[i+1:]...)
			return
		}
	}
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */
