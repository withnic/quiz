package main

import (
"crypto/sha1"
"fmt"
)

// https://twitter.com/davecheney/status/1041526653141147654
// #golang pop quiz: what does this program print?
// 461ec8144f
//[70 30 200 20 79]
// nothing, doesn't compile
func main() {
	input := []byte("Hello, playground")
	hash := sha1.Sum(input)[:5]
	fmt.Println(hash)
}