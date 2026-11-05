//010
//011
//110
//111

//x=10
//n=10

//x=110
//n=100
//10110

func minEnd(n int, x int) int {
	binaryN := strconv.FormatInt(int64(n-1), 2)
	binaryX := strconv.FormatInt(int64(x),2)
	m,n := len(binaryN),len(binaryX)
	ptrN,ptrX := m-1,n-1
	results := make([]byte,0)
	for ptrN >= 0 && ptrX >= 0{
		if binaryX[ptrX] == '0'{
			results = append(results, binaryN[ptrN])
			ptrN--
		} else {
			results = append(results, binaryX[ptrX])
		}
		ptrX--
	}
	for ptrX >= 0 {
		results = append(results, binaryX[ptrX])
		ptrX--
	}
	for ptrN >= 0 {
		results = append(results, binaryN[ptrN])
		ptrN--
	}
	l := len(results)
	for i := 0; i < l/2; i++ {
		results[i],results[l-1-i]=results[l-1-i],results[i]
	}
	result, _ := strconv.ParseInt(string(results), 2, 0)
	return int(result)
}
