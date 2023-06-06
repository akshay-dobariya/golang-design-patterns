package intermediate

import (
	"fmt"

	"github.com/akshay-dobariya/golang-design-patterns/creational/prototype/treenode"
)

type IntermediateNode struct {
	value    string
	children *[]treenode.TreeNode
}

func (in *IntermediateNode) SetValue(v string) {
	in.value = v
}

func (in *IntermediateNode) SetChildren(children *[]treenode.TreeNode) {
	in.children = children
}

func (in *IntermediateNode) Print(indent string) {
	fmt.Println(indent, in.value)
	for _, child := range *in.children {
		child.Print(indent + indent)
	}
}

func (in *IntermediateNode) Clone() treenode.TreeNode {
	intermediateNode := IntermediateNode{}
	intermediateNode.SetValue(in.value)
	clonedChildren := []treenode.TreeNode{}
	for _, child := range *in.children {
		clonedChildren = append(clonedChildren, child.Clone())
	}
	intermediateNode.children = &clonedChildren
	return &intermediateNode
}
