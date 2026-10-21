package main

import (
	"fmt"
)

// TreeNode is the standard LeetCode binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// averageOfSubtree is the function you need to implement.
// DO NOT change the signature.
func averageOfSubtree(root *TreeNode) int {
	result := 0
	calculateValue(root, 0, &result)
	return result
}

func calculateValue(root *TreeNode, count int, result *int) (int, int) {
	if root == nil {
		return 0, 0
	}

	valLeft, countLeft := calculateValue(root.Left, count, result)
	valRight, countRight := calculateValue(root.Right, count, result)

	totalData := countLeft + countRight + 1
	totalVal := valLeft + valRight + root.Val

	avg := totalVal / totalData
	if avg == root.Val {
		(*result)++
	}
	return totalVal, totalData
}

func main() {
	type testCase struct {
		name     string
		level    []any // level-order representation, use nil for missing nodes
		expected int
	}

	tests := []testCase{
		{
			name:     "Example 1",
			level:    []any{4, 8, 5, 0, 1, nil, 6},
			expected: 5,
		},
		//{
		//	name:     "Example 2",
		//	level:    []any{1},
		//	expected: 1,
		//},
		//
		//// Additional tests (hand-computed expected outputs):
		//{
		//	name:     "Perfect small tree",
		//	level:    []any{2, 1, 3},
		//	expected: 3,
		//},
		//{
		//	name:     "Mixed with nils",
		//	level:    []any{5, 0, 10, nil, 1},
		//	expected: 3,
		//},
		//{
		//	name:     "Left chain",
		//	level:    []any{1, 2, nil, 3, nil, 4, nil},
		//	expected: 2,
		//},
		//{
		//	name:     "All same values",
		//	level:    []any{7, 7, 7, 7, 7, 7, 7},
		//	expected: 7,
		//},
		//{
		//	name:     "Single zero",
		//	level:    []any{0},
		//	expected: 1,
		//},
	}

	fmt.Println("Running tests for averageOfSubtree(root *TreeNode) int")
	fmt.Println("----------------------------------------------------")

	passed := 0
	for i, tc := range tests {
		root := buildTreeLevelOrder(tc.level)

		var got int
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Printf("[#%d] %s: PANIC: %v\n", i+1, tc.name, r)
					fmt.Printf("      tree: %s\n", levelOrderString(root))
					fmt.Println()
				}
			}()
			got = averageOfSubtree(root)
		}()

		// If it panicked, got will be 0, but we already printed panic info.
		// Continue checking only if no panic occurred:
		if got == tc.expected {
			fmt.Printf("[#%d] %s: PASS (got %d)\n", i+1, tc.name, got)
			passed++
		} else {
			fmt.Printf("[#%d] %s: FAIL (got %d, expected %d)\n", i+1, tc.name, got, tc.expected)
			fmt.Printf("      tree: %s\n", levelOrderString(root))
		}
		fmt.Println()
	}

	fmt.Printf("Summary: %d/%d tests passed\n", passed, len(tests))
}

// buildTreeLevelOrder builds a binary tree from a level-order array representation.
// Use nil for missing nodes, e.g. []any{4, 8, 5, 0, 1, nil, 6}.
func buildTreeLevelOrder(level []any) *TreeNode {
	if len(level) == 0 || level[0] == nil {
		return nil
	}
	root := &TreeNode{Val: mustInt(level[0])}
	queue := []*TreeNode{root}

	i := 1
	for i < len(level) && len(queue) > 0 {
		cur := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(level) {
			if level[i] != nil {
				cur.Left = &TreeNode{Val: mustInt(level[i])}
				queue = append(queue, cur.Left)
			}
			i++
		}

		// Right child
		if i < len(level) {
			if level[i] != nil {
				cur.Right = &TreeNode{Val: mustInt(level[i])}
				queue = append(queue, cur.Right)
			}
			i++
		}
	}
	return root
}

func mustInt(v any) int {
	switch x := v.(type) {
	case int:
		return x
	case int32:
		return int(x)
	case int64:
		return int(x)
	case float64:
		return int(x)
	default:
		panic(fmt.Sprintf("non-int value in level array: %#v (type %T)", v, v))
	}
}

// levelOrderString prints a tree back into a compact level-order string (for debugging).
func levelOrderString(root *TreeNode) string {
	if root == nil {
		return "[]"
	}
	out := make([]any, 0)
	q := []*TreeNode{root}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		if node == nil {
			out = append(out, nil)
			continue
		}
		out = append(out, node.Val)
		q = append(q, node.Left, node.Right)
	}

	// Trim trailing nils
	for len(out) > 0 && out[len(out)-1] == nil {
		out = out[:len(out)-1]
	}

	return fmt.Sprintf("%v", out)
}
