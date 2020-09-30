# Go data Structures


Various examples on how to use data structures.

- [ArrayList]

ArrayList 是一个数组队列,相当于动态数组,提供了相关的添加、删除、修改、遍历等功能

- [ArrayStack]

使用数组来实现的栈，是一种后进先出(LIFO)的数据结构

- [LinkedStack]

使用链表来实现的栈，是一种后进先出(LIFO)的数据结构

- [AVLTree]

平衡二叉搜索树是二叉搜索树的升级版本。它拥有二叉搜索树的所有特性，同时它对子树的高度也进行了一定的限制。所以平衡二叉搜索树不会退化成链表的形式，而是维持了一个比较平衡的二叉树状态.在元素插入或删除后，如果树的平衡性被破坏了还要进行平衡性的修复。


- [BinaryHeap]

一种特殊的堆，二叉堆是完全二叉树或者是近似完全二叉树。二叉堆一般用数组来表示。如果根节点在数组中的位置是1，第n个位置的子节点分别在2n和 2n+1。
对于最大堆，删除根节点就是删除最大值；对于最小堆，是删除最小值。然后，把堆存储的最后那个节点移到填在根节点处。再从上而下调整父节点与它的子节点：对于最大堆，父节点如果小于具有最大值的子节点，则交换二者。
这一操作称作down-heap或bubble-down, percolate-down, sift-down, trickle down, heapify-down, cascade-down,extract-min/max等。直至当前节点与它的子节点满足“堆性质”为止。

- [BTree]

平衡多路查找树是为磁盘等外存储设备设计的一种树，该树是AVLTree的改进版本。B-Tree相对于AVLTree缩减了节点个数，使每次磁盘I/O取到内存的数据都发挥了作用，从而提高了查询效率。
系统从磁盘读取数据到内存时是以磁盘块（block）为基本单位的，位于同一个磁盘块中的数据会被一次性读取出来，B-Tree结构的数据可以让系统高效的找到数据所在的磁盘块

- [B+Tree]

B-Tree每个节点中不仅包含数据的key值，还有data值。而每一个页的存储空间是有限的，如果data数据较大时将会导致每个节点（即一个页）能存储的key的数量很小，当存储的数据量很大时同样会导致B-Tree的深度较大，增大查询时的磁盘I/O次数，进而影响查询效率。
在B+Tree中，所有数据记录节点都是按照键值大小顺序存放在同一层的叶子节点上，而非叶子节点上只存储key值信息，这样可以大大加大每个节点存储的key值数量，降低B+Tree的高度。

- [Custom Comparator](https://github.com/emirpasic/gods/blob/master/examples/customcomparator/customcomparator.go)

- [DoublyLinkedList]

双向链表的结点包括三个部分：前驱指针域、数据域和后继指针域

- [EnumerableWithIndex](https://github.com/emirpasic/gods/blob/master/examples/enumerablewithindex/enumerablewithindex.go)
- [EnumerableWithKey](https://github.com/emirpasic/gods/blob/master/examples/enumerablewithkey/enumerablewithkey.go)
- [HashBidiMap](https://github.com/emirpasic/gods/blob/master/examples/hashbidimap/hashbidimap.go)

- [HashMap]

不允许键值对集合中有重复的键

- [HashSet]

不允许集合中有重复的值

- [IteratorWithIndex](https://github.com/emirpasic/gods/blob/master/examples/iteratorwithindex/iteratorwithindex.go)
- [iteratorwithkey](https://github.com/emirpasic/gods/blob/master/examples/iteratorwithkey/iteratorwithkey.go)
- [IteratorWithKey](https://github.com/emirpasic/gods/blob/master/examples/linkedliststack/linkedliststack.go)
- [RedBlackTree](https://github.com/emirpasic/gods/blob/master/examples/redblacktree/redblacktree.go)
- [RedBlackTreeExtended](https://github.com/emirpasic/gods/blob/master/examples/redblacktreeextended/redblacktreeextended.go)
- [Serialization](https://github.com/emirpasic/gods/blob/master/examples/serialization/serialization.go)
- [SinglyLinkedList](https://github.com/emirpasic/gods/blob/master/examples/singlylinkedlist/singlylinkedlist.go)
- [Sort](https://github.com/emirpasic/gods/blob/master/examples/sort/sort.go)
- [TreeBidiMap](https://github.com/emirpasic/gods/blob/master/examples/treebidimap/treebidimap.go)
- [TreeMap]

HashMap中根据key值来进行排序，排序方法可以自己定义

- [TreeSet]

保存了对象的排列次序，排序方法可以自己定义