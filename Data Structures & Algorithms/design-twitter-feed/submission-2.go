var time = 0

func getTick()int{
	time++
	return time
}

type MaxHeap [][]Tweet
func (h MaxHeap) Len()int{return len(h)}
func (h MaxHeap) Less(i, j int)bool{return h[i][len(h[i])-1].time > h[j][len(h[j])-1].time}
func (h MaxHeap) Swap(i, j int){h[i], h[j] = h[j], h[i]}
func (h *MaxHeap) Push(x interface{}){*h = append(*h, x.([]Tweet))}
func (h *MaxHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

type Tweet struct{
	tweetId int
	time int
}

func NewTweet(tweetId int)Tweet{
	return Tweet{
		tweetId: tweetId,
		time: getTick(),
	}
}

type Twitter struct {
    mapFollowerToFollowee map[int]map[int]bool
	mapUserToTweets map[int][]Tweet
}

func Constructor() Twitter {
    return Twitter{
		mapFollowerToFollowee : make(map[int]map[int]bool),
		mapUserToTweets : make(map[int][]Tweet),
	}
}


func (this *Twitter) PostTweet(userId int, tweetId int)  {
    newTweet := NewTweet(tweetId)
	if this.mapUserToTweets[userId] == nil {
		this.mapUserToTweets[userId] = make([]Tweet,0)
	}
	this.mapUserToTweets[userId] = append(this.mapUserToTweets[userId], newTweet)
}


func (this *Twitter) GetNewsFeed(userId int) []int {
    maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	followees := this.mapFollowerToFollowee[userId]
	if followees != nil {
		for followee, stillFollow := range followees{
			if !stillFollow{continue}
			tweets := this.mapUserToTweets[followee]
			if tweets != nil {
				heap.Push(maxHeap, tweets)
			}
		}
	}
	userTweets := this.mapUserToTweets[userId]
	if userTweets != nil{
		heap.Push(maxHeap, userTweets)
	}

	result := make([]int,0)
	for maxHeap.Len() > 0 && len(result) < 10 {
		tweets := heap.Pop(maxHeap).([]Tweet)
		result = append(result, tweets[len(tweets)-1].tweetId)
		tweets = tweets[:len(tweets)-1]
		if len(tweets) > 0 {
			heap.Push(maxHeap, tweets)
		}
	}
	return result
}


func (this *Twitter) Follow(followerId int, followeeId int)  {
    if this.mapFollowerToFollowee[followerId] == nil {
		this.mapFollowerToFollowee[followerId] = make(map[int]bool)
	}
	this.mapFollowerToFollowee[followerId][followeeId]=true
}


func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    if this.mapFollowerToFollowee[followerId] != nil {
		this.mapFollowerToFollowee[followerId][followeeId]=false
	}
}
