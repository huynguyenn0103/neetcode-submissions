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
	if this.timeMap[key] != nil {
		n := len(this.timeMap[key])
		for i := n-1; i >= 0; i-- {
			item := this.timeMap[key][i]
			if item.timestamp <= timestamp{
				return item.val
			}
		}
	}
	return ""
}
