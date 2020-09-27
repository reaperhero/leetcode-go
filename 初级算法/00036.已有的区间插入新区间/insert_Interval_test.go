package _0036_已有的区间插入新区间

import (
	"fmt"
	structures "leetcode-go/数据结构"
	"testing"
)

// 可以分 3 段处理，先添加原来的区间，即在给的 newInterval 之前的区间。然后添加 newInterval ，注意这里可能需要合并多个区间。最后把原来剩下的部分添加到最终结果中即可。
// Input: intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
// Output: [[1,2],[3,10],[12,16]]
// Explanation: Because the new interval [4,8] overlaps with [3,5],[6,7],[8,10].

type Interval = structures.Interval

func insert(intervals []Interval, newInterval Interval) []Interval {
	res := make([]Interval, 0)
	if len(intervals) == 0 {
		res = append(res, newInterval)
		return res
	}
	curIndex := 0
	for curIndex < len(intervals) && intervals[curIndex].End < newInterval.Start {
		res = append(res, intervals[curIndex])
		curIndex++
	}

	for curIndex < len(intervals) && intervals[curIndex].Start <= newInterval.End {
		newInterval = Interval{Start: min(newInterval.Start, intervals[curIndex].Start), End: max(newInterval.End, intervals[curIndex].End)}
		curIndex++
	}
	res = append(res, newInterval)

	for curIndex < len(intervals) {
		res = append(res, intervals[curIndex])
		curIndex++
	}
	return res
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

func Test_insert(t *testing.T) {
	intervals := []Interval{{Start: 1, End: 2}, {Start: 3, End: 5}, {Start: 6, End: 7}, {Start: 8, End: 10}, {Start: 12, End: 16}}
	newInterval := Interval{Start: 4, End: 8}
	fmt.Println(insert(intervals, newInterval)) // [{1 2} {3 10} {12 16}]

}
