package main

import (
	"archive/zip"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var (
	url = flag.String("url", "http://idea.medeming.com/jihuoma/images/jihuoma.zip", "jihuoma")
)

type wrapReader struct {
	reader io.ReadCloser
}

func (w *wrapReader) ReadAt(p []byte, off int64) (n int, err error){
	return w.reader.Read(p)
}

func (w *wrapReader) Close() {
	w.reader.Close()
}


func main() {
	flag.Parse()

	resp, err := http.DefaultClient.Get(*url)
	// Open a zip archive for reading.
	file, err := ioutil.TempFile("/tmp", "zip")
	if err != nil {
		log.Fatal(err)
		return
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}

	file.Write(bytes)
	log.Println(file.Name())
	defer file.Close()
	tmpFilePath := file.Name()

	r, err := zip.OpenReader(file.Name())
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range r.File {
		fileReader, err := f.Open()

		fmt.Println(fmt.Sprintf("=================%s===================", f.Name))
		fmt.Println()
		if err != nil {
			log.Fatal(err)
		}
		result, err := ioutil.ReadAll(fileReader)
		fmt.Println(string(result))
		fileReader.Close()

		fmt.Println("\n\n\n")
	}
	os.Remove(tmpFilePath)
}


