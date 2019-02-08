# merkle
[![GoDoc](https://godoc.org/github.com/saromanov/merkle?status.png)](https://godoc.org/github.com/saromanov/merkle)
[![Go Report Card](https://goreportcard.com/badge/github.com/saromanov/merkle)](https://goreportcard.com/report/github.com/saromanov/merkle)
[![Build Status](https://travis-ci.org/saromanov/merkle.svg?branch=master)](https://travis-ci.org/saromanov/merkle)
[![Coverage Status](https://coveralls.io/repos/github/saromanov/merkle/badge.svg?branch=master)](https://coveralls.io/github/saromanov/merkle?branch=master)

Implementation of Merkle tree

### Example
```go
package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"

	"github.com/saromanov/merkle"
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

func main() {
	data, err := ioutil.ReadFile("file")
	if err != nil {
		fmt.Println(err)
		return
	}
	blocks := splitData(data, 32)

	tree := merkle.NewTree()
	err = tree.Generate(blocks, md5.New())
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Height: %d\n", tree.Height())
	fmt.Printf("Root: %v\n", tree.Root())
}

```
