package utils

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

func WriteFile(path string, data []byte) error {
	err := ioutil.WriteFile(path, data, 0755)
	return err
}

func ReadFile(path string) ([]byte, error) {
	data, err := ioutil.ReadFile(path)
	return data, err
}

func CreateFile(path string) error {
	return WriteFile(path, nil)
}

func IsFileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

func DownloadFile(url string, filepath string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func GetCWD() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
	return dir
}
