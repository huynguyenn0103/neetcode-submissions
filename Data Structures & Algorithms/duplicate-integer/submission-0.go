func hasDuplicate(nums []int) bool {
    mapVisited := make(map[int]bool)
    for _, num := range nums {
        if mapVisited[num]{return true}
        mapVisited[num]=true
    }
    return false
}
