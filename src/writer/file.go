package writer

import (
	"io"
	"os"
)

func Write(fileName string, content io.Reader) error {
	file, err := os.Create(fileName)

	if err != nil {
		return err
	}

	defer file.Close()

	_, err = io.Copy(file, content)
	if err != nil {
		return err
	}

	return nil
}
