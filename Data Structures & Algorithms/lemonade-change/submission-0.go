func lemonadeChange(bills []int) bool {
	mapMoneyToNum := make(map[int]int)
	for _, bill := range bills {
		if bill == 5 {
			mapMoneyToNum[5]++
		} else if bill == 10 {
			if mapMoneyToNum[5] > 0 {
				mapMoneyToNum[5]--
				mapMoneyToNum[10]++
			} else {
				return false
			}
		} else {
			if mapMoneyToNum[10] > 0 && mapMoneyToNum[5] > 0 {
				mapMoneyToNum[10]--
				mapMoneyToNum[5]--
				mapMoneyToNum[20]++
			} else if mapMoneyToNum[5] > 2 {
				mapMoneyToNum[5] = mapMoneyToNum[5]-3
				mapMoneyToNum[20]++
			} else {
				return false
			}
		}
	}
	return true
}
