package piscine

type TreeNode struct {
	Left, Right, Parent *TreeNode
	Data                string
}

func BTreeTransplant(root, node, rplc *TreeNode) *TreeNode {

	if node.Parent == nil {
		root = rplc
	} else if node == node.Parent.Left {
		node.Parent.Left = rplc
	} else if node == node.Parent.Right {
		node.Parent.Right = rplc
	}

	// if rplc has replaced node (is not nil) => rplc has the same parent as node
	if node != nil {
		rplc.Parent = node.Parent
	}

	return root
}
