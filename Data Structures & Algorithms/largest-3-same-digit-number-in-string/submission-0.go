func largestGoodInteger(num string) string {
	result := make([]byte,0)
	curBytes := make([]byte,0)
	curBytes = append(curBytes, num[0])
	for i := 1; i < len(num); i++ {
		if len(curBytes) == 0 {
			curBytes = append(curBytes, num[i])
		} else {
			if num[i] == curBytes[0]{
				curBytes = append(curBytes, num[i])
				if len(curBytes) == 3{
					if len(result) == 0 || curBytes[0] > result[0]{
						result = curBytes
					}
					curBytes = make([]byte,0)
				}
			} else {
				curBytes = []byte{num[i]}
			}
		}
	}
	return string(result)
}
