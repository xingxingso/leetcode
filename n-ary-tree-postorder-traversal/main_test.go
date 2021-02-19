package n_ary_tree_postorder_traversal

import (
	"reflect"
	"testing"
)

func Test_postorder(t *testing.T) {
	type args struct {
		root *Node
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "equal0",
			args: args{
				root: createNAryTree(),
			},
			want: []int{5, 6, 3, 2, 4, 1},
		},
		{
			name: "equal1",
			args: args{
				root: createNAryTree2(),
			},
			want: []int{2, 6, 14, 11, 7, 3, 12, 8, 4, 13, 9, 10, 5, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := postorderS1(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorder() = %v, want %v", got, tt.want)
			}
			if got := postorderS2(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("postorder() = %v, want %v", got, tt.want)
			}
		})
	}
}

func createNAryTree() *Node {
	return &Node{
		Val: 1,
		Children: []*Node{
			&Node{
				Val: 3,
				Children: []*Node{
					&Node{
						Val: 5,
					},
					&Node{
						Val: 6,
					},
				},
			},
			&Node{
				Val: 2,
			},
			&Node{
				Val: 4,
			},
		},
	}
}

func createNAryTree2() *Node {
	return &Node{
		Val: 1,
		Children: []*Node{
			&Node{
				Val: 2,
			},
			&Node{
				Val: 3,
				Children: []*Node{
					&Node{
						Val: 6,
					},
					&Node{
						Val: 7,
						Children: []*Node{
							&Node{
								Val: 11,
								Children: []*Node{
									&Node{
										Val: 14,
									},
								},
							},
						},
					},
				},
			},
			&Node{
				Val: 4,
				Children: []*Node{
					&Node{
						Val: 8,
						Children: []*Node{
							&Node{
								Val: 12,
							},
						},
					},
				},
			},
			&Node{
				Val: 5,
				Children: []*Node{
					&Node{
						Val: 9,
						Children: []*Node{
							&Node{
								Val: 13,
							},
						},
					},
					&Node{
						Val: 10,
					},
				},
			},
		},
	}
}
