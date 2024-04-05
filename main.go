package main

import (
	"encoding/json"
	"fmt"

	"github.com/gonzxlez/colibri"
	"github.com/gonzxlez/colibri/webextractor"
)

var rawRules = `{
    "method": "GET",
    "url":    "https://pkg.go.dev/search?q=xpath",
    "timeout": 10000,
    "selectors": {
        "packages": {
            "expr": "div.SearchSnippet",
            "all": true,
            "type": "css",
            "selectors": {
                "name": "//h2/a/text()",
                "path": "//h2/a/@href"
            }
        }
    }
}`

func main() {
	we, err := webextractor.New()
	if err != nil {
		panic(err)
	}

	var rules colibri.Rules
	err = json.Unmarshal([]byte(rawRules), &rules)
	if err != nil {
		panic(err)
	}

	output, err := we.Extract(&rules)
	if err != nil {
		panic(err)
	}

	fmt.Println("URL:", output.Response.URL())
	fmt.Println("Status code:", output.Response.StatusCode())
	fmt.Println("Content-Type", output.Response.Header().Get("Content-Type"))
	fmt.Println("Data:", output.Data)
}
