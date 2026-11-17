type Item struct{
	timestamp int
	val string
}
type TimeMap struct {
	timeMap map[string][]Item
}

func Constructor() TimeMap {
	return TimeMap{
		timeMap: make(map[string][]Item),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if this.timeMap[key] == nil {
		this.timeMap[key] = make([]Item,0)
	}
	this.timeMap[key] = append(this.timeMap[key], Item{timestamp: timestamp, val: value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	result := ""
	if this.timeMap[key] != nil {
		times := this.timeMap[key]
		n := len(times)
		left, right := 0, n-1
		for left <= right{
			mid := (left + right)/2
			if times[mid].timestamp == timestamp{
				return times[mid].val
			} else if times[mid].timestamp > timestamp{
				right = mid - 1
			} else {
				result = times[mid].val
				left = mid + 1
			}
		}
		
	}
	return result
}
