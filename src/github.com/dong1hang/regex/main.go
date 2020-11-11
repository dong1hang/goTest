package main

import (
	"fmt"
	"regexp"
)

const text = `My email is ccmouse@gmail.com
ccmouse1@gmail.com    
ccmouse3@gmail.com
dds  ccmouse4@gmail.com.cn`

/** 正则表达式 */
func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@[a-zA-Z0-9.]+\.[a-zA-Z0-9]+`)
	match := re.FindAllString(text, -1)
	match1 := re.FindAllStringSubmatch(text, -1)
	fmt.Println(match)
	//fmt.Println(match1)
	for _, m := range match1 {
		fmt.Println(m)
	}

}
