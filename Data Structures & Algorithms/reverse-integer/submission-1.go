const MIN = -1<<31
const MAX = 1<<31 - 1

func reverse(x int) int {
	var signed int
	if x < 0 {
		signed = -1
	} else {
		signed = 1
	}

	x = x /signed

	result := 0
	for x > 0 {
		remainder := x % 10
		result = result * 10 + remainder
		if signed * result > MAX || signed * result < MIN{
			return 0
		}
		x = x/10
	}
	return signed * result
}
