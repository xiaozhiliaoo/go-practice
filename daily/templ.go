package main

import (
	"bytes"
	"fmt"
	"text/template"
)

type Res struct {
}

type SearchResult struct {
	Title   string
	Content string
}

func main() {

	var a = []string{"1", "2", "3"}
	fmt.Println(a[0:len(a)])

	// 模板定义
	tmpl := `{{range $i, $result := .}}
{{if $i}}\n\n{{end}}{{.Title}} {{.Content}}{{end}}`

	// 多条检索结果
	results := []SearchResult{
		{
			Title:   "Golang",
			Content: "language.",
		},
		{
			Title:   "hello",
			Content: "world",
		},
		{
			Title:   "i",
			Content: "love",
		},
	}

	// 执行模板
	var buf bytes.Buffer
	err := template.Must(template.New("searchResults").Parse(tmpl)).Execute(&buf, results)
	if err != nil {
		panic(err)
	}

	// 输出结果
	fmt.Println(buf.String())
}
