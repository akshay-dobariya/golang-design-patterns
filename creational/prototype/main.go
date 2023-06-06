package main

import (
	"fmt"

	"github.com/akshay-dobariya/golang-design-patterns/creational/prototype/treenode"
	"github.com/akshay-dobariya/golang-design-patterns/creational/prototype/treenode/intermediate"
	"github.com/akshay-dobariya/golang-design-patterns/creational/prototype/treenode/leaf"
)

func createLeafNode(value string) *leaf.LeafNode {
	node := leaf.LeafNode{}
	node.SetValue(value)
	return &node
}

func createIntermediateNode(value string) *intermediate.IntermediateNode {
	node := intermediate.IntermediateNode{}
	node.SetValue(value)
	return &node
}

/*
    root -- node-1 -- node-1.1
				   -- node-1.2
		 -- node-2 -- node-2.1
				   -- node-2.2
				   -- node-2.3
		 -- node-3 -- node-3.1
				   -- node-3.2
				   -- node-3.3
				   -- node-3.4
*/
func getNewTree() treenode.TreeNode {
	// create leaf nodes
	intermediateNodes := []treenode.TreeNode{}
	for i := 1; i <= 3; i++ {
		leafNodes := []treenode.TreeNode{}
		for j := 1; j <= i; j++ {
			node := createLeafNode(fmt.Sprintf("node-%d.%d", i, j))
			leafNodes = append(leafNodes, node)
		}
		intermediateNode := createIntermediateNode(fmt.Sprintf("node-%d", i))
		intermediateNode.SetChildren(&leafNodes)
		intermediateNodes = append(intermediateNodes, intermediateNode)
	}
	root := createIntermediateNode("root")
	root.SetChildren(&intermediateNodes)
	return root
}
func main() {
	tree := getNewTree()
	tree.Print("  ")

	clonedTree := tree.Clone()
	clonedTree.Print("  ")
}
