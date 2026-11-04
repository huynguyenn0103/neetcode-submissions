var mapCharToInt = map[byte]int{
	'I':1,
	'V':5,
	'X':10,
	'L':50,
	'C':100,
	'D':500,
	'M':1000,
}
var changeDirectionCombination = map[string]bool{
	"IV": true,
	"IX": true,
	"XL": true,
	"XC": true,
	"CD": true,
	"CM": true,
}

func romanToInt(s string) int {
	n := len(s)
	result := 0
	for i := 0; i < n; i++ {
		if i < n-1 && changeDirectionCombination[s[i:i+2]] {
			result -= mapCharToInt[s[i]]
		} else {
			result += mapCharToInt[s[i]]
		}
	}
	return result
}
