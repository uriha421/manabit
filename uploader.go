package main

import (
	"fmt"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"net/http"
)

func upload(w http.ResponseWriter, r *http.Request) {
	// receive a *.md file.
	r.ParseMultipartForm(1024)

	// define attribute "uploaded" and get the *.md file.
	fileHeader := r.MultipartForm.File["uploaded"][0]
	file, err := fileHeader.Open()
	if err != nil {
		return
	}

	md, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	data := blackfriday.MarkdownCommon([]byte(md))
  ioutil.WriteFile("templates/text1.html", data, 0644)
	fmt.Fprintln(w, string(data))
}
