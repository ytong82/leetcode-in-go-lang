package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func doPathSum(root *TreeNode, sum int, path []int, res *[][]int) {
	if root == nil {
		return
	}

	remain := sum - root.Val
	path = append(path, root.Val)
	if remain == 0 {
		if root.Left == nil && root.Right == nil {
			pth := make([]int, len(path))
			copy(pth, path)
			*res = append(*res, pth)
			return
		}
	}
	doPathSum(root.Left, remain, path, res)
	doPathSum(root.Right, remain, path, res)
}

func pathSum(root *TreeNode, sum int) [][]int {
	var path []int
	var res [][]int
	doPathSum(root, sum, path, &res)
	return res
}

func makeTree() *TreeNode {
	root := TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}

	node := TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}

	node2 := TreeNode{
		Val:   8,
		Left:  nil,
		Right: nil,
	}

	node3 := TreeNode{
		Val:   11,
		Left:  nil,
		Right: nil,
	}

	node4 := TreeNode{
		Val:   13,
		Left:  nil,
		Right: nil,
	}

	node5 := TreeNode{
		Val:   4,
		Left:  nil,
		Right: nil,
	}

	node6 := TreeNode{
		Val:   7,
		Left:  nil,
		Right: nil,
	}

	node7 := TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}

	node8 := TreeNode{
		Val:   5,
		Left:  nil,
		Right: nil,
	}

	node9 := TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}

	root.Left = &node
	root.Right = &node2

	node.Left = &node3

	node2.Left = &node4
	node2.Right = &node5

	node3.Left = &node6
	node3.Right = &node7

	node5.Left = &node8
	node5.Right = &node9

	return &root
}

func makeTree2() *TreeNode {
	root := TreeNode{
		Val:   -2,
		Left:  nil,
		Right: nil,
	}

	node := TreeNode{
		Val:   -3,
		Left:  nil,
		Right: nil,
	}

	root.Right = &node

	return &root
}

func main() {
	root := makeTree()
	res := pathSum(root, 22)
	fmt.Println("final res: ", res)

	root = makeTree2()
	res = pathSum(root, -5)
	fmt.Println("final res: ", res)
}
