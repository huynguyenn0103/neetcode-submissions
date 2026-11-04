// AZ = 52 = 26*A + 26
// BA = 53 = 26*B + A

func log26(x int)int {
	return int(math.Log(float64(x))/math.Log(float64(26)))
}

func pow(x,n int)int{
	return int(math.Pow(float64(x),float64(n)))
}

func numToByte(x int)byte{
	return byte(int('A')-1+x)
}

func calBase(x int)int {
	base := 0
	sum := 26*pow(26,base)
	for x > sum{
		base++
		sum += 26*pow(26,base)
	}
	return base
}

func convertToTitle(columnNumber int) string {
	base := calBase(columnNumber)
	result := make([]byte, 0)
	for base >= 0 {
		div := columnNumber / pow(26,base)
		mod := columnNumber / pow(26,base)
		if mod == 0 && columnNumber != 26{
			div = div - 1
		}
		result = append(result, numToByte(div))
		columnNumber = columnNumber - div*pow(26,base)
		base--
	}
	return string(result)
}
