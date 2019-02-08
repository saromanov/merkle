package merkle

import (
	"errors"
	"fmt"
	"hash"
)

var errNewNode = errors.New("unable to create a new node")

// Node defines node of the Merkle tree
type Node struct {
	Hash  []byte
	Left  *Node
	Right *Node
}

// NewNode creates a new node at the tree
func NewNode(h hash.Hash, block []byte) (Node, error) {
	if h == nil || block == nil {
		return Node{}, errNewNode
	}
	defer h.Reset()
	_, err := h.Write(block)
	if err != nil {
		return Node{}, fmt.Errorf("unable write to hash: %v", err)
	}
	return Node{Hash: h.Sum(nil)}, nil
}
