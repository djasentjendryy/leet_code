package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// ListNode matches LeetCode's definition.
type ListNode struct {
	Val  int
	Next *ListNode
}

// -------------------------------------------------------------------
// TODO: Paste YOUR solution here (do NOT change the signature).
//
// func mergeNodes(head *ListNode) *ListNode {
//     ...
// }
//
// -------------------------------------------------------------------

// Temporary stub so this file compiles before you paste your solution.
// Delete this stub after you paste your own mergeNodes.
func mergeNodes(head *ListNode) *ListNode {
	nums := 0
	dummy := &ListNode{
		Val:  -1,
		Next: nil,
	}
	prev := &ListNode{
		Val:  -1,
		Next: nil,
	}
	dummy.Next = prev
	for head != nil {
		currVal := head.Val
		if currVal == 0 {
			if nums != 0 {
				prev = createNewNode(nums, prev)
				nums = 0
			}
			head = head.Next
		} else {
			nums += head.Val
			head = head.Next
		}
	}
	return dummy.Next.Next
}

func createNewNode(val int, prev *ListNode) *ListNode {
	newNode := &ListNode{
		Val:  val,
		Next: nil,
	}

	prev.Next = newNode
	return newNode
}

// --- Test Runner ---

type testCase struct {
	name     string
	input    []int
	expected []int
}

func runCase(tc testCase) (ok bool, got []int, panicked bool, panicMsg any) {
	defer func() {
		if r := recover(); r != nil {
			panicked = true
			panicMsg = r
			ok = false
		}
	}()

	head := buildList(tc.input)
	out := mergeNodes(head)
	got = listToSlice(out)

	ok = reflect.DeepEqual(got, tc.expected)
	return ok, got, false, nil
}

func main() {
	// Opt-in stdin mode so IDEs (GoLand) won't accidentally block on stdin.
	// Usage:
	//   go run main.go                 -> runs built-in tests
	//   echo "[0,3,1,0,4,5,2,0]" | go run main.go -stdin
	//   go run main.go -stdin < input.txt
	useStdin := flag.Bool("stdin", false, "read input list from stdin")
	flag.Parse()

	if *useStdin {
		b, _ := io.ReadAll(os.Stdin)
		line := strings.TrimSpace(string(b))
		if line == "" {
			fmt.Println("No stdin input detected.")
			return
		}
		vals, err := parseInputList(line)
		if err != nil {
			fmt.Printf("Input parse error: %v\n", err)
			return
		}
		head := buildList(vals)
		fmt.Printf("Input:  %v\n", vals)
		fmt.Printf("Output: %s\n", formatList(mergeNodes(head)))
		return
	}

	// Built-in tests
	tests := []testCase{
		{
			name:     "Sample_1",
			input:    []int{0, 3, 1, 0, 4, 5, 2, 0},
			expected: []int{4, 11},
		},
		{
			name:     "Single_segment",
			input:    []int{0, 7, 0},
			expected: []int{7},
		},
		{
			name:     "Many_segments",
			input:    []int{0, 1, 2, 3, 0, 4, 0, 5, 6, 0},
			expected: []int{6, 4, 11},
		},
		{
			name:     "Zeros_at_ends_small",
			input:    []int{0, 1, 0},
			expected: []int{1},
		},
		{
			name:     "Larger_values",
			input:    []int{0, 1000, 2000, 3000, 0, 4000, 0},
			expected: []int{6000, 4000},
		},
		{
			name:     "Alternating_short_segments",
			input:    []int{0, 9, 0, 8, 0, 7, 0},
			expected: []int{9, 8, 7},
		},
		{
			name:     "Longer_mixed",
			input:    []int{0, 2, 2, 2, 0, 1, 3, 5, 0, 10, 0},
			expected: []int{6, 9, 10},
		},
	}

	pass := 0
	for _, tc := range tests {
		ok, got, panicked, msg := runCase(tc)
		if panicked {
			fmt.Printf("❌ %s: PANIC: %v\n", tc.name, msg)
			fmt.Printf("   input:    %v\n", tc.input)
			fmt.Printf("   expected: %v\n", tc.expected)
			continue
		}
		if !ok {
			fmt.Printf("❌ %s: FAIL\n", tc.name)
			fmt.Printf("   input:    %v\n", tc.input)
			fmt.Printf("   expected: %v\n", tc.expected)
			fmt.Printf("   got:      %v\n", got)
			continue
		}
		fmt.Printf("✅ %s: PASS\n", tc.name)
		pass++
	}

	fmt.Printf("\nResult: %d/%d tests passed\n", pass, len(tests))
}

// --- Helpers ---

func buildList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	dummy := &ListNode{}
	cur := dummy
	for _, v := range vals {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func listToSlice(head *ListNode) []int {
	out := make([]int, 0)
	for cur := head; cur != nil; cur = cur.Next {
		out = append(out, cur.Val)
	}
	return out
}

func formatList(head *ListNode) string {
	vals := listToSlice(head)
	return fmt.Sprintf("%v", vals)
}

// parseInputList supports:
//   - JSON-like: [0,3,1,0,4,5,2,0]
//   - CSV-like: 0,3,1,0,4,5,2,0
//   - space-separated: 0 3 1 0 4 5 2 0
func parseInputList(s string) ([]int, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil, fmt.Errorf("empty input")
	}

	// JSON array format
	if strings.HasPrefix(s, "[") {
		var arr []int
		if err := json.Unmarshal([]byte(s), &arr); err != nil {
			return nil, fmt.Errorf("failed to parse JSON array: %w", err)
		}
		return arr, nil
	}

	// Otherwise: split by comma/space
	fields := strings.FieldsFunc(s, func(r rune) bool {
		return r == ',' || r == ' ' || r == '\t' || r == '\n' || r == '\r'
	})
	if len(fields) == 0 {
		return nil, fmt.Errorf("no numbers found")
	}
	arr := make([]int, 0, len(fields))
	for _, f := range fields {
		if f == "" {
			continue
		}
		n, err := strconv.Atoi(f)
		if err != nil {
			return nil, fmt.Errorf("bad integer %q: %w", f, err)
		}
		arr = append(arr, n)
	}
	return arr, nil
}
