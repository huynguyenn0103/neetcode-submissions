type MyHashMap struct {
	hashMap map[int]int
}

func Constructor() MyHashMap {
    return MyHashMap{
		hashMap: make(map[int]int),
	}
}

func (this *MyHashMap) Put(key int, value int) {
    this.hashMap[key]=value
}

func (this *MyHashMap) Get(key int) int {
    if val, existed := this.hashMap[key]; !existed {
		return -1
	} else {
		return val
	}
}

func (this *MyHashMap) Remove(key int) {
    delete(this.hashMap,key)
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */