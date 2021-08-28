package exam_room

import (
	"fmt"
	"testing"
)

func TestExamRoomEx1(t *testing.T) {
	room := ConstructorS1(10)
	if got := room.Seat(); got != 0 {
		t.Errorf("Seat(), want 0, got %d", got)
	}
	if got := room.Seat(); got != 9 {
		t.Errorf("Seat(), want 9, got %d", got)
	}
	if got := room.Seat(); got != 4 {
		t.Errorf("Seat(), want 4, got %d", got)
	}
	fmt.Println(room.seats, room.diffs)
	if got := room.Seat(); got != 2 {
		t.Errorf("Seat(), want 2, got %d", got)
	}
	fmt.Println(room.seats, room.diffs)
	room.Leave(4)
	if got := room.Seat(); got != 5 {
		t.Errorf("Seat(), want 5, got %d", got)
	}
}

func TestExamRoomEx2(t *testing.T) {
	room := ConstructorS1(4)
	if got := room.Seat(); got != 0 {
		t.Errorf("Seat(), want 0, got %d", got)
	}
	if got := room.Seat(); got != 3 {
		t.Errorf("Seat(), want 3, got %d", got)
	}
	if got := room.Seat(); got != 1 {
		t.Errorf("Seat(), want 1, got %d", got)
	}
	if got := room.Seat(); got != 2 {
		t.Errorf("Seat(), want 2, got %d", got)
	}
	room.Leave(1)
	room.Leave(3)
	if got := room.Seat(); got != 1 {
		t.Errorf("Seat(), want 1, got %d", got)
	}
}

// ["ExamRoom","seat","seat","seat","leave","leave","seat","seat","seat","seat","seat","seat","seat","seat","seat","leave"]
// [[10],[],[],[],[0],[4],[],[],[],[],[],[],[],[],[],[0]]
// [null,0,9,4,null,null,0,4,2,6,1,3,5,7,8,null]
func TestExamRoomEx3(t *testing.T) {
	room := ConstructorS1(10)
	if got := room.Seat(); got != 0 {
		t.Errorf("Seat(), want 0, got %d", got)
	}
	if got := room.Seat(); got != 9 {
		t.Errorf("Seat(), want 9, got %d", got)
	}
	if got := room.Seat(); got != 4 {
		t.Errorf("Seat(), want 4, got %d", got)
	}
	room.Leave(0)
	room.Leave(4)
	if got := room.Seat(); got != 0 {
		t.Errorf("Seat(), want 0, got %d", got)
	}
	if got := room.Seat(); got != 4 {
		t.Errorf("Seat(), want 4, got %d", got)
	}
	if got := room.Seat(); got != 2 {
		t.Errorf("Seat(), want 2, got %d", got)
	}
	if got := room.Seat(); got != 6 {
		t.Errorf("Seat(), want 6, got %d", got)
	}
	if got := room.Seat(); got != 1 {
		t.Errorf("Seat(), want 1, got %d", got)
	}
	if got := room.Seat(); got != 3 {
		t.Errorf("Seat(), want 3, got %d", got)
	}
	if got := room.Seat(); got != 5 {
		t.Errorf("Seat(), want 5, got %d", got)
	}
	if got := room.Seat(); got != 7 {
		t.Errorf("Seat(), want 7, got %d", got)
	}
	if got := room.Seat(); got != 8 {
		t.Errorf("Seat(), want 8, got %d", got)
	}
	room.Leave(0)
}

