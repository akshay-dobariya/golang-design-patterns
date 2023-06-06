package leaf

import (
	"fmt"

	"github.com/akshay-dobariya/golang-design-patterns/creational/prototype/treenode"
)

type LeafNode struct {
	value string
}

func (ln *LeafNode) SetValue(v string) {
	ln.value = v
}

func (ln *LeafNode) Print(indent string) {
	fmt.Println(indent, ln.value)
}

func (ln *LeafNode) Clone() treenode.TreeNode {
	leafNode := LeafNode{}
	leafNode.SetValue(ln.value)
	return &leafNode
}
