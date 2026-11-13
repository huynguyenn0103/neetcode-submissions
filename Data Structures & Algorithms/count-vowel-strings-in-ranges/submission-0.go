func isStartAndEndWithVowel(str string)bool{
	vowel := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	return vowel[str[0]] && vowel[str[len(str)-1]]
}

func vowelStrings(words []string, queries [][]int) []int {
   prefixSum := make([]int, 0)
   sum := 0
   for _, word := range words{
	  if isStartAndEndWithVowel(word){
		sum = sum + 1 
	  }
	  prefixSum = append(prefixSum, sum)
   } 
   result := make([]int, 0)
   for _, query := range queries{
	  from, to := query[0], query[1]
	  if from == 0 {
		result = append(result, prefixSum[to])
	  } else {
		result = append(result, prefixSum[to]-prefixSum[from-1])
	  }
   }
   return result
}