func abs(x int)int{
	if x >= 0 {return x}
	return -x
}

func diff(x,y byte)int{
	return abs(int(x)-int(y))
}


func scoreOfString(s string) int {
	result := 0
	for i := 1; i < len(s); i++ {
		result += diff(s[i],s[i-1])
	}
	return result
}
