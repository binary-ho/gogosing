package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type SplayTree struct {
	root *SplayTreeNode
}

type SplayTreeNode struct {
	key                 int64
	left, right, parent *SplayTreeNode
	size, sum, lazy     int64
}

func isNil(node *SplayTreeNode) bool {
	return node == nil
}

func isPresent(node *SplayTreeNode) bool {
	return node != nil
}

func (node *SplayTreeNode) updateTreeSizeAndSum() {
	node.size = 1
	node.sum = node.key

	if isPresent(node.left) {
		node.size += node.left.size
		node.sum += node.left.sum
	}

	if isPresent(node.right) {
		node.size += node.right.size
		node.sum += node.right.sum
	}
}

func (tree *SplayTree) Rotate(node *SplayTreeNode) {
	parent := node.parent
	if isNil(parent) {
		return
	}

	parent.pushLazyValue()
	node.pushLazyValue()

	node.setParentToChild()
	node.setGrandParentToParent()

	if isNil(node.parent) {
		tree.root = node
	}

	// update Child Count
	parent.updateTreeSizeAndSum()
	node.updateTreeSizeAndSum()
}

func (tree *SplayTree) Splay(node *SplayTreeNode) {
	if isNil(tree.root) || isNil(node) {
		//} node.isNil() || node == tree.root {
		return
	}

	if node.parent == tree.root {
		tree.Rotate(node)
		return
	}

	for isPresent(node.parent) {
		parent := node.parent
		grandParent := parent.parent

		if isPresent(grandParent) {
			if checkSameDirectionChildWithParent(node) {
				// Zig-Zig
				tree.Rotate(parent)
			} else {
				// Zig-Zag
				tree.Rotate(node)
			}
		}
		// 공통 rotate 작업
		tree.Rotate(node)
	}
}

func (tree *SplayTree) GetRangeSubtreeRootWithGather(start, end int64) *SplayTreeNode {
	tree.gather(start, end)

	subtreeRoot := tree.root.right.left
	subtreeRoot.pushLazyValue()
	return subtreeRoot
}

func (tree *SplayTree) gather(start, end int64) {
	tree.GetKthNodeAndPush(end + 1)
	endNode := tree.root

	tree.GetKthNodeAndPush(start - 1)
	startNode := tree.root

	//tree.PrintDFS()
	tree.splayAndSetChild(startNode, endNode)
}

func (tree *SplayTree) splayAndSetChild(rootNode *SplayTreeNode, child *SplayTreeNode) {
	// TODO : 여기 의심
	if isNil(tree.root) || isNil(child) {
		return
	}

	for child.parent != rootNode && isPresent(child.parent) {
		parent := child.parent
		grandParent := parent.parent

		if grandParent == rootNode {
			tree.Rotate(child)
			break
		}

		//if checkSameDirectionChildWithParent(child) {
		//	tree.Rotate(parent)
		//} else {
		//	tree.Rotate(child)
		//}
		//// 공통 rotate 작업
		//tree.Rotate(child)

		if checkSameDirectionChildWithParent(child) {
			// Zig-Zig
			tree.Rotate(parent)
			tree.Rotate(child)
		} else {
			// Zig-Zag
			tree.Rotate(child)
			tree.Rotate(child)
		}
	}

	tree.root.updateTreeSizeAndSum()
	tree.root.pushLazyValue()
	rootNode.updateTreeSizeAndSum()
	rootNode.pushLazyValue()
	child.updateTreeSizeAndSum()
	child.pushLazyValue()
	if isNil(rootNode) {
		tree.root = child
	}
}

func (tree *SplayTree) Find(key int64) *SplayTreeNode {
	//fmt.Println("Find key: ", key)
	if isNil(tree.root) {
		return nil
	}

	node, parent := tree.findNodeAndParent(key)
	if isPresent(node) {
		tree.Splay(node)
		return node
	} else {
		tree.Splay(parent)
		return nil
	}
}

func (tree *SplayTree) findNodeAndParent(key int64) (node, parent *SplayTreeNode) {
	node = tree.root
	for isPresent(node) && key != node.key {
		parent = node
		if key < node.key {
			node = node.left
		} else {
			node = node.right
		}
	}

	return node, parent
}

func (tree *SplayTree) Insert(key int64) {
	//fmt.Println("Insert key: ", key)
	if isNil(tree.root) {

		tree.root = &SplayTreeNode{key: key, size: 1, sum: key}
		return
	}

	_, parent := tree.findNodeAndParent(key)
	//if isPresent(node) {
	//	return
	//}

	newNode := &SplayTreeNode{key: key, parent: parent, size: 1, sum: key}
	if key < parent.key {
		parent.left = newNode
	} else {
		parent.right = newNode
	}
	tree.splayAndSetChild(nil, newNode)
}

func (tree *SplayTree) Delete(key int64) {
	//fmt.Println("Delete key: ", key)
	if isNil(tree.Find(key)) {
		return
	}

	switch root := tree.root; true {
	case isPresent(root.left) && isPresent(root.right):
		tree.root = root.left
		tree.root.parent = nil

		node := tree.root
		for isPresent(node.right) {
			node = node.right
		}
		// 생각해보면, 왼쪽 서브 트리의 가장 오른쪽 노드는, 오른쪽 서브 트리의 루트 보다 작을 수 밖에 없다
		node.right = root.right
		root.right.parent = node

	case isPresent(root.left):
		tree.root = root.left
		tree.root.parent = nil

	case isPresent(root.right):
		tree.root = root.right
		tree.root.parent = nil

	default:
		tree.root = nil
	}
}

