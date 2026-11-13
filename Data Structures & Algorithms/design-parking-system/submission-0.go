type ParkingSystem struct {
	maxBig int
	maxMedium int
	maxSmall int
	curBig int
	curMedium int
	curSmall int
}
func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		maxBig: big,
		maxMedium: medium,
		maxSmall: small,
	}
}
func (this *ParkingSystem) AddCar(carType int) bool {
	if carType == 1{
		if this.curBig == this.maxBig {
			return false
		}
		this.curBig++
	} else if carType == 2{
		if this.curMedium == this.maxMedium{
			return false
		}
		this.curMedium++
	} else {
		if this.curSmall == this.maxSmall{
			return false
		}
		this.curSmall++
	}
	return true
}
