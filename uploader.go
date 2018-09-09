package main

import (
	"fmt"
	"gopkg.in/russross/blackfriday.v2"
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

	data := blackfriday.Run([]byte(md))
  text := `<link rel="stylesheet" type="text/css" href="../css/text.css">
  ` + string(data)
  ioutil.WriteFile("templates/text1.html", []byte(text), 0644)
	fmt.Fprintln(w, text)
}
