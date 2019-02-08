package merkle

// Tree defines merkle tree
type Tree struct {
	Nodes  []Node
	Levels [][]Node
}

// NewTree create a tree
func NewTree() Tree {
	return Tree{
		Nodes:  nil,
		Levels: nil,
	}
}