func (tree *SplayTree) SumRange(start, end, value int64) {
	subtreeRoot := tree.GetRangeSubtreeRootWithGather(start, end)
	tree.root.updateTreeSizeAndSum()
	if isNil(subtreeRoot) {
		return
	}
	subtreeRoot.sum += subtreeRoot.size * value
	subtreeRoot.lazy += value
}

func (tree *SplayTree) GetKthNode(k int64) int64 {
	//fmt.Printf("Get %d th Node", k)
	k -= 1
	node := tree.root
	for isPresent(node) {
		for isPresent(node.left) && node.left.size > k {
			node = node.left
		}

		if isPresent(node.left) {
			k -= node.left.size
		}

		if k == 0 {
			break
		}

		k--
		node = node.right
	}

	tree.Splay(node)
	//fmt.Printf(" -> %d\n", node.key)
	return node.key
}

func (tree *SplayTree) GetKthNodeAndPush(k int64) {
	//fmt.Printf("Get %d th Node And Push\n", k)
	//k -= 1	// 더미 떄문에 삭제
	node := tree.root
	node.pushLazyValue()

	for isPresent(node) {
		for isPresent(node.left) && node.left.size > k {
			node = node.left
			node.pushLazyValue()
		}

		if isPresent(node.left) {
			k -= node.left.size
		}

		if k == 0 {
			break
		}

		k--
		node = node.right
		if isPresent(node) {
			node.pushLazyValue()
		}
	}

	tree.splayAndSetChild(nil, node)
	//fmt.Printf(" -> %d\n", node.key)
	return
}

func (node *SplayTreeNode) pushLazyValue() {
	lazyValue := node.lazy
	node.lazy = 0

	//if node.key != math.MinInt64 && node.key != math.MaxInt64 {
	node.key += lazyValue
	//}

	if left := node.left; isPresent(left) {
		left.lazy += lazyValue
		left.sum += left.size * lazyValue
	}

	if right := node.right; isPresent(right) {
		right.lazy += lazyValue
		right.sum += right.size * lazyValue
	}
}

func (node *SplayTreeNode) setGrandParentToParent() {
	parent := node.parent
	grandParent := parent.parent

	// change Parent
	node.parent = grandParent
	parent.parent = node

	if isNil(node.parent) {
		return
	}

	// TODO : grandParent로 바꾸기
	if parent == node.parent.left {
		node.parent.left = node
	} else {
		node.parent.right = node
	}

	// TODO : 중복 제거
	node.updateTreeSizeAndSum()
	parent.updateTreeSizeAndSum()
}

func (node *SplayTreeNode) setParentToChild() {
	parent := node.parent
	var newChild *SplayTreeNode

	if node == parent.left {
		newChild = node.right
		parent.left = newChild
		node.right = parent
	} else {
		newChild = node.left
		parent.right = newChild
		node.left = parent
	}

	if newChild != nil {
		newChild.parent = parent
	}
}

func checkSameDirectionChildWithParent(node *SplayTreeNode) bool {
	parent := node.parent
	grandParent := parent.parent

	isNodeLeft := node == parent.left
	isParentLeft := parent == grandParent.left
	return isNodeLeft == isParentLeft
}

// print 함수
func (tree *SplayTree) PrintDFS() {
	printDFS(tree.root, 0, "root")
	fmt.Println() // 보기 편하려고 넣음
}

func printDFS(node *SplayTreeNode, level int, direction string) {
	if isNil(node) {
		return
	}

	fmt.Printf("%s[%s] : Node %d\n", getIndent(level), direction, node.key)

	printDFS(node.left, level+1, "left")
	printDFS(node.right, level+1, "right")
}

// 출력 시 들여쓰기를 위한 함수
func getIndent(level int) string {
	indent := ""
	for i := 0; i < level; i++ {
		indent += "  "
	}
	return indent
}

func mainOrigin2() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var N, M, K int
	//fmt.Fscanf(reader, "%d %d %d", &N, &M, &K)
	fmt.Fscanln(reader, &N, &M, &K)

	tree := &SplayTree{root: nil}
	tree.Insert(math.MinInt32)
	tree.root.updateTreeSizeAndSum()
	for i := 0; i < N; i++ {
		var value int64
		//fmt.Fscanf(reader, "%d", &value)
		fmt.Fscanln(reader, &value)
		//fmt.Println(value)
		tree.Insert(value)
		tree.root.updateTreeSizeAndSum()
	}
	tree.Insert(math.MaxInt32)
	tree.root.updateTreeSizeAndSum()

	const (
		SUM     = 1
		GET_SUM = 2
	)

	for i := 0; i < M+K; i++ {
		var command, start, end, value int64
		//fmt.Fscanf(reader, "%d %d %d", &command, &start, &end)
		fmt.Fscanln(reader, &command, &start, &end, &value)
		//fmt.Printf("%d %d %d %d\n", command, start, end, value)

		switch command {
		case SUM:
			//fmt.Fscanf(reader, "%d", &value)
			tree.SumRange(start, end, value)
			//fmt.Fprintf(writer, "%d %d %d %d\n", command, start, end, value)

		case GET_SUM:
			subtreeRoot := tree.GetRangeSubtreeRootWithGather(start, end)
			fmt.Fprintf(writer, "%d\n", subtreeRoot.sum)
			//fmt.Println(subtreeRoot.sum)
			//fmt.Fprintf(writer, "%d %d %d %d\n", command, start, end, value)
		}
	}

	writer.Flush()
}
