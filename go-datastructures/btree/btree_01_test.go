package btree

import (
	"fmt"
	"github.com/emirpasic/gods/trees/btree"
	"testing"
)

func Test_tree_01(t *testing.T) {
	tree := btree.NewWithIntComparator(3) // empty (keys are of type int)

	tree.Put(1, "x") // 1->x
	tree.Put(2, "b") // 1->x, 2->b (in order)
	tree.Put(1, "a") // 1->a, 2->b (in order, replacement)
	tree.Put(3, "c") // 1->a, 2->b, 3->c (in order)
	tree.Put(4, "d") // 1->a, 2->b, 3->c, 4->d (in order)
	tree.Put(5, "e") // 1->a, 2->b, 3->c, 4->d, 5->e (in order)
	tree.Put(6, "f") // 1->a, 2->b, 3->c, 4->d, 5->e, 6->f (in order)
	tree.Put(7, "g") // 1->a, 2->b, 3->c, 4->d, 5->e, 6->f, 7->g (in order)
	iterator := tree.Iterator()
	for iterator.Next() { // 根据key排序，依次打印key value
		fmt.Println(iterator.Key())
		fmt.Println(iterator.Value())
	}
	//fmt.Println(tree)
	// BTree
	//         1
	//     2
	//         3
	// 4
	//         5
	//     6
	//         7

	//fmt.Println(tree.RightValue()) // 获取最右边的value，相当于最大key的value
	//fmt.Println(tree.RightKey())   // 获取最右边的key，相当于最大key
	//fmt.Println(tree.LeftKey()) // 获取最左边的value，相当于最小key的value
	//fmt.Println(tree.LeftValue())   // 获取最左边的key，相当于最小key
	_ = tree.Values() // []interface {}{"a", "b", "c", "d", "e", "f", "g"} (in order)
	_ = tree.Keys()   // []interface {}{1, 2, 3, 4, 5, 6, 7} (in order)

	//for _, v := range tree.Keys() {
	//	fmt.Println(v) // 从最小key开始打印
	//}
	//for _, v := range tree.Values() {
	//	fmt.Println(v) // 根据key的排序，依次打印value
	//}
	tree.Remove(2) // 1->a, 3->c, 4->d, 5->e, 6->f (in order)
	//fmt.Println(tree)
	//BTree
	//	1
	//	3
	//4
	//	5
	//6
	//	7

	tree.Clear() // empty
	tree.Empty() // true
	tree.Size()  // 0

	// Other:
	tree.Height()     // gets the height of the tree
	tree.Left()       // gets the left-most (min) node
	tree.Right()      // get the right-most (max) node
}
