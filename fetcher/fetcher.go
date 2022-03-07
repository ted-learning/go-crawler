package fetcher

import (
	"bufio"
	"fmt"
	"go-crawler/common"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetcher(url string) ([]byte, error) {
	log.Printf("Fetching URL %s\n", url)
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			common.PanicErr(err)
		}
	}(response.Body)

	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code: %d", response.StatusCode)
	}
	encode := determineEncoding(response.Body)
	reader := transform.NewReader(response.Body, encode.NewDecoder())
	return ioutil.ReadAll(bufio.NewReader(reader))
}

func determineEncoding(reader io.Reader) encoding.Encoding {
	peek, err := bufio.NewReader(reader).Peek(1024)
	if err != nil {
		return unicode.UTF8
	}

	encode, _, _ := charset.DetermineEncoding(peek, "")
	return encode
}
