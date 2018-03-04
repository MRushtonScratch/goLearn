package main

import (
	"github.com/MRushtonScratch/goLearn/bytestrings"
)

func main() {
	err := bytestrings.WorkWithBuffer()

	if(err != nil) {
		panic(err)
	}

	bytestrings.SearchString()
	bytestrings.ModifyString()
	bytestrings.StringReader()
}