//["ExamRoom","seat","seat","seat","leave","leave","seat","seat","seat","seat","seat","seat","seat","seat","seat","leave","leave","seat","seat","leave","seat","leave","seat","leave","seat","leave","seat","leave","leave","seat","seat","leave","leave","seat","seat","leave"]
//[[10],[],[],[],[0],[4],[],[],[],[],[],[],[],[],[],[0],[4],[],[],[7],[],[3],[],[3],[],[9],[],[0],[8],[],[],[0],[8],[],[],[2]]
//[null,0,9,4,null,null,0,4,2,6,1,3,5,7,8,null,null,0,4,null,7,null,3,null,3,null,9,null,null,0,8,null,null,0,8,null]
func TestExamRoomEx4(t *testing.T) {
	room := ConstructorS1(10)
	if got := room.Seat(); got != 0 {
		t.Errorf("Seat(), want 0, got %d", got)
	}
	if got := room.Seat(); got != 9 {
		t.Errorf("Seat(), want 9, got %d", got)
	}
	if got := room.Seat(); got != 4 {
		t.Errorf("Seat(), want 4, got %d", got)
	}
	room.Leave(0)
	room.Leave(4)
	if got := room.Seat(); got != 0 {
		t.Errorf("Seat(), want 0, got %d", got)
	}
	if got := room.Seat(); got != 4 {
		t.Errorf("Seat(), want 4, got %d", got)
	}
	if got := room.Seat(); got != 2 {
		t.Errorf("Seat(), want 2, got %d", got)
	}
	if got := room.Seat(); got != 6 {
		t.Errorf("Seat(), want 6, got %d", got)
	}
	if got := room.Seat(); got != 1 {
		t.Errorf("Seat(), want 1, got %d", got)
	}
	if got := room.Seat(); got != 3 {
		t.Errorf("Seat(), want 3, got %d", got)
	}
	if got := room.Seat(); got != 5 {
		t.Errorf("Seat(), want 5, got %d", got)
	}
	if got := room.Seat(); got != 7 {
		t.Errorf("Seat(), want 7, got %d", got)
	}
	if got := room.Seat(); got != 8 {
		t.Errorf("Seat(), want 8, got %d", got)
	}
	room.Leave(0)
	room.Leave(4)
	if got := room.Seat(); got != 0 {
		t.Errorf("Seat(), want 0, got %d", got)
	}
	if got := room.Seat(); got != 4 {
		t.Errorf("Seat(), want 4, got %d", got)
	}
	room.Leave(7)
	if got := room.Seat(); got != 7 {
		t.Errorf("Seat(), want 7, got %d", got)
	}
	room.Leave(3)
	if got := room.Seat(); got != 3 {
		t.Errorf("Seat(), want 3, got %d", got)
	}
	room.Leave(3)
	if got := room.Seat(); got != 3 {
		t.Errorf("Seat(), want 3, got %d", got)
	}
	room.Leave(9)
	if got := room.Seat(); got != 9 {
		t.Errorf("Seat(), want 9, got %d", got)
	}
	room.Leave(0)
	room.Leave(8)
	if got := room.Seat(); got != 0 {
		t.Errorf("Seat(), want 0, got %d", got)
	}
	if got := room.Seat(); got != 8 {
		t.Errorf("Seat(), want 8, got %d", got)
	}
	room.Leave(2)
}

//["ExamRoom","seat","seat","leave","leave","seat"]
//[[10],[],[],[0],[9],[]]
func TestExamRoomEx5(t *testing.T) {
	room := ConstructorS1(10)
	if got := room.Seat(); got != 0 {
		t.Errorf("Seat(), want 0, got %d", got)
	}
	if got := room.Seat(); got != 9 {
		t.Errorf("Seat(), want 9, got %d", got)
	}
	room.Leave(0)
	room.Leave(9)
	if got := room.Seat(); got != 0 {
		t.Errorf("Seat(), want 0, got %d", got)
	}
}

//["ExamRoom","seat","seat","seat","leave","leave","seat","seat","seat","seat","seat","seat","seat"]
//[[8],[],[],[],[0],[7],[],[],[],[],[],[],[]]
//[null,0,7,3,null,null,7,0,5,1,2,4,6]
func TestExamRoomEx6(t *testing.T) {
	room := ConstructorS1(8)
	if got := room.Seat(); got != 0 {
		t.Errorf("Seat(), want 0, got %d", got)
	}
	if got := room.Seat(); got != 7 {
		t.Errorf("Seat(), want 7, got %d", got)
	}
	if got := room.Seat(); got != 3 {
		t.Errorf("Seat(), want 3, got %d", got)
	}
	room.Leave(0)
	room.Leave(7)
	fmt.Println(room.seats, room.diffs)
	if got := room.Seat(); got != 7 {
		t.Errorf("Seat(), want 7, got %d", got)
	}
	fmt.Println(room.seats, room.diffs)
	if got := room.Seat(); got != 0 {
		t.Errorf("Seat(), want 0, got %d", got)
	}
	if got := room.Seat(); got != 5 {
		t.Errorf("Seat(), want 5, got %d", got)
	}
	if got := room.Seat(); got != 1 {
		t.Errorf("Seat(), want 1, got %d", got)
	}
	if got := room.Seat(); got != 2 {
		t.Errorf("Seat(), want 2, got %d", got)
	}
	if got := room.Seat(); got != 4 {
		t.Errorf("Seat(), want 4, got %d", got)
	}
	if got := room.Seat(); got != 6 {
		t.Errorf("Seat(), want 6, got %d", got)
	}
}
