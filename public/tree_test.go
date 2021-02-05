package public

import (
	"reflect"
	"testing"
)

func Test_createTreeByArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "equal",
			args: args{
				nums: []int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 0, 1},
			},

			want: want11(),
		},
		{
			name: "equal",
			args: args{
				nums: []int{3, 9, 8, 4, 6, 1, 7, 0, 0, 0, 2, 5},
			},

			want: want12(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CreateTreeByArray(tt.args.nums); !isSameTree(got, tt.want) {
				t.Errorf("createTreeByArray() = %v, want %v", got, tt.want)
			}
		})
	}
}

func want12() *TreeNode {
	node1 := &TreeNode{
		Val: 2,
	}
	node2 := &TreeNode{
		Val: 5,
	}

	node3 := &TreeNode{
		Val: 4,
	}

	node4 := &TreeNode{
		Val:   6,
		Right: node1,
	}

	node5 := &TreeNode{
		Val:  1,
		Left: node2,
	}

	node6 := &TreeNode{
		Val: 7,
	}

	node7 := &TreeNode{
		Val:   9,
		Left:  node3,
		Right: node4,
	}

	node8 := &TreeNode{
		Val:   8,
		Left:  node5,
		Right: node6,
	}

	node9 := &TreeNode{
		Val:   3,
		Left:  node7,
		Right: node8,
	}

	return node9
}

func want11() *TreeNode {
	node1 := &TreeNode{
		Val: 7,
	}
	node2 := &TreeNode{
		Val: 2,
	}

	node3 := &TreeNode{
		Val: 1,
	}

	node4 := &TreeNode{
		Val:   11,
		Left:  node1,
		Right: node2,
	}

	node5 := &TreeNode{
		Val:   13,
		Left:  nil,
		Right: nil,
	}

	node6 := &TreeNode{
		Val:   4,
		Left:  nil,
		Right: node3,
	}

	node7 := &TreeNode{
		Val:   4,
		Left:  node4,
		Right: nil,
	}

	node8 := &TreeNode{
		Val:   8,
		Left:  node5,
		Right: node6,
	}

	node9 := &TreeNode{
		Val:   5,
		Left:  node7,
		Right: node8,
	}

	return node9
}

func Test_bfs(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "equal",
			args: args{
				root: want11(),
			},

			want: []int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 0, 1},
		},
		{
			name: "equal",
			args: args{
				root: CreateTreeByArray([]int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 0, 1}),
			},

			want: []int{5, 4, 8, 11, 0, 13, 4, 7, 2, 0, 0, 0, 1},
		},
		{
			name: "equal",
			args: args{
				root: CreateTreeByArray([]int{3, 9, 8, 4, 6, 1, 7, 0, 0, 0, 2, 5}),
			},

			want: []int{3, 9, 8, 4, 6, 1, 7, 0, 0, 0, 2, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bfs(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bfs() = %v, want %v", got, tt.want)
			}
		})
	}
}
