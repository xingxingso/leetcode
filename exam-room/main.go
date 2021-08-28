/*
Package exam_room
https://leetcode-cn.com/problems/exam-room/

855. 考场就座

在考场里，一排有 N 个座位，分别编号为 0, 1, 2, ..., N-1 。

当学生进入考场后，他必须坐在能够使他与离他最近的人之间的距离达到最大化的座位上。
如果有多个这样的座位，他会坐在编号最小的座位上。(另外，如果考场里没有人，那么学生就坐在 0 号座位上。)

返回 ExamRoom(int N) 类，它有两个公开的函数：其中，函数 ExamRoom.seat() 会返回一个 int （整型数据），代表学生坐的位置；
函数 ExamRoom.leave(int p) 代表坐在座位 p 上的学生现在离开了考场。每次调用 ExamRoom.leave(p) 时都保证有学生坐在座位 p 上。

提示：
	1. 1 <= N <= 10^9
	2. 在所有的测试样例中 ExamRoom.seat() 和 ExamRoom.leave() 最多被调用 10^4 次。
	3. 保证在调用 ExamRoom.leave(p) 时有学生正坐在座位 p 上。

*/
package exam_room

//type ExamRoom struct {}
//func Constructor(n int) ExamRoom {}
//func (this *ExamRoom) Seat() int {}
//func (this *ExamRoom) Leave(p int) {}

/**
 * Your ExamRoom object will be instantiated and called as such:
 * obj := Constructor(n);
 * param_1 := obj.Seat();
 * obj.Leave(p);
 */

// --- 自己

/*
ExamRoomS1
方法一：

时间复杂度：
空间复杂度：
*/
type ExamRoomS1 struct {
	max   int
	seats []int
	diffs []int
}

func ConstructorS1(n int) ExamRoomS1 {
	return ExamRoomS1{
		max:   n - 1,
		seats: []int{},
		diffs: []int{n - 1},
	}
}

func (this *ExamRoomS1) Seat() int {
	if len(this.seats) == 0 {
		this.seat(0)
		return 0
	}

	maxDiffIndex := 0
	for i := 1; i < len(this.diffs); i++ {
		if this.diffs[i] > this.diffs[maxDiffIndex] {
			maxDiffIndex = i
		}
	}
	idx := -1
	if maxDiffIndex == 0 {
		idx = 0
	} else if maxDiffIndex == len(this.diffs)-1 {
		idx = this.max
	} else {
		idx = this.seats[maxDiffIndex-1] + (this.seats[maxDiffIndex]-this.seats[maxDiffIndex-1])>>1
	}
	this.seat(idx)
	return idx
}

func (this *ExamRoomS1) seat(p int) {
	newSeats := make([]int, len(this.seats)+1)
	pdx := -1
	for i := 0; i < len(newSeats); i++ {
		if i < len(this.seats) && this.seats[i] < p {
			newSeats[i] = this.seats[i]
		} else if pdx == -1 {
			pdx = i
			newSeats[i] = p
		} else {
			newSeats[i] = this.seats[i-1]
		}
	}
	this.seats = newSeats
	this.calDiff()
}

func (this *ExamRoomS1) calDiff() {
	newDiffs := make([]int, len(this.seats)+1)
	for i := 0; i < len(this.seats); i++ {
		if i > 0 {
			newDiffs[i] = (this.seats[i] - this.seats[i-1]) / 2
		}
	}
	newDiffs[0], newDiffs[len(newDiffs)-1] = this.max, this.max
	if len(this.seats) > 0 {
		newDiffs[0] = this.seats[0] - 0
		newDiffs[len(newDiffs)-1] = this.max - this.seats[len(this.seats)-1]
	}
	this.diffs = newDiffs
}

func (this *ExamRoomS1) Leave(p int) {
	if len(this.seats) == 1 {
		this.seats = []int{}
		this.diffs = []int{this.max}
		return
	}
	newSeats := make([]int, len(this.seats)-1)
	for i := 0; i < len(newSeats); i++ {
		if this.seats[i] < p {
			newSeats[i] = this.seats[i]
		} else {
			newSeats[i] = this.seats[i+1]
		}
	}
	this.seats = newSeats
	this.calDiff()
}
