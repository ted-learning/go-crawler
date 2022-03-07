package main

import (
	"bufio"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"io"
)

func determineEncoding(reader io.Reader) encoding.Encoding {
	peek, err := bufio.NewReader(reader).Peek(1024)
	panicErr(err)

	encode, _, _ := charset.DetermineEncoding(peek, "")
	return encode
}

func panicErr(err error) {
	if err != nil {
		panic(err)
	}
}
