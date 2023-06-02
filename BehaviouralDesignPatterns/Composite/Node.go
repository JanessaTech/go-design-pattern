package composite

import "fmt"

type iNode interface {
	getValue() int
}

type NonLeafNode struct {
	nodes []iNode
}

func (NonLeafNode *NonLeafNode) getValue() int {
	if NonLeafNode.nodes != nil {
		sum := 0
		for _, n := range NonLeafNode.nodes {
			sum += n.getValue()
		}
		return sum
	}
	return 0
}
func (NonLeafNode *NonLeafNode) addNode(node iNode) {
	NonLeafNode.nodes = append(NonLeafNode.nodes, node)
}

type LeafNode struct {
	value int
}

func (LeafNode *LeafNode) getValue() int {
	return LeafNode.value
}

func Main() {
	leafNode1 := LeafNode{value: 1}
	leafNode2 := LeafNode{value: 2}

	root := NonLeafNode{nodes: []iNode{}}
	root.addNode(&leafNode1)
	root.addNode(&leafNode2)

	value := root.getValue()
	fmt.Println("value =", value)

}
