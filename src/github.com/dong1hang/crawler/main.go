package main

import (
	"github.com/dong1hang/crawler/engine"
	"github.com/dong1hang/crawler/zhenai/parser"
)

func main() {
	engine.Run(
		engine.Request{Url: "https://www.zhenai.com/zhenghun",
			ParserFunc: parser.ParseCityList,
		})

}
