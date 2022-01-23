package writer

import (
	"io"
	"os"
	"regexp"
)

func Write(fileName string, content io.Reader) error {
	fileName = EscapeFileName(fileName)

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

func EscapeFileName(fileName string) string {
	var blackList = regexp.MustCompile(`(&|>|<|\[|\]|:|\n|\r)*`)

	return blackList.ReplaceAllString(fileName, "")
}

func removeIfExist(path string) error {
	if _, err := os.Stat(path); err == nil {
		return os.Remove(path)
	}

	return nil
}
