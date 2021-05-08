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

// --- 他人
// https://mp.weixin.qq.com/s?__biz=MzAxODQxMDM0Mw==&mid=2247484499&idx=1&sn=64f75d4bdbb4c5777ba199aae804d138&chksm=9bd7fa5baca0734dc51f588af913140560b994e3811dac6a7fa8ccfc2a31aca327f1faf964c2&scene=21#wechat_redirect

/*
方法一: 堆

时间复杂度：
空间复杂度：
*/
type Twitter struct {
	userMap map[int]*User
}

var timestamp int

type Tweet struct {
	id   int
	time int
	next *Tweet
}

type User struct {
	id       int
	followed map[int]struct{}
	head     *Tweet
}

func (user *User) follow(userId int) {
	user.followed[userId] = struct{}{}
}

func (user *User) unfollow(userId int) {
	// 不可以取关自己
	if userId != user.id {
		delete(user.followed, userId)
	}
}

func (user *User) post(tweetId int) {
	tweet := &Tweet{id: tweetId, time: timestamp}
	timestamp++
	// 将新建的推文插入链表头
	// 越靠前的推文 time 值越大
	tweet.next = user.head
	user.head = tweet
}

/** Initialize your data structure here. */
func Constructor() Twitter {
	return Twitter{
		userMap: map[int]*User{},
	}
}

func (this *Twitter) createIfUserNotExit(userId int) {
	var user *User
	var ok bool
	// 若 userId 不存在，则新建
	if user, ok = this.userMap[userId]; !ok {
		user = &User{
			id:       userId,
			followed: map[int]struct{}{},
		}
		// 关注自己
		user.follow(userId)
		this.userMap[userId] = user
	}
}

/** Compose a new tweet. */
func (this *Twitter) PostTweet(userId int, tweetId int) {
	this.createIfUserNotExit(userId)
	this.userMap[userId].post(tweetId)
}

/** Follower follows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Follow(followerId int, followeeId int) {
	this.createIfUserNotExit(followerId)
	this.createIfUserNotExit(followeeId)
	this.userMap[followerId].follow(followeeId)
}

/** Follower unfollows a followee. If the operation is invalid, it should be a no-op. */
func (this *Twitter) Unfollow(followerId int, followeeId int) {
	this.createIfUserNotExit(followerId)
	this.createIfUserNotExit(followeeId)
	this.userMap[followerId].unfollow(followeeId)
}

// 大顶堆类
type IntHeap []*Tweet

func (h IntHeap) Len() int            { return len(h) }
func (h IntHeap) Less(i, j int) bool  { return h[i].time > h[j].time }
func (h IntHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x interface{}) { *h = append(*h, x.(*Tweet)) }
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

/** Retrieve the 10 most recent tweet ids in the user's news feed. Each item in the news feed must be posted by users who the user followed or by the user herself. Tweets must be ordered from most recent to least recent. */
func (this *Twitter) GetNewsFeed(userId int) []int {
	if _, ok := this.userMap[userId]; !ok {
		return nil
	}
	users := this.userMap[userId].followed
	h := &IntHeap{}
	for id := range users {
		twt := this.userMap[id].head
		if twt == nil {
			continue
		}
		heap.Push(h, twt)
	}

	res := make([]int, 0)
	for h.Len() > 0 {
		// 最多返回10条就够了
		if len(res) >= 10 {
			break
		}
		twt := heap.Pop(h).(*Tweet)
		res = append(res, twt.id)
		if twt.next != nil {
			heap.Push(h, twt.next)
		}
	}
	return res
}
