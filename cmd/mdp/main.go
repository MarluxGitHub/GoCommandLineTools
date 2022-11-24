package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday/v2"
)

const (
	header = `<!Doctype html>
	<html>
	<head>
		<meta http-equiv="content-type" content="text/html" charset="utf-8">
		<title>Markdown Preview</title>
	</head>
	<body>`

	footer = `</body>
	</html>`
)

func main() {
	filename := flag.String("f", "", "markdown file to preview")

	flag.Parse()

	if *filename == "" {
		flag.Usage()
		os.Exit(1)
	}

	if err := run(*filename); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filename string) error {
	input, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	htmlData := parseContent(input)

	outName := fmt.Sprintf("%s.html", filepath.Base(filename))

	fmt.Println(outName)

	return saveHtml(outName, htmlData)
}

func parseContent(data []byte) []byte {
	output := blackfriday.Run(data)
	body := bluemonday.UGCPolicy().SanitizeBytes(output)

	var buffer bytes.Buffer

	buffer.WriteString(header)
	buffer.Write(body)
	buffer.WriteString(footer)

	return buffer.Bytes()
}

func saveHtml(filename string, data []byte) error {
	return ioutil.WriteFile(filename, data, 0644)
}