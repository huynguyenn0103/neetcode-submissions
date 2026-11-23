type CountSquares struct {
    mapPointToFreq map[[2]int]int
}

func Constructor() CountSquares {
    return CountSquares{
		mapPointToFreq : make(map[[2]int]int),
	}
}

func abs(x int)int{
	if x >= 0 {return x}
	return -x
}

func validDiagonal(p1,p2 [2]int)bool{
	if abs(p1[0]-p2[0]) == abs(p1[1]-p2[1]) && p1[0]-p2[0] != 0 {
		return true
	}
	return false
}

func (this *CountSquares) Add(point []int)  {
    p := [2]int{point[0],point[1]}
	this.mapPointToFreq[p]++
}

func (this *CountSquares) Count(point []int) int {
	target := [2]int{point[0],point[1]}
	result := 0
    for c := range this.mapPointToFreq{
		if validDiagonal(target,c) && this.mapPointToFreq[[2]int{c[0],target[1]}] > 0 && this.mapPointToFreq[[2]int{target[0],c[1]}] > 0{
			result += this.mapPointToFreq[c]*this.mapPointToFreq[[2]int{c[0],target[1]}]*this.mapPointToFreq[[2]int{target[0],c[1]}]
		}
	}
	return result
}
