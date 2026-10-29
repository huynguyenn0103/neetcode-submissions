func isValid(bytes []byte)bool{
	num, _ := strconv.Atoi(string(bytes))
	return num >= 1 && num <= 26
}


func numDecodings(s string) int {
    n := len(s)
	result := make([]int, n)
	for i:=n-1; i>=0; i--{
		if s[i]=='0'{continue}
		tmp := 0
		bytes := make([]byte,0)
		for j:=i; j <n; j++{
			bytes = append(bytes, s[j])
			if !isValid(bytes){
				break
			}
			if j == n-1{
				tmp = tmp + 1
			} else {
				tmp = tmp + result[j+1]
			}
		}
		result[i]=tmp
	}
	return result[0]
}
