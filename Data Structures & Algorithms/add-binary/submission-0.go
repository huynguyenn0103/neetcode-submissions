func addBinary(a string, b string) string {
	m,n := len(a),len(b)
	ptrA,ptrB := m-1,n-1
	carry := 0
	resultInInt := make([]int,0)
	for ptrA >= 0 && ptrB >= 0{
		var sum int
		if (a[ptrA] == '1' && b[ptrB] == '0')||(a[ptrA] == '0' && b[ptrB] == '1'){
			if carry == 1{
				sum = 0
			} else {
				sum = 1
				carry = 0
			}
		} else if a[ptrA] == '0' && b[ptrB] == '0'{
			if carry == 1{
				sum = 1
			} else {
				sum = 0
			}
			carry = 0
		} else {
			if carry == 1{
				sum = 1
			} else {
				sum = 0
			}
			carry = 1
		}
		resultInInt = append(resultInInt, sum)
		ptrB--
		ptrA--
	}
	for ptrA >= 0 {
		var sum int
		if a[ptrA] == '0'{
			if carry == 1{
				sum = 1
			} else {
				sum = 0
			}
			carry = 0
		} else {
			if carry == 1{
				sum = 0
			} else {
				sum = 1
			}
		}
		resultInInt = append(resultInInt, sum)
		ptrA--
	}
	for ptrB >= 0 {
		var sum int
		if b[ptrB] == '0'{
			if carry == 1{
				sum = 1
			} else {
				sum = 0
			}
			carry = 0
		} else {
			if carry == 1{
				sum = 0
			} else {
				sum = 1
			}
		}
		resultInInt = append(resultInInt, sum)
		ptrB--
	}
	if carry == 1{
		resultInInt = append(resultInInt, carry)
	}
	resultInBytes := make([]byte,0)
	l := len(resultInInt)
	for i:=l-1; i>=0; i--{
		resultInBytes = append(resultInBytes, byte(resultInInt[i]+int('0')))
	}
	return string(resultInBytes)
}
