package main

func main() {
	// 예제 트리 생성
	tree := &SplayTree{
		root: &SplayTreeNode{
			key: 10,
		},
	}

	// 트리 구조 출력
	tree.PrintDFS()

	tree.Insert(5)
	tree.PrintDFS()

	tree.Insert(3)
	tree.PrintDFS()

	tree.Find(5)
	tree.PrintDFS()

	tree.Insert(2)
	tree.PrintDFS()

	tree.Insert(1)
	tree.PrintDFS()

	tree.Find(10)
	tree.PrintDFS()

	tree.Insert(7)
	tree.PrintDFS()

	tree.Find(2)
	tree.PrintDFS()

	tree.Insert(8)
	tree.PrintDFS()

	tree.Insert(9)
	tree.PrintDFS()

	tree.Find(3)
	tree.PrintDFS()

	tree.Delete(7)
	tree.PrintDFS()

	tree.Delete(3)
	tree.PrintDFS()

	tree.Delete(1)
	tree.PrintDFS()
}
