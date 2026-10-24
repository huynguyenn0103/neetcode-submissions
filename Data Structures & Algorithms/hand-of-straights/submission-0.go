import "slices"
type GroupInfo struct {
	size int
	max int
}

func isNStraightHand(hand []int, groupSize int) bool {
   if len(hand) % groupSize != 0 {return false}
   mapValToLatestGroup := make(map[int]int)
   mapGroupInfo := make(map[int]GroupInfo)
   slices.SortFunc(hand, func(i1, i2 int)int{return i1-i2})
   startGroup := 0
   // 1 2 2 3 3 4 4 5
   for _, val := range hand {
	  var groupVal int
	  if _, existed := mapValToLatestGroup[val]; existed {
		  mapValToLatestGroup[val]++
		  groupVal = mapValToLatestGroup[val]
	  } else {
		 mapValToLatestGroup[val] = startGroup
		 groupVal = startGroup
	  }
	  var newSize int

	  if mapGroupInfo[groupVal].size == 0{
		newSize = 1
	  } else {
		if val - mapGroupInfo[groupVal].max != 1 {return false}
		newSize = mapGroupInfo[groupVal].size + 1
	  }
	  mapGroupInfo[groupVal] = GroupInfo{size: newSize, max: val}
	  if newSize == groupSize {
		startGroup++
	  }
   }
	return len(mapGroupInfo) == (len(hand) / groupSize)
//    0: 1 2 3 4
//    1: 2 3 4

//    1: 1
//    2: 2 
//    3: 2
//    4: 2
//    5:
}
// 1,2,2,3,3,4,4,5
// 1,2,3,3,3,4,4,5,5