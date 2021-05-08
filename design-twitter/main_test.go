package design_twitter

import (
	"reflect"
	"testing"
)

func TestTwitter(t *testing.T) {
	s := Constructor()
	testEx1(&s, t)

	s2 := Constructor()
	testEx2(&s2, t)
}

type TwitterInterface interface {
	PostTweet(userId int, tweetId int)
	GetNewsFeed(userId int) []int
	Follow(followerId int, followeeId int)
	Unfollow(followerId int, followeeId int)
}

func testEx1(twitter TwitterInterface, t *testing.T) {

	// 用户1发送了一条新推文 (用户id = 1, 推文id = 5).
	twitter.PostTweet(1, 5)

	// 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
	if got := twitter.GetNewsFeed(1); !reflect.DeepEqual(got, []int{5}) {
		t.Errorf("GetNewsFeed() = %v, want %v", got, []int{5})
	}

	// 用户1关注了用户2.
	twitter.Follow(1, 2)

	// 用户2发送了一个新推文 (推文id = 6).
	twitter.PostTweet(2, 6)

	// 用户1的获取推文应当返回一个列表，其中包含两个推文，id分别为 -> [6, 5].
	// 推文id6应当在推文id5之前，因为它是在5之后发送的.
	if got := twitter.GetNewsFeed(1); !reflect.DeepEqual(got, []int{6, 5}) {
		t.Errorf("GetNewsFeed() = %v, want %v", got, []int{6, 5})
	}

	// 用户1取消关注了用户2.
	twitter.Unfollow(1, 2)

	// 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
	// 因为用户1已经不再关注用户2.
	if got := twitter.GetNewsFeed(1); !reflect.DeepEqual(got, []int{5}) {
		t.Errorf("GetNewsFeed() = %v, want %v", got, []int{5})
	}
}

func testEx2(twitter TwitterInterface, t *testing.T) {
	twitter.PostTweet(1, 5)
	twitter.PostTweet(1, 3)

	// 用户1的获取推文应当返回一个列表，其中包含一个id为5的推文.
	if got := twitter.GetNewsFeed(1); !reflect.DeepEqual(got, []int{3, 5}) {
		t.Errorf("GetNewsFeed() = %v, want %v", got, []int{3, 5})
	}
}
