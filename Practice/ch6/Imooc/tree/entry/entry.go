package main

import "imooc/tree"
import "fmt"

type myTreeNode struct {
	node *tree.Node
}

func (myNode *myTreeNode) postOrder() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	left.postOrder()
	right := myTreeNode{myNode.node.Right}
	right.postOrder()
	myNode.node.Print()
}

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}

	root.Left = &tree.Node{}

	root.Right = &tree.Node{Value: 5}

	root.Right.Left = new(tree.Node)

	root.Left.Right = tree.CreateNode(2)

	root.SetValue(4)

	root.Traverse()

	myNode := myTreeNode{&root}
	fmt.Println()
	myNode.postOrder()

	nodeCount := 0
	root.TraverseFunc(func(n *tree.Node) {
		n.Value += 1
		nodeCount++
	})

	fmt.Println()

	root.Traverse()

	fmt.Println()

	fmt.Println(nodeCount)
}
