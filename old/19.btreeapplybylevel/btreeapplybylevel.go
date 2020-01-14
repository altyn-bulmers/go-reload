package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	levels := BTreeLevelCount(root)
	for i := 1; i <= levels; i++ {
		Levels(root, i, f)

	}
}

func Levels(root *TreeNode, level int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if level == 1 {
		f(root.Data)
	} else {
		Levels(root.Left, level-1, f)
		Levels(root.Right, level-1, f)
	}
}
