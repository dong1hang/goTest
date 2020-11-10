package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"

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
	fmt.Printf("%s\n", all)

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
