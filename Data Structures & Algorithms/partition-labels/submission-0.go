type Interval struct {
	start int
	end int
}

func isOverlapAndGetCandidate(i1, i2 Interval) (bool, Interval){
	if i1.end < i2.start {
		return false, i2
	} else {
		if i1.end > i2.end {
			return true, i1
		}
		return true, Interval{start: i1.start, end: i2.end}
	}
}

func partitionLabels(s string) []int {
    intervals := make(map[byte]Interval)
	listUniqueChar := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		c := s[i]
		if _, existed := intervals[c]; existed {
			intervals[c]=Interval{start: intervals[c].start, end: i}
		} else {
			intervals[c]=Interval{start: i, end: i}
			listUniqueChar = append(listUniqueChar, c)
		}
	}
	result := make([]int, 0)
	candidate := intervals[listUniqueChar[0]]
	for i := 1; i < len(listUniqueChar); i++ {
		c := listUniqueChar[i]
		isOverlap, newCandidate := isOverlapAndGetCandidate(candidate, intervals[c])
		if !isOverlap {
			result = append(result, candidate.end-candidate.start+1)
		} 
		if i == len(listUniqueChar)-1{
			result = append(result, newCandidate.end-newCandidate.start+1)
		}
		candidate = newCandidate
	}
	return result
}

// x: 0:3
// y: 1:4
// z: 5:7
// b: 6:9
// i: 10:10
// s: 11:11
// l: 12:12