
func canConstruct(ransomNote string, magazine string) bool {
	mapCharToFreq := make(map[byte]int)
	for i := 0; i < len(magazine); i++{
		mapCharToFreq[magazine[i]]++
	}
	for i := 0; i < len(ransomNote); i++{
		mapCharToFreq[ransomNote[i]]--
		if mapCharToFreq[ransomNote[i]] < 0 {
			return false
		}
	}
	return true
}
