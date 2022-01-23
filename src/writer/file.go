package writer

import (
	"io"
	"os"
	"regexp"
)

func Write(fileName string, content io.Reader) error {
	var blackList = regexp.MustCompile(`(&|>|<|\[|\]|:|\n|\r)*`)
	fileName = blackList.ReplaceAllString(fileName, "")

	e := removeIfExist(fileName)

	if nil != e {
		return e
	}

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

func removeIfExist(path string) error {
	if _, err := os.Stat(path); err == nil {
		return os.Remove(path)
	}

	return nil
}
