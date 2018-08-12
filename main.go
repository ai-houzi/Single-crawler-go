package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
	"golang.org/x/text/transform"
	"io"
	"golang.org/x/net/html/charset"
	"bufio"
	"golang.org/x/text/encoding"
	"regexp"
)

func main() {


	printCityListAll(all)
	//fmt.Printf("%s\n", all)

}



/**
	正则匹配需要的内容
 */
func printCityListAll(contents []byte) {
	r := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

	matches := r.FindAllSubmatch(contents, -1)

	for _, m := range matches {

		fmt.Printf("city: %s, url: %s \n", m[2],m[1])

	}

	fmt.Printf("matches fond %d", len(matches))

}
