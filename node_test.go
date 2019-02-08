package merkle

import (
	"crypto/md5"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNode(t *testing.T) {
	_, err := NewNode(nil, []byte("test"))
	assert.Error(t, err, "should return error")

	_, err = NewNode(md5.New(), nil)
	assert.Error(t, err, "should return error")

	_, err = NewNode(md5.New(), []byte("test"))
	assert.NoError(t, err, "should't return error")
}
