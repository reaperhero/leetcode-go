[![Travis-ci Status](https://travis-ci.org/cizixs/go-algorithms.svg?branch=master)](https://travis-ci.org/cizixs/go-algorithms)
[![Coverage Status](https://coveralls.io/repos/github/cizixs/go-algorithms/badge.svg)](https://coveralls.io/github/cizixs/go-algorithms)

# go-algorithms

- Data Structures
    - [x] Queue
    - [x] List
    - [X] Stack
    - [x] Deque
    - [x] Binary Tree
    - [x] AVL tree
    - [x] B-Trees
    - [x] Hash
    - [x] Priority Queue
    - [x] Set
    - [x] Heap

- Algorithms
    - [x] Sort
    
    
    
## 二叉树


- 满二叉树

除最后一层无任何子节点外，每一层上的所有结点都有两个子结点二叉树

- 完全二叉树

一棵二叉树至多只有最下面的一层上的结点的度数可以小于2，并且最下层上的结点都集中在该层最左边的若干位置上，则此二叉树成为完全二叉树。

- 平衡二叉树

它是一棵空树或它的左右两个子树的高度差的绝对值不超过1，并且左右两个子树都是一棵平衡二叉树

- 二叉搜索树

它或者是一棵空树，或者是具有下列性质的二叉树： 若它的左子树不空，则左子树上所有结点的值均小于它的根结点的值； 若它的右子树不空，则右子树上所有结点的值均大于它的根结点的值； 它的左、右子树也分别为二叉排序树

- 红黑树

平衡二叉搜索树

- 二叉堆

一种特殊的堆，二叉堆是完全二叉树或者是近似完全二叉树。二叉堆一般用数组来表示。如果根节点在数组中的位置是1，第n个位置的子节点分别在2n和 2n+1。
对于最大堆，删除根节点就是删除最大值；对于最小堆，是删除最小值。然后，把堆存储的最后那个节点移到填在根节点处。再从上而下调整父节点与它的子节点：对于最大堆，父节点如果小于具有最大值的子节点，则交换二者。
这一操作称作down-heap或bubble-down, percolate-down, sift-down, trickle down, heapify-down, cascade-down,extract-min/max等。直至当前节点与它的子节点满足“堆性质”为止。

- B树

B树也称B-树,它是一颗多路平衡查找树。我们描述一颗B树时需要指定它的阶数，阶数表示了一个结点最多有多少个孩子结点，一般用字母m表示阶数。当m取2时，就是我们常见的二叉搜索树

- B+树

B+树与B树最大的不同是内部结点不保存数据，只用于索引，所有数据（或者说记录）都保存在叶子结点中