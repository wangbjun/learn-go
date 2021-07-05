package handler

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func ListFileHandler(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/list/"):]
	log.Println("Path is:", path)
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	content, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(content)

	return nil
}
