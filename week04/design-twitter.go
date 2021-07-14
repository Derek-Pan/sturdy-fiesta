package main

//设计推特

// 二叉堆中存储堆数据结构
type Node struct {
	time      int64
	twitterId int
}

// 二叉堆实现方式
type BinaryHeap struct {
	heap []Node
}

// push元素,从底向上比较
func (b *BinaryHeap) Push(node Node) {
	b.heap = append(b.heap, node)
	size := len(b.heap)
	father := size / 2
	curr := size - 1
	for father >= 1 {
		if node.time > b.heap[father].time {
			tmp := b.heap[father]
			b.heap[father] = node
			b.heap[curr] = tmp
			curr = father
			father = curr / 2
		} else {
			break
		}
	}
}

// pop一个元素, 首尾交换,从顶向下比较
func (b *BinaryHeap) Pop() Node {
	node := b.heap[1]
	size := len(b.heap) - 1
	b.heap[1] = b.heap[size]
	b.heap = b.heap[0:size]

	root := 1
	child := root * 2

	for child < size {
		otherChild := child + 1
		if otherChild < size && b.heap[otherChild].time > b.heap[child].time {
			child = otherChild
		}
		if b.heap[root].time < b.heap[child].time {
			tmp := b.heap[root]
			b.heap[root] = b.heap[child]
			b.heap[child] = tmp
			root = child
			child = 2 * root
		} else {
			break
		}
	}
	return node
}

// 判断二叉堆是否为空
func (b *BinaryHeap) IsEmpty() bool {
	return len(b.heap) == 1
}

// Twitter数据结构
type Twitter struct {
	sequence     int64                    // 用来计算twitter发送顺序
	follows      map[int]map[int]struct{} // 存放关注列表
	sendTwitters map[int][]Node           // 每个用户发送的推文列表
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		sequence:     0,
		follows:      make(map[int]map[int]struct{}),
		sendTwitters: make(map[int][]Node),
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	if _, ok := this.sendTwitters[userId]; !ok {
		this.sendTwitters[userId] = make([]Node, 0, 0)
	}
	st := Node{
		time:      this.sequence,
		twitterId: tweetId,
	}
	this.sequence++
	this.sendTwitters[userId] = append(this.sendTwitters[userId], st)
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	ans := make([]int, 0, 0)
	binHeap := &BinaryHeap{
		heap: make([]Node, 1, 1),
	}
	// 这里感觉只要关注列表里的用户每个取前10条推文就可以了, 没有必要全部获取
	for _, twt := range this.sendTwitters[userId] {
		binHeap.Push(twt)
	}
	for fw := range this.follows[userId] {
		for _, twt := range this.sendTwitters[fw] {
			binHeap.Push(twt)
		}
	}
	for len(ans) < 10 && !binHeap.IsEmpty() {
		ans = append(ans, binHeap.Pop().twitterId)
	}
	return ans
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	_, ok := this.follows[followerId]
	if !ok {
		this.follows[followerId] = map[int]struct{}{followeeId: {}}
	} else {
		this.follows[followerId][followeeId] = struct{}{}
	}
	return
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	if _, ok := this.follows[followerId]; !ok {
		return
	}
	if _, ok := this.follows[followerId][followeeId]; !ok {
		return
	}
	delete(this.follows[followerId], followeeId)
	return
}

/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */
