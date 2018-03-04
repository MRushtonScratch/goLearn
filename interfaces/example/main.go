package main

import (
	"bytes"
	"fmt"
	"github.com/MRushtonScratch/goLearn/interfaces"
)

func main () {
	in := bytes.NewReader([]byte("Example"))
	out := &bytes.Buffer{}

	fmt.Print("stdout on Copy = ")
	if err := interfaces.Copy(in, out); err != nil {
		panic(err)
	}

	fmt.Println("out bytes buffer =", out.String())

	fmt.Print("stdout on PipeExample = ")

	if err := interfaces.PipeExample(); err != nil {
		panic(err)
	}
}

