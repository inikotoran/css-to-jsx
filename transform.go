package main

import (
	"errors"
	"fmt"
	"strings"
)

func transform(css string) (jsx string, err error) {
	attributes := strings.Split(css, "\n")
	for _, attribute := range attributes {
		keyVal := strings.Split(attribute, ": ")
		if len(keyVal) != 2 {
			err = errors.New("invalid css (maybe)")
			return
		}
		cssKey := keyVal[0]
		cssVal := keyVal[1]

		jsxKey := kebabToCamel(cssKey)
		jsxVal := cssVal[0:len(cssVal)-1]

		jsx = fmt.Sprintf("%v\n%v: '%v',", jsx, jsxKey, jsxVal)

	}

	return
}

func kebabToCamel(kebab string) (camel string) {
	meats := strings.Split(kebab, "-")
	for index, meat := range meats {
		if index == 0 {
			camel = meat
			continue
		}
		camel = fmt.Sprintf("%v%v", camel, strings.Title(meat))
	}
	return
}