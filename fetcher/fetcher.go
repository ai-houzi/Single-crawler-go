package fetcher

import (
	"bufio"
	"golang.org/x/net/html/charset"
	"io"
	"net/http"
	"fmt"
	"io/ioutil"
	"golang.org/x/text/transform"
	"golang.org/x/text/encoding"
	"golang.org/x/text/encoding/unicode"
	"log"
)

func Fetch(url string) ([]byte, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil,
			fmt.Errorf("wrong status code :", resp.StatusCode)
	}
	e := determineEncoding(resp.Body)

	utf8reader := transform.NewReader(resp.Body, e.NewDecoder())

	return ioutil.ReadAll(utf8reader)

}

/**
 * 判定网页内容的编码格式
 *
 */
func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)

	if err != nil {
		log.Printf("fetcher error: %v" ,err)
		return unicode.UTF8
	}

	encoding, _, _ := charset.DetermineEncoding(bytes, "")

	return encoding
}
