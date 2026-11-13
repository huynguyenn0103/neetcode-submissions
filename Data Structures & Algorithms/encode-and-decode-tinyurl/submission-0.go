type Solution struct {}
func Constructor() Solution {
	return Solution{}
}
func (this *Solution) Encode(longUrl string) string {
	return longUrl
}
func (this *Solution) Decode(shortUrl string) string {
	return shortUrl
}
