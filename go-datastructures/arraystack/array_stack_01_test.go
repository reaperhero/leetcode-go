package array_stack_test

import (
	"github.com/emirpasic/gods/stacks/arraystack"
	"testing"
)

// 数组-栈
func Test_array_stack_01(t *testing.T) {
	stack := arraystack.New()  // empty
	stack.Push(1)       // 1
	stack.Push(2)       // 1, 2
	stack.Values()            // 2, 1 (LIFO order)
	_, _ = stack.Peek()       // 2,true
	_, _ = stack.Pop()        // 2, true
	_, _ = stack.Pop()        // 1, true
	_, _ = stack.Pop()        // nil, false (nothing to pop)
	stack.Push(1)       // 1
	stack.Clear()             // empty
	stack.Empty()             // true
	stack.Size()              // 0
}




