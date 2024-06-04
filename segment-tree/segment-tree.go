package main

import (
	"bufio"
	"fmt"
	"os"
)

type SegmentTree struct {
	children []*SegmentTreeNode
}

func (tree *SegmentTree) GetSize() int {
	return len(tree.children) - 1
}

func (tree *SegmentTree) IsNotLeaf(nodeNumber int) bool {
	end := tree.GetSize()
	start := (end + 1) / 2
	return !(start <= nodeNumber && nodeNumber <= end)
}

type SegmentTreeNode struct {
	start, end int
	sum, lazy  int64
}

func InitSegmentTree(N int, values []int64) (tree *SegmentTree) {
	size := 4 * N
	tree = &SegmentTree{children: make([]*SegmentTreeNode, size+1)}
	for i := 0; i <= size; i++ {
		tree.children[i] = &SegmentTreeNode{}
	}
	tree.initSegmentTree(1, 1, N, values)
	return tree
}

func (tree *SegmentTree) initSegmentTree(nodeNumber, start, end int, values []int64) int64 {
	if start == end {
		node := tree.children[nodeNumber]
		node.setValues(values[start], start, end)
		return node.sum
	}

	mid := (start + end) / 2
	leftSum := tree.initSegmentTree(nodeNumber*2, start, mid, values)    // left child
	rightSum := tree.initSegmentTree(nodeNumber*2+1, mid+1, end, values) // right child

	node := tree.children[nodeNumber]
	node.setValues(leftSum+rightSum, start, end)
	return node.sum
}

func (tree *SegmentTree) Query(start, end int) int64 {
	return tree.query(1, start, end)
}

func (tree *SegmentTree) query(nodeNumber, start, end int) int64 {
	node := tree.children[nodeNumber]
	tree.pushLazy(nodeNumber)

	if start > node.end || node.start > end {
		return 0
	}

	if start <= node.start && node.end <= end {
		return node.sum
	}

	leftSum := tree.query(nodeNumber*2, start, end)
	rightSum := tree.query(nodeNumber*2+1, start, end)
	return leftSum + rightSum
}

func (tree *SegmentTree) Update(targetStart, targetEnd int, newValue int64) {
	tree.update(1, targetStart, targetEnd, newValue)
}

func (tree *SegmentTree) update(nodeNumber, targetStart, targetEnd int, newValue int64) int64 {
	node := tree.children[nodeNumber]
	tree.pushLazy(nodeNumber)

	if node.start > targetEnd || targetStart > node.end {
		return node.sum
	}

	if targetStart <= node.start && node.end <= targetEnd {
		node.sum += newValue * int64(node.end-node.start+1)
		if tree.IsNotLeaf(nodeNumber) {
			tree.children[nodeNumber*2].lazy += newValue
			tree.children[nodeNumber*2+1].lazy += newValue
		}
		return node.sum
	}

	leftSum := tree.update(nodeNumber*2, targetStart, targetEnd, newValue)
	rightSum := tree.update(nodeNumber*2+1, targetStart, targetEnd, newValue)
	node.sum = leftSum + rightSum
	return node.sum
}

func (tree *SegmentTree) pushLazy(nodeNumber int) {
	node := tree.children[nodeNumber]
	if node.lazy == 0 {
		return
	}

	node.sum += node.lazy * int64(node.end-node.start+1)
	if tree.IsNotLeaf(nodeNumber) {
		tree.children[nodeNumber*2].lazy += node.lazy
		tree.children[nodeNumber*2+1].lazy += node.lazy
	}
	node.lazy = 0
}

func (node *SegmentTreeNode) setValues(value int64, start, end int) {
	node.sum = value
	node.start = start
	node.end = end
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	var N, M, K int
	fmt.Fscanln(reader, &N, &M, &K)

	var values = make([]int64, N+1)
	for i := 1; i <= N; i++ {
		fmt.Fscanln(reader, &values[i])
	}
	tree := InitSegmentTree(N, values)

	const (
		SUM     = 1
		GET_SUM = 2
	)

	for i := 0; i < M+K; i++ {
		var start, end int
		var command, value int64
		fmt.Fscanln(reader, &command, &start, &end, &value)

		switch command {
		case SUM:
			tree.Update(start, end, value)

		case GET_SUM:
			result := tree.Query(start, end)
			fmt.Fprintf(writer, "%d\n", result)
		}
	}

	//for i := 0; i < len(tree.children); i++ {
	//	node := tree.children[i]
	//	fmt.Fprintf(writer, "id: %d, sum: %d, lazy: %d, start: %d, end: %d\n", i, node.sum, node.lazy, node.start, node.end)
	//}
	writer.Flush()
}
