package merkle

import (
	"errors"
	"hash"
)

var errEmptyTree = errors.New("tree is empty")

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

// Root retruns root of the tree
func (t *Tree) Root() *Node {
	if t.Nodes == nil {
		return nil
	}

	return &t.Levels[0][0]
}

// Height returns tree's height
func (t *Tree) Height() uint64 {
	return uint64(len(t.Levels))
}

// Generate provides generation of the node
func (t *Tree) Generate(blocks [][]byte, hashf hash.Hash) error {
	blockCount := uint64(len(blocks))
	if blockCount == 0 {
		return errEmptyTree
	}
	height, nodeCount := CalculateHeightAndNodeCount(blockCount)
	levels := make([][]Node, height)
	nodes, err := makeNodes(nodeCount, hashf, blocks)
	if err != nil {
		return err
	}
	levels[height-1] = nodes[:len(blocks)]

	currentNode := nodes[len(blocks):]
	h := height - 1
	for ; h > 0; h-- {
		below := levels[h]
		wrote, err := createNodeLevel(below, currentNode, hashf)
		if err != nil {
			return err
		}
		levels[h-1] = currentNode[:wrote]
		currentNode = currentNode[wrote:]
	}

	t.Nodes = nodes
	t.Levels = levels
	return nil
}

// makeNodes create nodes for the tree
func makeNodes(nodeCount uint64, hashf hash.Hash, blocks [][]byte) ([]Node, error) {
	nodes := make([]Node, nodeCount)
	for i, block := range blocks {
		node, err := NewNode(hashf, block)
		if err != nil {
			return nil, err
		}
		nodes[i] = node
	}
	return nodes, nil
}

func createNodeLevel(below []Node, current []Node,
	h hash.Hash) (uint64, error) {
	h.Reset()
	size := h.Size()
	data := make([]byte, size*2)
	end := (len(below) + (len(below) % 2)) / 2
	for i := 0; i < end; i++ {
		// Concatenate the two children hashes and hash them, if both are
		// available, otherwise reuse the hash from the lone left node
		node := Node{}
		ileft := 2 * i
		iright := 2*i + 1
		left := &below[ileft]
		var right *Node = nil
		if len(below) > iright {
			right = &below[iright]
		}
		if right == nil {
			b := data[:size]
			copy(b, left.Hash)
			node = Node{Hash: b}
		} else {
			copy(data[:size], below[ileft].Hash)
			copy(data[size:], below[iright].Hash)
			var err error
			node, err = NewNode(h, data)
			if err != nil {
				return 0, err
			}
		}
		// Point the new node to its children and save
		node.Left = left
		node.Right = right
		current[i] = node

		// Reset the data slice
		data = data[:]
	}
	return uint64(end), nil
}
