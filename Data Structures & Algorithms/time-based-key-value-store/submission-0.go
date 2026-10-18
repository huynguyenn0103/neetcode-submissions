type ValueItem struct {
	val string
	timestamp int
}

type TimeMap struct {
	store map[string][]ValueItem
}

func Constructor() TimeMap {
	return TimeMap{
		store: make(map[string][]ValueItem),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if this.store[key] == nil {
		this.store[key] = make([]ValueItem, 0)
	}
	this.store[key]=append(this.store[key], ValueItem{val: value, timestamp: timestamp})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	listOfEvents := this.store[key]
	n := len(listOfEvents)
	if n == 0 {return ""}
	left, right := 0, n-1
	var result string
	for left <= right {
		mid := (left + right) / 2
		itemInMid := listOfEvents[mid]
		if itemInMid.timestamp == timestamp{
			result = itemInMid.val
			break
		} else if itemInMid.timestamp > timestamp{
			right = mid - 1
		} else {
			left = mid + 1
			result = itemInMid.val
		}
	}
	return result
}
