func atoi(num string) []int {
	result := make([]int, 0)
	for i := 0; i < len(num); i++ {
		result = append(result, int(num[i]-'0'))
	}
	return result
}

func multiplyOneDigit(nums1 []int, num2 int) []int {
	if num2 == 0 {
		return []int{0}
	}
	n := len(nums1)
	base := make([]int, 0)
	for i := 0; i < n; i++ {
		base = append(base, nums1[n-1-i])
	}
	carry := 0
	for i := 0; i < n; i++ {
		newVal := base[i]*num2 + carry
		base[i] = newVal % 10
		carry = newVal / 10
	}
	if carry > 0 {
		base = append(base, carry)
	}
	return base
}

func multiply(num1 string, num2 string) string {
	nums1 := atoi(num1)
	nums2 := atoi(num2)
	n2 := len(nums2)
	maths := make([][]int, n2)
	results := make([]int, 0)
	for i := n2 - 1; i >= 0; i-- {
		maths[n2-1-i] = multiplyOneDigit(nums1, nums2[i])
	}

	carry := 0
	r := 0
	for {
		isUpdated := false
		newVal := carry
		for j := 0; j <= r; j++ {
			idx := r - j
			if j >= len(maths) || idx >= len(maths[j]) {
				continue
			}
			newVal += maths[j][idx]
			isUpdated = true
		}
		if !isUpdated {
			break
		}
		results = append(results, newVal%10)
		carry = newVal / 10
		r++
	}
	if carry != 0 {
		results = append(results, carry)
	}

	m := len(results)
	for i := 0; i < m/2; i++ {
		results[i], results[m-1-i] = results[m-1-i], results[i]
	}
	resultsInBytes := make([]byte, 0)
	for _, val := range results {
		resultsInBytes = append(resultsInBytes, byte(val+'0'))
	}
	for {
		if len(resultsInBytes) == 1 || resultsInBytes[0] != 0 {
			break
		}
		resultsInBytes = resultsInBytes[1:]
	}
	return string(resultsInBytes)
}