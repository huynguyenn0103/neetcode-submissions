const NumOfRecent = 10
var curTimeStamp = 1

type Tweet struct {
	id int
	timestamp int
}

type MaxHeap []Tweet
func (h MaxHeap) Len()int{return len(h)}
func (h MaxHeap) Less(i,j int)bool{return h[i].timestamp > h[j].timestamp}
func (h MaxHeap) Swap(i,j int){h[i],h[j]=h[j],h[i]}
func (h *MaxHeap) Push(x interface{}){*h = append(*h, x.(Tweet))}
func (h *MaxHeap) Pop()interface{}{
	n := len(*h)
	item := (*h)[n-1]
	*h = (*h)[:n-1]
	return item
}

type Twitter struct {
    userToFollower map[int]map[int]bool 
	userToTweet map[int][]Tweet
}


func Constructor() Twitter {
	return Twitter{
		userToFollower: make(map[int]map[int]bool),
		userToTweet: make(map[int][]Tweet),
	}
}

func (this *Twitter) PostTweet(userId int, tweetId int)  {
    if this.userToTweet[userId] == nil {
		this.userToTweet[userId]=make([]Tweet, 0)
	}
	this.userToTweet[userId] = append(this.userToTweet[userId], Tweet{id: tweetId, timestamp: curTimeStamp})
	curTimeStamp++
}


func (this *Twitter) GetNewsFeed(userId int) []int {
    maxHeap := &MaxHeap{}
	heap.Init(maxHeap)
	if this.userToTweet[userId] != nil {
		for _, tweet := range this.userToTweet[userId]{
			heap.Push(maxHeap, tweet)
		}
	}
	if this.userToFollower[userId] != nil {
		for followerId, isStillFollow := range this.userToFollower[userId] {
			if isStillFollow {
				for _, tweet := range this.userToTweet[followerId]{
					heap.Push(maxHeap, tweet)
				} 
			}
		}
	}
	numOfPost := NumOfRecent
	result := make([]int, 0)
	for numOfPost > 0 && maxHeap.Len() > 0 {
		result = append(result, heap.Pop(maxHeap).(Tweet).id)
		numOfPost--
	}
	return result
}

// log(1)
func (this *Twitter) Follow(followerId int, followeeId int)  {
    if this.userToFollower[followerId] == nil {
		this.userToFollower[followerId] = make(map[int]bool)
	}
	this.userToFollower[followerId][followeeId] = true
}

// log(1)
func (this *Twitter) Unfollow(followerId int, followeeId int)  {
    if this.userToFollower[followerId] != nil {
		this.userToFollower[followerId][followeeId] = false
	}	
}
