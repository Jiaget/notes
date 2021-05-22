package algorithm

// https://studygolang.com/articles/7014
const (
	RED   bool = true
	BLACK bool = false
)

type RBNode struct {
	value               int
	color               bool
	left, right, parent *RBNode
}

func (rbnode *RBNode) getGrandParent() *RBNode {
	if rbnode.parent == nil {
		return nil
	}
	return rbnode.parent.parent
}

func (rbnode *RBNode) getBrother() *RBNode {
	if rbnode.parent == nil {
		return nil
	}
	if rbnode == rbnode.parent.left {
		return rbnode.parent.right
	}
	return rbnode.parent.left
}

func (rbnode *RBNode) getUncle() *RBNode {
	if rbnode.parent == nil {
		return nil
	}
	if rbnode.parent == rbnode.getGrandParent().left {
		return rbnode.getGrandParent().right
	}
	return rbnode.getGrandParent().left
}
