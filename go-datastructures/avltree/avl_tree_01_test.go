package avltree

import (
	"fmt"
	avl "github.com/emirpasic/gods/trees/avltree"
	"testing"
)

// 平衡二叉查找树
// 平衡因子:左子树的高度减去它的右子树的高度,平衡因子可以直接存储在每个节点中，或从可能存储在节点中的子树高度计算出来。平衡因子 -2或2的节点被认为是不平衡的，并需要重新平衡这个树
// 任一节点对应的两棵子树的最大高度差为1，因此它也被称为高度平衡树
// 查找、插入和删除在平均和最坏情况下的时间复杂度都是O(log{n})
// 增加和删除元素的操作则可能需要借由一次或多次树旋转，以实现树的重新平衡

func Test_avltree_01(t *testing.T) {
	tree := avl.NewWithIntComparator() // empty(keys are of type int)

	tree.Put(1, "x") // 1->x
	tree.Put(2, "b") // 1->x, 2->b (in order)
	tree.Put(1, "a") // 1->a, 2->b (in order, replacement)
	tree.Put(3, "c") // 1->a, 2->b, 3->c (in order)
	tree.Put(4, "d") // 1->a, 2->b, 3->c, 4->d (in order)
	tree.Put(5, "e") // 1->a, 2->b, 3->c, 4->d, 5->e (in order)
	tree.Put(6, "f") // ...
	tree.Put(7, "f")
	tree.Put(8, "f")
	tree.Put(9, "f")
	tree.Put(10, "f")
	tree.Put(11, "f")
	tree.Put(12, "f")
	tree.Put(13, "f")
	tree.Put(14, "f")
	tree.Put(15, "f")

	fmt.Println(tree)
	//	│           ┌── 15
	//	│       ┌── 14
	//	│       │   └── 13
	//	│   ┌── 12
	//	│   │   │   ┌── 11
	//	│   │   └── 10
	//	│   │       └── 9
	//	└── 8
	//			│       ┌── 7
	//			│   ┌── 6
	//			│   │   └── 5
	//			└── 4
	//					│   ┌── 3
	//					└── 2
	//							└── 1

	_ = tree.Values() // []interface {}{a b c d e f f f f f f f f f f}  (in order)
	_ = tree.Keys()   // []interface {}{1 2 3 4 5 6 7 8 9 10 11 12 13 14 15} (in order)
	if v, ok := tree.Get(12); ok {
		fmt.Println(v) // f
	}
	tree.Remove(2)

	fmt.Println(tree)
	//AVLTree
	//│           ┌── 15
	//│       ┌── 14
	//│       │   └── 13
	//│   ┌── 12
	//│   │   │   ┌── 11
	//│   │   └── 10
	//│   │       └── 9
	//└── 8
	//    │       ┌── 7
	//    │   ┌── 6
	//    │   │   └── 5
	//    └── 4
	//        └── 3
	//            └── 1

	tree.Clear() // empty
	tree.Empty() // true
	tree.Size()  // 0
}
