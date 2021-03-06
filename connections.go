package main

import (
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strconv"
)

func readResponse(resp *http.Response, err error) []byte {

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Println(err)
	}

	return body

}

func connectWithHash(jarFingerprint int) []byte {
	requestbody := []byte("[" + strconv.Itoa(jarFingerprint) + "]")
	resp, err := http.Post(API+"fingerprint", "application/json", bytes.NewBuffer(requestbody))
	return readResponse(resp, err)
}

func connectWithProjectID(projectID string) []byte {
	resp, err := http.Get(API + "addon/" + projectID + "/files")
	return readResponse(resp, err)
}

/*
	Source: https://golangcode.com/download-a-file-from-a-url/
*/

// DownloadFile will download a url to a local file. It's efficient because it will
// write as it downloads and not load the whole file into memory.
func DownloadFile(fileName string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	// Create the file within the folder
	out, err := os.Create(path.Join(*downloadPath, fileName))
	if err != nil {
		return err
	}

	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
