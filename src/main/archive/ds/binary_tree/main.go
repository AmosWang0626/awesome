package main

import "fmt"

type binaryTree struct {
	value int
	left  *binaryTree
	right *binaryTree
}

// 二叉树 先序、中序、后序遍历
func main() {

	binTree1 := &binaryTree{value: 1}
	binTree2 := &binaryTree{value: 2}
	binTree3 := &binaryTree{value: 3}
	binTree4 := &binaryTree{value: 4}
	binTree5 := &binaryTree{value: 5}
	binTree6 := &binaryTree{value: 6}
	binTree7 := &binaryTree{value: 7}

	binTree1.left = binTree2
	binTree1.right = binTree3
	binTree2.left = binTree4
	binTree2.right = binTree5
	binTree3.left = binTree6
	binTree3.right = binTree7

	binTree1.leftRange()
	fmt.Println()
	binTree1.centerRange()
	fmt.Println()
	binTree1.rightRange()

}

func (current *binaryTree) leftRange() {
	fmt.Print(current.value, "\t")
	if current.left != nil {
		current.left.leftRange()
	}
	if current.right != nil {
		current.right.leftRange()
	}
}

func (current *binaryTree) centerRange() {
	if current.left != nil {
		current.left.centerRange()
	}
	fmt.Print(current.value, "\t")
	if current.right != nil {
		current.right.centerRange()
	}
}

func (current *binaryTree) rightRange() {
	if current.left != nil {
		current.left.rightRange()
	}
	if current.right != nil {
		current.right.rightRange()
	}
	fmt.Print(current.value, "\t")
}
