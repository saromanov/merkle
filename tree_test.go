package merkle

import (
	"crypto/md5"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func splitData(data []byte, size int) [][]byte {
	count := len(data) / size
	blocks := make([][]byte, 0, count)
	for i := 0; i < count; i++ {
		block := data[i*size : (i+1)*size]
		blocks = append(blocks, block)
	}
	if len(data)%size != 0 {
		blocks = append(blocks, data[len(blocks)*size:])
	}
	return blocks
}

func TestTree(t *testing.T) {
	data, err := ioutil.ReadFile("tree.go")
	if err != nil {
		t.Errorf("unable to open file: %v", err)
	}
	assert.NoError(t, err, "shouldn't contain error")
	blocks := splitData(data, 32)

	tree := NewTree()
	err = tree.Generate(blocks, md5.New())
	assert.NoError(t, err, "shouldn't contain error")
	assert.Equal(t, tree.Height(), uint64(8), "error")
	assert.Equal(t, tree.Root().Hash, []byte{0xae, 0x7, 0x83, 0x58, 0x5d, 0x27, 0x37, 0x53, 0x98, 0xe7, 0x67, 0xf7, 0x9c, 0x79, 0x5f, 0x4b}, "error")
}
