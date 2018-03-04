package bytestrings

import (
	"fmt"
	"bytes"
	"bufio"
)

func WorkWithBuffer() error {
	rawString := "it's easy to encode unicode into a byte array ྿"

	b := Buffer(rawString)
	fmt.Println(b.String())

	s, err := toString(b)
	if err != nil {
		return err
	}

	fmt.Println(s)

	reader := bytes.NewReader([]byte(rawString))
	scanner := bufio.NewScanner(reader)

	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}

	return nil
}
