package main

import (
	"fmt"
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

	data, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	fmt.Fprintln(w, string(data))
}
