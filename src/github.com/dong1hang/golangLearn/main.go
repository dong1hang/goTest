package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"regexp"

	"golang.org/x/net/html/charset"

	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
)

func main() {
	resp, err := http.Get("https://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error : status code", resp.StatusCode)
		return
	}
	e := determineEncoding(resp.Body)
	utf8Reader := transform.NewReader(
		resp.Body, e.NewDecoder())
	fmt.Println(utf8Reader)
	all, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	//fmt.Printf("%s\n", all)
	printAllCity(all)

}

func determineEncoding(r io.Reader) encoding.Encoding {
	bytes, err := bufio.NewReader(r).Peek(1024)

	bufio.NewReader(r)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(bytes, "")
	return e
}

func printAllCity(counts []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)".[^>]*([^<]+)</a>`)
	matches := re.FindAllSubmatch(counts, -1)
	for _, m := range matches {
		// for _, subMatch := range m {

		// 	fmt.Printf("%s\n", subMatch)
		// }
		// fmt.Println()
		fmt.Printf("City:%s,URL:%s\n", m[2], m[1])
	}
	fmt.Printf("Matches found: %d\n", len(matches))

}
