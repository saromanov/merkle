package merkle

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPowersOfTwo(t *testing.T) {
	r := isPowerOfTwo(5)
	assert.Equal(t, r, false, "not equal")

	r = isPowerOfTwo(16)
	assert.Equal(t, r, true, "not equal")
}

func TestNextPowersOfTwo(t *testing.T) {
	r := nextPowerOfTwo(5)
	assert.Equal(t, r, uint64(8), "not equal")

	r = nextPowerOfTwo(10)
	assert.Equal(t, r, uint64(16), "not equal")
}
