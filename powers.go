package merkle

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
