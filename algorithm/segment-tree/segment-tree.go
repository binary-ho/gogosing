package main

//
//import (
//	"bufio"
//	"fmt"
//	"os"
//)
//
//type SegmentTree struct {
//	values [][]int64
//}
//
//func InitSegmentTree(N int, values [][]int64) (tree *SegmentTree) {
//	size := 4 * (N * N)
//	tree = &SegmentTree{values: make([][]int64, size+1)}
//	for i := 0; i <= size; i++ {
//		tree.values[i] = make([]int64, size+1)
//	}
//	tree.initSegmentTree(1, 1, N, 1, N, values)
//	return tree
//}
//
//func (tree *SegmentTree) initSegmentTree(nodeNumber, rowStart, rowEnd, colStart, colEnd int, values [][]int64) int64 {
//	if rowStart == rowEnd && colStart == colEnd {
//		tree.values[nodeNumber] = values[rowStart][colStart]
//		return tree.values[nodeNumber]
//	}
//
//	rowMid := (rowStart + rowEnd) / 2
//	colMid := (colStart + colEnd) / 2
//	leftUpSide := tree.initSegmentTree(nodeNumber*4-2, rowStart, rowMid, colStart, colMid, values)    // left child
//	rightUpSide := tree.initSegmentTree(nodeNumber*4-1, rowStart, rowMid, colMid+1, colEnd, values)   // left child
//	leftDownSide := tree.initSegmentTree(nodeNumber*4, rowMid+1, rowEnd, colStart, colMid, values)    // right child
//	rightDownSide := tree.initSegmentTree(nodeNumber*4+1, rowMid+1, rowEnd, colMid+1, colEnd, values) // right child
//
//	node := tree.children[nodeNumber]
//	node.setValues(leftUpSide+rightUpSide+leftDownSide+rightDownSide, rowStart, rowEnd, colStart, colEnd)
//	return node.sum
//}
//
//func (tree *SegmentTree) Query(rowStart, rowEnd, colStart, colEnd int) int64 {
//	return tree.query(1, rowStart, rowEnd, colStart, colEnd)
//}
//
//func (tree *SegmentTree) query(nodeNumber, rowStart, rowEnd, colStart, colEnd int) int64 {
//	node := tree.children[nodeNumber]
//	//fmt.Println("[Query] : ", nodeNumber, " // ", node.rowStart, " ~ ", node.rowEnd, ", ", node.colStart, " ~ ", node.colEnd)
//	if (rowStart > node.rowEnd || node.rowStart > rowEnd) ||
//		(colStart > node.colEnd || node.colStart > colEnd) {
//		//fmt.Println("[END] ", nodeNumber, "- 범위 완전 아웃 0")
//		return 0
//	}
//
//	if (rowStart <= node.rowStart && node.rowEnd <= rowEnd) &&
//		(colStart <= node.colStart && node.colEnd <= colEnd) {
//		//fmt.Println("[END] ", nodeNumber, "- 범위 완전 포함 ", node.sum)
//		return node.sum
//	}
//
//	leftUpSide := tree.query(nodeNumber*4-2, rowStart, rowEnd, colStart, colEnd)
//	rightUpSide := tree.query(nodeNumber*4-1, rowStart, rowEnd, colStart, colEnd)
//	leftDownSide := tree.query(nodeNumber*4, rowStart, rowEnd, colStart, colEnd)
//	rightDownSide := tree.query(nodeNumber*4+1, rowStart, rowEnd, colStart, colEnd)
//	//fmt.Println("[END] ", nodeNumber, "- 범위 걸침 ", leftUpSide+rightUpSide+leftDownSide+rightDownSide)
//	return leftUpSide + rightUpSide + leftDownSide + rightDownSide
//}
//
//func (tree *SegmentTree) Update(targetY, targetX int, newValue int64) {
//	tree.update(1, targetY, targetX, newValue)
//}
//
//func (tree *SegmentTree) update(nodeNumber, targetY, targetX int, newValue int64) int64 {
//	node := tree.children[nodeNumber]
//
//	if (targetY > node.rowEnd || node.rowStart > targetY) ||
//		(targetX > node.colEnd || node.colStart > targetX) {
//		return node.sum
//	}
//
//	if (targetY == node.rowStart && node.rowEnd == targetY) &&
//		(targetX == node.colStart && node.colEnd == targetX) {
//		node.sum = newValue
//		return node.sum
//	}
//
//	leftUpSide := tree.update(nodeNumber*4-2, targetY, targetX, newValue)
//	rightUpSide := tree.update(nodeNumber*4-1, targetY, targetX, newValue)
//	leftDownSide := tree.update(nodeNumber*4, targetY, targetX, newValue)
//	rightDownSide := tree.update(nodeNumber*4+1, targetY, targetX, newValue)
//	node.sum = leftUpSide + rightUpSide + leftDownSide + rightDownSide
//	return node.sum
//}
//
//func (node *SegmentTreeNode) setValues(value int64, rowStart, rowEnd, colStart, colEnd int) {
//	node.sum = value
//	node.rowStart = rowStart
//	node.rowEnd = rowEnd
//	node.colStart = colStart
//	node.colEnd = colEnd
//}
//
//func main() {
//	reader := bufio.NewReader(os.Stdin)
//	writer := bufio.NewWriter(os.Stdout)
//
//	var N, M int
//	fmt.Fscanln(reader, &N, &M)
//
//	var values = make([][]int64, N+1)
//	for i := 1; i <= N; i++ {
//		values[i] = make([]int64, N+1)
//		for j := 1; j <= N; j++ {
//			fmt.Fscanf(reader, "%d ", &values[i][j])
//		}
//	}
//
//	//for i := 1; i <= N; i++ {
//	//	fmt.Fprintln(writer, &values[i])
//	//	// 이렇게 하는거 맞는지 확인
//	//	//fmt.Fscanln(reader)
//	//}
//	//
//	//writer.Flush()
//	//return
//	tree := InitSegmentTree(N, values)
//	//for i := 0; i < len(tree.children); i++ {
//	//	node := tree.children[i]
//	//	fmt.Fprintf(writer, "id: %d, sum: %d, rowStart: %d, rowEnd: %d, colStart: %d, colEnd: %d\n", i, node.sum, node.rowStart, node.rowEnd, node.colStart, node.colEnd)
//	//}
//
//	const (
//		UPDATE = 0
//		QUERY  = 1
//	)
//
//	for i := 0; i < M; i++ {
//		var command, rowStart, rowEnd, colStart, colEnd int
//		fmt.Fscanln(reader, &command, &rowStart, &colStart, &rowEnd, &colEnd)
//
//		switch command {
//		case QUERY:
//			result := tree.Query(rowStart, rowEnd, colStart, colEnd)
//			fmt.Fprintf(writer, "%d\n", result)
//
//		case UPDATE:
//			tree.Update(rowStart, colStart, int64(rowEnd))
//		}
//	}
//
//	//for i := 0; i < len(tree.children); i++ {
//	//	node := tree.children[i]
//	//	fmt.Fprintf(writer, "id: %d, sum: %d, lazy: %d, rowStart: %d, rowEnd: %d, colStart: %d, colEnd: %d\n", i, node.sum, node.lazy, node.rowStart, node.rowEnd, node.colStart, node.colEnd)
//	//}
//	writer.Flush()
//}
