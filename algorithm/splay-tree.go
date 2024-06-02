package main

import (
	"fmt"
)

type SplayTree struct {
	root *SplayTreeNode
}

type SplayTreeNode struct {
	key                 int
	left, right, parent *SplayTreeNode
	size                int
}

func (node *SplayTreeNode) isNil() bool {
	return node == nil
}

func (node *SplayTreeNode) isPresent() bool {
	return node != nil
}

func (node *SplayTreeNode) updateSize() {
	node.size = 1
	if node.left.isPresent() {
		node.size += node.left.size
	}

	if node.right.isPresent() {
		node.size += node.right.size
	}
}

func (tree *SplayTree) Rotate(node *SplayTreeNode) {
	parent := node.parent
	if parent.isNil() {
		return
	}

	node.setParentToChild()
	node.setGrandParentToParent()

	if node.parent.isNil() {
		tree.root = node
	}

	// update Child Count
	node.updateSize()
	parent.updateSize()
}

func (tree *SplayTree) Splay(node *SplayTreeNode) {
	if tree.root.isNil() {
		return
	}

	if node.parent == tree.root {
		tree.Rotate(node)
		return
	}

	for node.parent.isPresent() {
		parent := node.parent
		grandParent := parent.parent

		if grandParent.isPresent() {
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

func (tree *SplayTree) gather(start, end int) {
	tree.GetKthNode(end + 1)
	endNode := tree.root

	tree.GetKthNode(start - 1)
	startNode := tree.root

	tree.splayAndSetChild(startNode, endNode)
}

func (tree *SplayTree) splayAndSetChild(rootNode *SplayTreeNode, child *SplayTreeNode) {
	for child.parent != rootNode {
		parent := child.parent

		if parent.parent == rootNode {
			tree.Rotate(child)
			break
		}

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

	if rootNode.isNil() {
		tree.root = child
	}
}

func (tree *SplayTree) Find(key int) *SplayTreeNode {
	fmt.Println("Find key: ", key)
	if tree.root.isNil() {
		return nil
	}

	node, parent := tree.findNodeAndParent(key)
	if node.isPresent() {
		tree.Splay(node)
		return node
	} else {
		tree.Splay(parent)
		return nil
	}
}

func (tree *SplayTree) findNodeAndParent(key int) (node, parent *SplayTreeNode) {
	node = tree.root
	for node.isPresent() && key != node.key {
		parent = node
		if key < node.key {
			node = node.left
		} else {
			node = node.right
		}
	}

	return node, parent
}

func (tree *SplayTree) Insert(key int) {
	fmt.Println("Insert key: ", key)
	if tree.root.isNil() {
		tree.root = &SplayTreeNode{key: key}
		return
	}

	node, parent := tree.findNodeAndParent(key)
	if node.isPresent() {
		return
	}

	newNode := &SplayTreeNode{key: key, parent: parent}
	if key < parent.key {
		parent.left = newNode
	} else {
		parent.right = newNode
	}
	tree.Splay(newNode)
}

func (tree *SplayTree) Delete(key int) {
	fmt.Println("Delete key: ", key)
	if tree.Find(key).isNil() {
		return
	}

	switch root := tree.root; true {
	case root.left.isPresent() && root.right.isPresent():
		tree.root = root.left
		tree.root.parent = nil

		node := tree.root
		for node.right.isPresent() {
			node = node.right
		}
		// 생각해보면, 왼쪽 서브 트리의 가장 오른쪽 노드는, 오른쪽 서브 트리의 루트 보다 작을 수 밖에 없다
		node.right = root.right
		root.right.parent = node

	case root.left.isPresent():
		tree.root = root.left
		tree.root.parent = nil

	case root.right.isPresent():
		tree.root = root.right
		tree.root.parent = nil

	default:
		tree.root = nil
	}
}

func (tree *SplayTree) GetKthNode(k int) int {
	fmt.Printf("Get %d th Node", k)
	k -= 1
	node := tree.root
	for node.isPresent() {
		for node.left.isPresent() && node.left.size > k {
			node = node.left
		}

		if node.left.isPresent() {
			k -= node.left.size
		}

		if k == 0 {
			break
		}

		k--
		node = node.right
	}

	tree.Splay(node)
	fmt.Printf(" -> %d\n", node.key)
	return node.key
}

func (node *SplayTreeNode) setGrandParentToParent() {
	parent := node.parent
	grandParent := parent.parent

	// change Parent
	node.parent = grandParent
	parent.parent = node

	if grandParent.isNil() {
		return
	}

	if parent == grandParent.left {
		grandParent.left = node
	} else {
		grandParent.right = node
	}
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
	if node.isNil() {
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
