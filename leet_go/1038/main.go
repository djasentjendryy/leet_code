package main

import (
	"fmt"
	"math"
	"reflect"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TODO: Implement your solution here (do NOT change signature).
func bstToGst(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}

	biggestVal := getBiggestVal(root)

	return assignVal(root, biggestVal)
}

func getBiggestVal(root *TreeNode) int {
	if root == nil {
		return 0
	}

	rightVal := getBiggestVal(root.Right)
	leftVal := getBiggestVal(root.Left)

	return int(math.Max(math.Max(float64(rightVal), float64(leftVal)), float64(root.Val)))
}

func assignVal(root *TreeNode, biggestVal int) *TreeNode {
	if root == nil {
		return root
	}

	assignVal(root.Left, biggestVal)
	assignVal(root.Right, biggestVal)

	result := 0
	for i := root.Val; i <= biggestVal; i++ {
		result += i
	}
	root.Val = result

	return root
}

func main() {
	tests := []TestCase{
		{
			name: "LeetCode Example 1",
			input: []interface{}{
				4, 1, 6, 0, 2, 5, 7, nil, nil, nil, 3, nil, nil, nil, 8,
			},
			expected: []interface{}{
				30, 36, 21, 36, 35, 26, 15, nil, nil, nil, 33, nil, nil, nil, 8,
			},
		},
		{
			name:     "LeetCode Example 2",
			input:    []interface{}{0, nil, 1},
			expected: []interface{}{1, nil, 1},
		},
		{
			name:     "Single node",
			input:    []interface{}{5},
			expected: []interface{}{5},
		},
		{
			name:     "Small balanced",
			input:    []interface{}{2, 1, 3},
			expected: []interface{}{5, 6, 3},
		},
		{
			name:     "Another small BST",
			input:    []interface{}{1, 0, 2},
			expected: []interface{}{3, 3, 2},
		},
		{
			name:     "Right skewed",
			input:    []interface{}{0, nil, 1, nil, 2, nil, 3},
			expected: []interface{}{6, nil, 6, nil, 5, nil, 3},
		},
		{
			name:     "Left skewed",
			input:    []interface{}{3, 2, nil, 1},
			expected: []interface{}{3, 5, nil, 6},
		},
		{
			name:     "Empty tree",
			input:    []interface{}{},
			expected: []interface{}{},
		},
	}

	passed := 0
	for _, tc := range tests {
		root := buildTree(tc.input)
		gotRoot := bstToGst(root)
		got := treeToLevelOrder(gotRoot)

		if reflect.DeepEqual(got, tc.expected) {
			fmt.Printf("✅ PASS: %s\n", tc.name)
			passed++
		} else {
			fmt.Printf("❌ FAIL: %s\n", tc.name)
			fmt.Printf("   expected: %v\n", tc.expected)
			fmt.Printf("   got:      %v\n", got)
		}
	}

	fmt.Printf("\n%d/%d tests passed\n", passed, len(tests))
}

type TestCase struct {
	name     string
	input    []interface{}
	expected []interface{}
}

// buildTree builds a binary tree from level-order representation.
// Example: []interface{}{4, 1, 6, 0, 2, 5, 7, nil, nil, nil, 3, nil, nil, nil, 8}
func buildTree(level []interface{}) *TreeNode {
	if len(level) == 0 || level[0] == nil {
		return nil
	}

	root := &TreeNode{Val: level[0].(int)}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(level) {
		node := queue[0]
		queue = queue[1:]

		// Left child
		if i < len(level) && level[i] != nil {
			node.Left = &TreeNode{Val: level[i].(int)}
			queue = append(queue, node.Left)
		}
		i++

		// Right child
		if i < len(level) && level[i] != nil {
			node.Right = &TreeNode{Val: level[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

// treeToLevelOrder serializes a binary tree into level-order with nil placeholders,
// trimming trailing nils for stable comparisons.
func treeToLevelOrder(root *TreeNode) []interface{} {
	if root == nil {
		return []interface{}{}
	}

	res := make([]interface{}, 0)
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node == nil {
			res = append(res, nil)
			continue
		}

		res = append(res, node.Val)
		queue = append(queue, node.Left, node.Right)
	}

	// Trim trailing nils
	for len(res) > 0 && res[len(res)-1] == nil {
		res = res[:len(res)-1]
	}

	return res
}
