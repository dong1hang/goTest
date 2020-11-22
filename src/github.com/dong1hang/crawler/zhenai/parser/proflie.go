package parser

import (
	"regexp"
	"strconv"

	"github.com/dong1hang/crawler/engine"
	"github.com/dong1hang/crawler/model"
)

const ageRe = ``

func ParseProfile(contents []byte) engine.ParseResult {
	profile := model.Profile{}
	re := regexp.MustCompile(ageRe)
	match := re.FindSubmatch(contents)
	if match != nil {
		age, err := strconv.Atoi(string(match[1]))
		if err != nil {
			profile.Age = age
		}
	}
}
