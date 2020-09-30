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

它是一棵空树或它的左右两个子树的高度差的绝对值不超过1，1 又称平衡因子并且左右两个子树都是一棵平衡二叉树

- 二叉搜索树

若它的左子树不空，则左子树上所有结点的值均小于它的根结点的值； 
若它的右子树不空，则右子树上所有结点的值均大于它的根结点的值；
AVL树适合用于插入与删除次数比较少，但查找多的情况

- 红黑树

平衡二叉搜索树,节点是红色或黑色,根是黑色，叶子都是黑色，每个红色节点必须有两个黑色的子节点，从每个叶子到根的所有路径上不能有两个连续的红色节点，从任一节点到其每个叶子的所有简单路径都包含相同数目的黑色节点
红黑树确保没有一条路径会比其它路径长出两倍，因此，红黑树是一种弱平衡二叉树,在相同的节点情况下，AVL树的高度低于红黑树,相对于要求严格的AVL树来说，它的旋转次数少，所以对于搜索，插入，删除操作较多的情况下，我们就用红黑树。
因为每一个红黑树也是一个特化的二叉查找树，因此红黑树上的只读操作与普通二叉查找树上的只读操作相同。
然而，在红黑树上进行插入操作和删除操作会导致不再符合红黑树的性质。恢复红黑树的性质需要少量O(log n)的颜色变更（实际是非常快速的）和不超过三次树旋转（对于插入操作是两次）。虽然插入和删除很复杂，但操作时间仍可以保持为O(log n)次。

- B树

B树也称B-树,它是一颗多路平衡查找树。我们描述一颗B树时需要指定它的阶数，阶数表示了一个结点最多有多少个孩子结点，一般用字母m表示阶数。当m取2时，就是我们常见的二叉搜索树

- B+树

B+树与B树最大的不同是内部结点不保存数据，只用于索引，所有数据（或者说记录）都保存在叶子结点中