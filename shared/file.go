package shared

import (
	"io/ioutil"
	"os"
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
