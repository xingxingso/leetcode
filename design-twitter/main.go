/*
https://leetcode-cn.com/problems/design-twitter/

355. 设计推特

设计一个简化版的推特(Twitter)，可以让用户实现发送推文，关注/取消关注其他用户，能够看见关注人（包括自己）的最近十条推文。你的设计需要支持以下的几个功能：

	1. postTweet(userId, tweetId): 创建一条新的推文
	2. getNewsFeed(userId): 检索最近的十条推文。每个推文都必须是由此用户关注的人或者是用户自己发出的。推文必须按照时间顺序由最近的开始排序。
	3. follow(followerId, followeeId): 关注一个用户
	4. unfollow(followerId, followeeId): 取消关注一个用户

*/
package design_twitter

import "container/heap"

// type Twitter struct {}
// /** Initialize your data structure here. */
// func Constructor() Twitter {}
// /** Compose a new tweet. */
// func (this *Twitter) PostTweet(userId int, tweetId int) {}
// /** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
// func (this *Twitter) GetNewsFeed(userId int) []int {}
// /** Follower follows a followee. If the operation is invalid, it should be a no-op. */
// func (this *Twitter) Follow(followerId int, followeeId int) {}
// /** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
// func (this *Twitter) Unfollow(followerId int, followeeId int) {}
/**
 * Your Twitter object will be instantiated and called as such:
 * obj := Constructor();
 * obj.PostTweet(userId,tweetId);
 * param_2 := obj.GetNewsFeed(userId);
 * obj.Follow(followerId,followeeId);
 * obj.Unfollow(followerId,followeeId);
 */

// --- 自己

/*
方法一:

时间复杂度：
空间复杂度：
*/
type Twitter struct {
	follow     map[int][]int
	news       map[int][]tweet
	tweetCount int
}

type tweet struct {
	id    int
	count int
}

// 大顶堆类
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.tweetCount++
	this.news[userId] = append(this.news[userId], tweet{id: tweetId, count: this.tweetCount})
	// this.news[userId] = this.news[userId][:10]
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	h := &IntHeap{}
	res := make([]int, 0)
	for _, userId := range this.follow[userId] {
		news := this.news[userId]
		if len(news) == 0 {
			continue
		}
		heap.Push(h, news[len(news)-1])
	}

}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	this.follow[followerId] = append(this.follow[followerId], followeeId)
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	key := 0
	for k, userId := range this.follow[followerId] {
		if userId == followeeId {
			key = k
			break
		}
	}
	temp := make([]int, 0)
	temp = append(temp, this.follow[followerId][:key]...)
	temp = append(temp, this.follow[followerId][key+1:]...)
	this.follow[followerId] = temp
}
