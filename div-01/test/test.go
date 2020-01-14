package main

import (
	"fmt"
	student ".."
)

func main() {
	root := &student.TreeNode{Data: "4"}
	student.BTreeInsertData(root, "1")
	student.BTreeInsertData(root, "7")
	student.BTreeInsertData(root, "5")
	node := student.BTreeSearchItem(root, "4")
	fmt.Println("Before delete:")
	student.BTreeApplyInorder(root, fmt.Println)
	root = student.BTreeDeleteNode(root, node)
	fmt.Println("After delete:")
	student.BTreeApplyInorder(root, fmt.Println)
}