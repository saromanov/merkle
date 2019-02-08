package merkle

import (
	"errors"
	"hash"
	"math"
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

	t.Levels = levels
	t.Nodes = nodes
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

func createNodeLevel(below []Node, currentNode []Node, h hash.Hash) (uint64, error) {

	h.Reset()
	size := h.Size()
	data := make([]byte, size*2)
	end := getEndTree(below)
	for i := 0; i < end; i++ {
		node := Node{}
		ileft := 2 * i
		iright := 2*i + 1
		left := &below[ileft]
		var right *Node
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
		node.Left = left
		node.Right = right
		currentNode[i] = node
		data = data[:]
	}
	return uint64(end), nil
}

func getEndTree(t []Node) int {
	return (len(t) + (len(t) % 2)) / 2
}

// CalculateHeightAndNodeCount returns height and number of nodes
func CalculateHeightAndNodeCount(leaves uint64) (height, nodeCount uint64) {
	height = calculateTreeHeight(leaves)
	return height, calculateNodeCount(height, leaves)
}

func calculateNodeCount(height, size uint64) uint64 {
	if isPowerOfTwo(size) {
		return 2*size - 1
	}
	count := size
	prev := size
	i := uint64(1)
	for ; i < height; i++ {
		next := (prev + (prev % 2)) / 2
		count += next
		prev = next
	}
	return count
}

// calculateTreeHeight provides getting of height of the binary tree
func calculateTreeHeight(nodeCount uint64) uint64 {
	if nodeCount == 1 {
		return 2
	}
	if nodeCount == 0 {
		return nodeCount
	}
	return uint64(math.Log2(float64(nextPowerOfTwo(nodeCount)))) + 1
}

// isPowerOfTwo returns true if n is power of 2
func isPowerOfTwo(n uint64) bool {
	return n != 0 && (n&(n-1)) == 0
}

func nextPowerOfTwo(n uint64) uint64 {
	if n == 0 {
		return 1
	}
	// http://graphics.stanford.edu/~seander/bithacks.html#RoundUpPowerOf2
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n |= n >> 32
	n++
	return n
}
